package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/fxfrancky/go-api-eshop/config"
	"github.com/fxfrancky/go-api-eshop/internal/initializers"
	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/fxfrancky/go-api-eshop/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Signup (Register) func to creates a new user.
// @Description Signup/Register a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param signInput body SignUpInput true "SignInput"
// @Success 201 {object} models.UserResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		502				{string}	string	"Status BadGateway"
// @Router /api/v1/auth/register [post]
func (h *Handler) SignUpUser(c *fiber.Ctx) error {
	var payload *SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.NewError(err))
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Photo:    payload.Photo,
		IsAdmin:  payload.IsAdmin,
	}

	// result := initializers.DB.Create(&newUser)
	result := h.userRepository.RegisterUser(newUser)
	if result != nil && strings.Contains(result.Error(), "duplicated key not allowed") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.BadGateway())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}

// SignInUser func to login a user to the API
// @Description Signin/Login  a user to the API.
// @Summary login a user to the API.
// @Tags User
// @Accept json
// @Produce json
// @Param signInput body SignInInput true "Sign In Input"
// @Success 200 {string}	string	"answer"
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Router /api/v1/auth/login [post]
func (h *Handler) SignInUser(c *fiber.Ctx) error {
	var payload *SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	// var user models.User
	user, err := h.userRepository.GetUserByEmail(payload.Email)

	if user == nil && err == nil {
		return c.Status(fiber.StatusForbidden).JSON(utils.InvalidCredentials())
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.InvalidCredentials())
	}

	config, _ := config.LoadConfig(".")

	accessTokenDetails, err := utils.CreateToken(utils.UIntToString(user.ID), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	refreshTokenDetails, err := utils.CreateToken(utils.UIntToString(user.ID), config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	ctx := context.TODO()
	now := time.Now()

	errAccess := initializers.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, utils.UIntToString(user.ID), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(errAccess))
	}

	errRefresh := initializers.RedisClient.Set(ctx, refreshTokenDetails.TokenUuid, utils.UIntToString(user.ID), time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(errRefresh))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	userRoles := []string{}
	userRoles = append(userRoles, "customer")
	if user.IsAdmin {
		userRoles = append(userRoles, "admin")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "user": models.FilterUserRecord(user), "roles": userRoles})
}

// RefreshAccessToken func to refresh the current TOKEN
// @Description refresh the current TOKEN.
// @Summary refresh the current TOKEN
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} status "success"
// @Security ApiKeyAuth
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Router /api/v1/auth/refresh [get]
func (h *Handler) RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := config.LoadConfig(".")
	ctx := context.TODO()

	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(utils.NewError(err))
	}

	userid, err := initializers.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	// var user models.User
	// err = initializers.DB.First(&user, "id = ?", userid).Error
	uId, err := utils.StringToInt(userid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	user, err := h.userRepository.GetUserById(uId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(utils.NewError(err))

		}
	}

	accessTokenDetails, err := utils.CreateToken(utils.UIntToString(user.ID), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	now := time.Now()

	errAccess := initializers.RedisClient.Set(ctx, accessTokenDetails.TokenUuid, utils.UIntToString(user.ID), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(errAccess))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

// LogoutUser func to logout of the the API.
// @Description Logout  of the the API.
// @Summary Logout  of the the API.
// @Tags User
// @Accept json
// @Produce json
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {string}  string "success"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/logout [get]
func (h *Handler) LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := config.LoadConfig(".")
	ctx := context.TODO()

	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	access_token_uuid := c.Locals("access_token_uuid").(string)
	_, err = initializers.RedisClient.Del(ctx, tokenClaims.TokenUuid, access_token_uuid).Result()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(utils.NewError(err))
	}

	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
