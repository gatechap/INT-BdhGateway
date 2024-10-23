package intutilities

import (
	"errors"
	"strconv"
)

func CheckUsername(xChannel string, username string, xUser string) string {
	var xUsername string

	if len(xChannel) > 0 {
		xUsername = xChannel
	} else if len(username) > 0 {
		xUsername = username
	} else {
		xUsername = xUser
	}
	return xUsername
}

func ToString(v interface{}) (string, error) {
	val, err := v.(string)
	if !err {
		return "",
			errors.New("can't convert to string")
	}
	return val, nil
}

// imit : pageSize
// offset :  pageNum * (pageNum-1) + 1

func GetLimit(pageSize string) string {
	var limit string
	limit = pageSize
	return limit
}

func GetOffset(pageNum int, pageSize int) string {
	if pageNum != 0 {
		var offset int
		offset = pageSize * (pageNum - 1)
		os := strconv.Itoa(offset)
		return os
	} else {
		return ""
	}

}

func StringToIngeter(param string) int {
	if param == "" {
		return 0
	} else {
		i, _ := strconv.Atoi(param)
		return i
	}

}
