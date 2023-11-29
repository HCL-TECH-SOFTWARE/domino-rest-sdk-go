/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package types

import (
	"io"
	"net/http"
)

type DominoRestConnector struct {
	BaseUrl         string                                                                                                                           `json:"baseUrl"`
	Meta            DominoApiMeta                                                                                                                    `json:"meta"`
	Request         func(dominoAccess DominoRestAccess, operationId string, options DominoRequestOptions) (map[string]interface{}, error)            `json:"request"`
	GetOperation    func(operationId string) (DominoRestOperation, error)                                                                            `json:"getOperation"`
	GetUrl          func(operation DominoRestOperation, scope string, params map[string]string) (string, error)                                      `json:"getUrl"`
	GetFetchOptions func(dominoAccess DominoRestAccess, operation DominoRestOperation, request DominoRequestOptions) (map[string]interface{}, error) `json:"getFetchOptions"`
}

type DominoRestOperation struct {
	Method   string        `json:"method"`
	Url      string        `json:"url"`
	Params   []interface{} `json:"params"`
	MimeType string        `json:"mimetype"`
}

type DominoRequestOptions struct {
	DataSource string            `json:"dataSource"`
	Params     map[string]string `json:"params"`
	Body       interface{}       `json:"body"`
}

type DominoRestResponse struct {
	Body    io.Reader
	headers http.Header
	ok      bool
}
