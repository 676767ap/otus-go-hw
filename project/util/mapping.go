package util

import (
	"errors"
	"reflect"

	"github.com/676767ap/otus-go-hw/project/util/log"
)

func MappingStructure(input interface{}, output interface{}) error {
	if !compare(input, output) {
		err := errors.New("структуры не совпадают")
		log.Error(err)
		return err
	}
	fields := reflect.TypeOf(input).Elem()
	inputValues := reflect.ValueOf(input).Elem()
	outputValues := reflect.ValueOf(output).Elem()
	for i := 0; i < fields.NumField(); i++ {
		tagValue := fields.Field(i).Tag.Get("map")
		if tagValue != "-" {
			fieldName := fields.Field(i).Name
			valueField := inputValues.FieldByName(fieldName)
			outputValues.FieldByName(fieldName).Set(valueField)
		}
	}
	return nil
}

func compare(first interface{}, second interface{}) bool {
	firstType := reflect.TypeOf(first)
	secondType := reflect.TypeOf(second)
	return firstType == secondType
}
