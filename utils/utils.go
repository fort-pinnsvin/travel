package utils

import "os"

func GetValue(variable string, def string) string {
	str := os.Getenv(variable)
	if str == "" {
		return def
	} else {
		return str
	}
}
