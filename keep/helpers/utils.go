/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package helpers

import (
	"errors"
	"reflect"
	"slices"
	"strings"
)

func SetArrayAny(objectArray []interface{}, object []interface{}) []interface{} {

	if len(objectArray) == 0 {
		return object
	}

	for _, v := range object {
		objectArray = append(objectArray, v)
	}

	return objectArray
}

func Omit(data interface{}, fields ...string) interface{} {

	newData := reflect.ValueOf(data)
	structType := newData.Type()

	newStruct := make([]reflect.StructField, 0)
	for i := 0; i < newData.NumField(); i++ {
		field := structType.Field(i)
		if !slices.Contains(fields, field.Name) {
			sf := reflect.StructField{
				Name:      field.Name,
				Type:      field.Type,
				PkgPath:   field.PkgPath,
				Tag:       field.Tag,
				Offset:    field.Offset,
				Index:     field.Index,
				Anonymous: field.Anonymous,
			}
			newStruct = append(newStruct, sf)
		}
	}
	newType := reflect.StructOf(newStruct)
	return reflect.New(newType).Interface()
}

func Pick(data interface{}, fields ...string) reflect.Value {
	newData := reflect.ValueOf(data)
	structType := newData.Type()

	newStruct := make([]reflect.StructField, 0)
	for i := 0; i < newData.NumField(); i++ {
		field := structType.Field(i)
		if slices.Contains(fields, field.Name) {
			sf := reflect.StructField{
				Name:      field.Name,
				Type:      field.Type,
				PkgPath:   field.PkgPath,
				Tag:       field.Tag,
				Offset:    field.Offset,
				Index:     field.Index,
				Anonymous: field.Anonymous,
			}
			newStruct = append(newStruct, sf)
		}
	}
	newType := reflect.StructOf(newStruct)

	return reflect.New(newType)
}

func StructTranspose(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // Non-structural return error
		return nil, errors.New("Unable to convert to map, invalid struct.")
	}

	t := v.Type()
	// Traversing structure fields
	// Specify the tagName value as the key in the map; the field value as the value in the map
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if string(field.Tag) != "" {
			out[strings.Split(field.Tag.Get("json"), ",")[0]] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func HasFieldStructProps(object interface{}, fieldName string) bool {

	field, isExist := reflect.TypeOf(object).Elem().FieldByName(fieldName)
	if !isExist {
		if string(field.Tag) != "" {
			return true
		}
		return false
	}

	return true

}
