/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestParams struct {
	Method string            `json:"method"`
	Url    string            `json:"url"`
	Body   io.Reader         `json:"body"`
	Header map[string]string `json:"header"`
}

func Request(rp RequestParams) ([]byte, error) {
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func FetchHttpMethods() []string {

	list := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	return list

}

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

func GetBytes(key interface{}) ([]byte, error) {
	buf, ok := key.([]byte)
	if !ok {
		return nil, fmt.Errorf("ooops, did not work")
	}

	return buf, nil
}
