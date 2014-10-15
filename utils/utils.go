package utils

import (
	"os"
	"fmt"
)

func GetValue(variable string, def string) string {
	str := os.Getenv(variable)
	fmt.Printf("[utils] var: %s, value: %s\n", variable, str)
	if str == "" {
		return def
	} else {
		return str
	}
}
