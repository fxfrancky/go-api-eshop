package handlers

import (
	"net/http"
	"strconv"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/fxfrancky/go-api-eshop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetMe func to get current user.
// @Description Get current user.
// @Summary Logout  of the the API.
// @Tags User
// @Accept json
// @Produce json
// @Param	Authorization	header		string	true	"Authentication header"
// @Security ApiKeyAuth
// @Router /api/v1/users/me [get]
func (h *Handler) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

// GetUser func to get a user
// @Summary Get a user
// @Description Get a user. Auth required
// @ID get-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param email path string true "Email of the user to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} userResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/users/{email} [get]
func (h *Handler) GetUserByEmail(c *fiber.Ctx) error {

	email := c.Params("email")

	u, err := h.userRepository.GetUserByEmail(email)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

// UpdateUser func to update a user
// @Summary Update a user
// @Description Update a User. Auth is required
// @ID update-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the user to update"
// @Param user body updateUserRequest true "User to update"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} userResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [put]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	u, err := h.userRepository.GetUserById(id)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	payload := &updateUserRequest{}
	payload.populateUpdateUser(u)
	if err := payload.bindUpdateUser(c, u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if err = h.userRepository.UpdateUser(u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

// DeleteUser func to delete a User
// @Summary Delete a User
// @Description Delete a User. Auth is required
// @ID delete-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the user to delete"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} userResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	u, err := h.userRepository.GetUserById(id)

	if u == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("User"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err = h.userRepository.DeleteUser(u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "product deleted!"})
}

// AllUser func to display all User
// @Summary Display all Users
// @Description Display all Users. Auth is required
// @ID all-users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of products returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of products (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} userListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/users/{limit}/{offset} [get]
func (h *Handler) AllUsers(c *fiber.Ctx) error {
	var (
		users []models.User
		count int64
	)
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	users, count, err = h.userRepository.AllUsers(offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newUserListResponse(users, count))
}
