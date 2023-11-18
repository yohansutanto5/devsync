package util

import "fmt"

func ConvertToInt(val interface{}) int {
	// Use a type assertion to check if val is of type int
	if intValue, ok := val.(int); ok {
		return intValue
	}
	// If val is not of type int, return 0
	return 0
}

func ConvertToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
