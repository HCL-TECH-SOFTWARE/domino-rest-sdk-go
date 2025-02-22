/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package utils

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
)

// IsExpired utility checks expiration time validity in nanoseconds.
func IsExpired(expiry int) bool {
	if expiry == 0 {
		return true
	}
	return false
}

// RequestParameters structure is required parameters for using Request
// utility function.
type RequestParameters struct {
	Method string            `json:"method"`
	Url    string            `json:"url"`
	Body   io.Reader         `json:"body"`
	Header map[string]string `json:"header"`
}

// Request utility function is a simplified and customized http request
// to shorten http request calls.
func (rp *RequestParameters) Request() ([]byte, error) {
	request, err := http.NewRequest(strings.ToUpper(rp.Method), rp.Url, rp.Body)
	// Setting up headers
	for key, value := range rp.Header {
		request.Header.Set(key, value)
	}

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// FetchHttpMethods utility function provides list of method types.
func FetchHttpMethods() []string {

	list := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	return list

}

// SetArrayAny utilty gets an array list of objects, it overwrites the objects and removes
// duplicate entries.
func SetArrayAny(objectArray []interface{}, object []interface{}) []interface{} {

	if len(objectArray) == 0 {
		return object
	}

	for _, v := range object {
		objectArray = append(objectArray, v)
	}

	return objectArray
}

// GetMimeTypes utility gets an array list of mimes available from domino.
func GetMimeTypes() []string {
	list := []string{
		"application/json",
		"application/xml",
		"application/mime",
		"application/javascript",
		"multipart/form-data",
		"text/markdown",
		"text/plain",
		"",
	}

	return list
}

func GetApiListName() []string {

	list := []string{
		"basis",
		"setup",
		"admin",
	}

	return list

}

func StructToMap(in interface{}) (map[string]interface{}, error) {
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
		fieldVal := v.Field(i)

		if string(field.Tag) != "" {
			if field.Type == reflect.TypeOf(false) {
				out[strings.Split(field.Tag.Get("json"), ",")[0]] = fieldVal.Bool()
			} else if fieldVal.IsZero() {
				continue
			} else {
				out[strings.Split(field.Tag.Get("json"), ",")[0]] = fieldVal.Interface()
			}
		}
	}
	return out, nil
}
