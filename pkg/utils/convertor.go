package utils

import (
	"log"
	"strconv"
)

func UIntToString(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		// return 0, errors.New("couldn't convert a string to int")
		log.Println("************** Printing this ", s)
		return 0, nil
	}
	return num, nil
}

// func GetOptionalParamInt(c *fiber.Ctx, paramName string) (int, error) {

// 	// strParam := fmt.Sprintf("%s", c.Params(paramName))
// 	strParam := c.Params(paramName)
// 	log.Println("******* value of Param : ", strParam)
// 	var param int
// 	if len(strParam) > 0 {
// 		parm, err := StringToInt(strParam)
// 		if err != nil {
// 			// param = 0
// 			return 0, err
// 		}
// 		param = parm
// 	}
// 	return param, nil
// }
