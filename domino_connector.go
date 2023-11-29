/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var (
	baseConnector = new(BaseDominoConnector)
	schema        = make(map[string]t.DominoRestOperation)
)

type BaseDominoConnector struct {
	baseUrl string
	t.DominoApiMeta
	t.DominoRestConnector
}

func DominoConnector(baseUrl string, meta t.DominoApiMeta) t.DominoRestConnector {
	baseConnector.baseUrl = baseUrl
	baseConnector.DominoApiMeta = meta
	baseConnector.Request = request
	baseConnector.GetUrl = getUrl
	baseConnector.GetOperation = getOperation
	baseConnector.GetFetchOptions = getFetchOptions
	return baseConnector.DominoRestConnector
}

func getUrl(operation t.DominoRestOperation, scope string, params map[string]string) (string, error) {

	rawUrl := operation.Url
	workingUrl, _ := url.Parse(baseConnector.baseUrl + rawUrl)
	query := workingUrl.Query()

	for paramIndex, _ := range operation.Params {
		paramsKey := operation.Params[paramIndex].(map[string]interface{})["in"]
		if paramsKey != nil {
			var pname string
			var isRequired bool
			in := paramsKey.(string)

			if operation.Params[paramIndex].(map[string]interface{})["name"] != nil {
				pname = operation.Params[paramIndex].(map[string]interface{})["name"].(string)
			}
			if operation.Params[paramIndex].(map[string]interface{})["required"] != nil {
				isRequired = operation.Params[paramIndex].(map[string]interface{})["required"].(bool)
			}

			if in != "header" || in != "cookie" {
				err := checkMandatory(isRequired, params, pname, scope)
				if err != nil {
					return "", err
				}
				var newValue string
				if pname == "dataSource" {
					newValue = scope
				} else {
					newValue = params[pname]
				}
				if in == "path" {
					var searchFor = fmt.Sprintf("{%s}", pname)
					rawUrl = strings.ReplaceAll(rawUrl, searchFor, newValue)
				} else {
					//if params[pname] != "" {
					query.Set(pname, newValue)
					//}
				}
			}

		}

	}
	rawUrl = baseConnector.MountPath + rawUrl
	workingUrl.Path = rawUrl
	workingUrl.RawQuery = query.Encode()

	return workingUrl.String(), nil
}

func request(dominoAccess t.DominoRestAccess, operationId string, options t.DominoRequestOptions) (map[string]interface{}, error) {

	scopeVal := options.DataSource
	if len(schema) == 0 {
		err := connectorApiLoader()
		if err != nil {
			return nil, err
		}
	}

	operation, opErr := getOperation(operationId)
	if opErr != nil {
		return nil, opErr
	}

	url, urlErr := getUrl(operation, scopeVal, options.Params)
	if urlErr != nil {
		return nil, urlErr
	}

	params, fetchErr := getFetchOptions(dominoAccess, operation, options)
	if fetchErr != nil {
		return nil, fetchErr
	}

	req := h.RequestParams{}
	req.Url = url
	req.Method = params["method"].(string)
	req.Header = params["header"].(map[string]string)

	var myMap map[string]interface{}
	data, _ := json.Marshal(options.Body)
	json.Unmarshal(data, &myMap)

	req.Body = bytes.NewBuffer(data)
	response, err := h.Request(req)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	errs := json.Unmarshal(response, &result)
	if errs != nil {
		return nil, errs
	}
	return result, nil
}

func getOperation(operationId string) (t.DominoRestOperation, error) {

	if len(schema) == 0 {
		err := connectorApiLoader()
		if err != nil {
			return t.DominoRestOperation{}, err
		}
	}

	_, hasOperationId := schema[operationId]
	if hasOperationId {
		return schema[operationId], nil
	}

	return t.DominoRestOperation{}, errors.New("OperationId " + operationId + " is not available")
}

