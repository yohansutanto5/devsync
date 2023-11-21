package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertToInt(val string) int {
	// Use a type assertion to check if val is of type int
	result, err := strconv.Atoi(val)
	// If val is not of type int, return 0
	if err != nil {
		return 0
	} else {
		return result
	}

}

func ConvertToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func ConvertStruct(source interface{}, target interface{}) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Type().Field(i)
		targetField, exists := targetValue.Type().FieldByName(sourceField.Tag.Get("convert"))

		if exists {
			sourceFieldValue := sourceValue.Field(i)
			targetFieldValue := targetValue.FieldByName(targetField.Name)

			if sourceFieldValue.Kind() == reflect.Struct && targetFieldValue.Kind() == reflect.Struct {
				// If the field is a struct, recursively call convertStruct for nested struct conversion
				ConvertStruct(sourceFieldValue.Interface(), targetFieldValue.Addr().Interface())
			} else {
				// Copy the field value
				targetFieldValue.Set(sourceFieldValue)
			}
		}
	}
}
