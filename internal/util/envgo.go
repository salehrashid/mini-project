package util

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(path string, predefined string) string {
	res := os.Getenv(path)
	if res != "" {
		return res
	}
	return predefined
}

func GetPort(path string) string {
	res := GetInt(path)
	if res > 0 {
		return fmt.Sprintf(":%d", res)
	}
	return ""
}

func GetInt(path string) int {
	res := os.Getenv(path)
	if res != "" {
		ires, err := strconv.Atoi(res)
		if err == nil {
			return ires
		}
	}
	return 0
}