func connectorApiLoader() error {

	actualUrl := fmt.Sprintf("%s%s%s", baseConnector.baseUrl, baseConnector.MountPath, baseConnector.FileName)
	req := h.RequestParams{}
	req.Url = actualUrl
	req.Header = map[string]string{"content-type": "application/json"}
	response, err := h.Request(req)
	if err != nil {
		return err
	}

	resData := make(map[string]interface{})
	json.Unmarshal(response, &resData)

	oplErr := connectorOperationLoader(resData)
	if oplErr != nil {
		return oplErr
	}

	return nil
}

func connectorOperationLoader(openApi map[string]interface{}) error {
	var mimeType string
	paths := openApi["paths"].(map[string]interface{})

	for url, _ := range paths {
		urlInfo := paths[url].(map[string]interface{})

		for key, _ := range urlInfo {
			hasMethod := slices.Contains(h.FetchHttpMethods(), strings.ToUpper(key))
			if key == "parameters" || hasMethod {

				if hasMethod {
					operationId := urlInfo[key].(map[string]interface{})["operationId"]
					paramList := []interface{}{}

					if operationId != nil || operationId != "" {

						// http method level parameters
						methodParams := urlInfo[key].(map[string]interface{})["parameters"]
						if methodParams != nil {
							paramList = h.SetArrayAny(paramList, methodParams.([]interface{}))
						}

						// path level parameters
						pathParams := urlInfo["parameters"]
						if pathParams != nil {
							paramList = h.SetArrayAny(paramList, pathParams.([]interface{}))
						}

						// set mimetype if exist
						requestBody := urlInfo[key].(map[string]interface{})["requestBody"]
						if requestBody != nil {
							content := requestBody.(map[string]interface{})["content"]
							if content != nil {
								for mime, _ := range content.(map[string]interface{}) {
									if slices.Contains(h.GetMimeTypes(), mime) {
										mimeType = mime
									}
								}
							}
						} else {
							mimeType = ""
						}

						// set domino rest operation
						dro := t.DominoRestOperation{}
						dro.Method = key
						dro.Url = url
						dro.Params = paramList
						dro.MimeType = mimeType
						schema[operationId.(string)] = dro

					}
				}
			}
		}
	}

	return nil
}

func getFetchOptions(dominoAccess t.DominoRestAccess, operation t.DominoRestOperation, request t.DominoRequestOptions) (map[string]interface{}, error) {

	var params = make(map[string]string)
	var headers = make(map[string]string)
	var result = make(map[string]interface{})

	params = request.Params
	result["method"] = operation.Method

	if request.Body != nil {
		result["body"] = request.Body
	}

	for idx, _ := range operation.Params {
		ops := operation.Params[idx].(map[string]interface{})
		for pname, _ := range ops {

			var (
				in         string
				isRequired bool
			)

			if ops["in"] != nil && ops["required"] != nil {
				in = ops["in"].(string)
				isRequired = ops["required"].(bool)
			}

			// Check for mandatory parameters missing
			if ops[pname] == nil {
				if isRequired && in == "header" {
					return nil, errors.New("Parameter " + pname + " is mandatory.")
				}
			}

			if pname == "name" {
				if params[pname] != "" && in == "header" {
					headers[pname] = params[pname]
				}
			}
		}

	}

	if operation.MimeType != "" {
		headers["Content-Type"] = operation.MimeType
	}

	token, tokenErr := dominoAccess.AccessToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	headers["Authorization"] = "Bearer " + token
	result["header"] = headers

	return result, nil
}

func checkMandatory(required bool, params map[string]string, paramName string, scope string) error {

	var candidate string
	if required {

		if paramName == "dataSource" {
			candidate = scope
			if candidate == "" {
				candidate = params[paramName]
			}
		} else {
			candidate = params[paramName]
		}
		if candidate == "" {
			return errors.New("Parameter " + paramName + " is mandatory.")
		}

		return nil
	}

	return nil
}
