package utils

import (
	"errors"
	"strconv"
)

func UIntToString(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("couldn't convert a string to int")
	}
	return num, nil
}
