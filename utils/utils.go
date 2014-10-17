package utils

import (
	"fmt"
	"os"
)

func GetValue(variable string, def string) string {
	str := os.Getenv(variable)
	if str == "" {
		str = def
	}
	fmt.Printf("[utils] var: %s, result: %s\n", variable, str)
	return str
}
