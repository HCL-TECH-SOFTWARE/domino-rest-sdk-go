/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

// Operation list details by operationId
var schema = make(map[string]*DominoRestOperation)
var config = new(ConnectorConfig)

// ConnectorConfig is a required parameter in order to use connector instance.
type ConnectorConfig struct {
	BaseUrl string
	ApiMeta ApiMeta
}

// ApiMeta is a sub requirements in order to use connector instance.
type ApiMeta struct {
	FileName  string `json:"fileName"`
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Version   string `json:"version"`
}

// DominoRequestParameters structure is required properties to get an
// instance of connector.
type DominoRequestParameters struct {
	DominoAccess AccessMethods
	OperationID  string
	DominoRequestOptions
}

// DominoRequestOptions structure contains fields mapped data values
// being fetch by getFetchOptions function.
type DominoRequestOptions struct {
	DataSource string            `json:"dataSource"`
	Params     map[string]string `json:"params"`
	Body       interface{}       `json:"body"`
}

// DominoRestOperation struct contains fields required to fill and generate
// proper url endpoint to request in domino.
type DominoRestOperation struct {
	Method   string        `json:"method"`
	Url      string        `json:"url"`
	Params   []interface{} `json:"params"`
	Mimetype string        `json:"mimetype"`
}

// FetchOptionsParameters are required arguments for getFetchOptions function.
type FetchOptionsParameters struct {
	AccessMethods        AccessMethods
	DominoRestOperations *DominoRestOperation
	DominoRequestOptions DominoRequestOptions
}

// ConnectorMethods structure contains list of connector function that can be
// used externally.
type ConnectorMethods struct {
	Request         func(*DominoRequestParameters) (map[string]interface{}, error)
	GetUrl          func(dro *DominoRestOperation, scope string, params map[string]string) (string, error)
	GetOperation    func(operationId string) (*DominoRestOperation, error)
	GetFetchOptions func(dro *FetchOptionsParameters) (map[string]interface{}, error)
}

// DominoConnector function provides connector instance and connector methods
// that can be used externally.
func (c *ConnectorConfig) DominoConnector() *ConnectorMethods {

	config = c

	connector := new(ConnectorMethods)
	connector.Request = (*DominoRequestParameters).dominoRequest
	connector.GetUrl = (*DominoRestOperation).getUrl
	connector.GetOperation = getOperation
	connector.GetFetchOptions = (*FetchOptionsParameters).getFetchOptions

	return connector
}

// dominoRequest private function that directly request to domino with clean and
// transformed request data, forms or parameters.
func (drp *DominoRequestParameters) dominoRequest() (map[string]interface{}, error) {

	var scope string
	scope = drp.DataSource

	// validates if schema array list is empty
	// if true, will call apiLoader.
	if len(schema) == 0 {

		url := fmt.Sprintf("%s%s%s", config.BaseUrl, config.ApiMeta.MountPath, config.ApiMeta.FileName)
		loaderParams := new(ApiLoaderParameters)
		loaderParams.URL = url

		data, err := loaderParams.ApiLoader()
		if err != nil {
			return nil, err
		}

		err = operationLoader(data)
		if err != nil {
			return nil, err
		}

	}

	// getOperation by operationId
	operation, err := getOperation(drp.OperationID)
	if err != nil {
		return nil, err
	}

	// produces complete domino url with parameter values for request
	url, err := operation.getUrl(scope, drp.Params)
	if err != nil {
		return nil, err
	}

	options := new(FetchOptionsParameters)
	options.AccessMethods = drp.DominoAccess
	options.DominoRestOperations = operation
	options.DominoRequestOptions = drp.DominoRequestOptions

	params, err := options.getFetchOptions()
	if err != nil {
		return nil, err
	}

	request := new(utils.RequestParameters)
	request.Url = url
	request.Method = params["method"].(string)
	request.Header = params["header"].(map[string]string)

	data, err := json.Marshal(options.DominoRequestOptions.Body)
	if err != nil {
		return nil, err
	}

	request.Body = bytes.NewBuffer(data)
	response, err := request.Request()
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// getUrl private function that parses required url endpoint to request in domino.
func (dro *DominoRestOperation) getUrl(scope string, params map[string]string) (string, error) {

	rawURL := dro.Url
	workingURL, err := url.Parse(config.BaseUrl + rawURL)
	if err != nil {
		return "", err
	}

	query := workingURL.Query()

	for paramIndex, _ := range dro.Params {

		paramsKey := dro.Params[paramIndex].(map[string]interface{})["in"]
		if paramsKey != nil {

			var (
				pname      string
				isRequired bool
			)

			in := paramsKey.(string)
			if dro.Params[paramIndex].(map[string]interface{})["name"] != nil {
				pname = dro.Params[paramIndex].(map[string]interface{})["name"].(string)
			}
			if dro.Params[paramIndex].(map[string]interface{})["required"] != nil {
				isRequired = dro.Params[paramIndex].(map[string]interface{})["required"].(bool)
			}

			if in != "header" && in != "cookie" {
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
					rawURL = strings.ReplaceAll(rawURL, searchFor, newValue)
				} else {
					// if params[pname] != "" {
					// 	query.Set(pname, newValue)
					// }
					query.Set(pname, newValue)
				}
			}

		}

	}

	rawURL = config.ApiMeta.MountPath + rawURL
	workingURL.Path = rawURL
	workingURL.RawQuery = query.Encode()

	return workingURL.String(), nil
}

// getOperation private function provides operation based on operationId parsed
// from domino api list information.
func getOperation(operationId string) (*DominoRestOperation, error) {

	// validates if schema array list is empty
	// if true, will call apiLoader.
	if len(schema) == 0 {

		url := fmt.Sprintf("%s%s%s", config.BaseUrl, config.ApiMeta.MountPath, config.ApiMeta.FileName)
		loaderParams := new(ApiLoaderParameters)
		loaderParams.URL = url

		data, err := loaderParams.ApiLoader()
		if err != nil {
			return nil, err
		}

		err = operationLoader(data)
		if err != nil {
			return nil, err
		}

	}

	_, hasOperationId := schema[operationId]
	if hasOperationId {
		return schema[operationId], nil
	}

	return nil, errors.New("OperationId " + operationId + " is not available")
}

// getFetchOptions private function provide other additional fields to request in
// domino.
func (fo *FetchOptionsParameters) getFetchOptions() (map[string]interface{}, error) {

	var (
		params  = make(map[string]string)
		headers = make(map[string]string)
		result  = make(map[string]interface{})
	)

	params = fo.DominoRequestOptions.Params
	result["method"] = fo.DominoRestOperations.Method

	if fo.DominoRequestOptions.Body != nil {
		result["body"] = fo.DominoRequestOptions.Body
	}

	for index, _ := range fo.DominoRestOperations.Params {

		ops := fo.DominoRestOperations.Params[index].(map[string]interface{})

		for pname, _ := range ops {

			var (
				in         string
				isRequired bool
			)

			if ops["in"] != nil && ops["required"] != nil {
				in = ops["in"].(string)
				isRequired = ops["required"].(bool)
			}

			if ops[pname] == nil {

				if isRequired && in == "header" {
					headers[pname] = params[pname]
				}

			}

		}

	}

	if fo.DominoRestOperations.Mimetype != "" {
		headers["Content-Type"] = fo.DominoRestOperations.Mimetype
	}

	token, err := fo.AccessMethods.GetAccessToken()
	if err != nil {
		return nil, err
	}

	headers["Authorization"] = "Bearer " + token
	result["header"] = headers

	return result, nil
}

// operationLoader private function loads up operation information from domino endpoint
// api list.
func operationLoader(data map[string]interface{}) error {

	var mimeType string
	// OpenAPI list of endpoint details
	paths := data["paths"].(map[string]interface{})

	for url, _ := range paths {

		urlInfo := paths[url].(map[string]interface{})

		for key, _ := range urlInfo {

			hasMethod := slices.Contains(utils.FetchHttpMethods(), strings.ToUpper(key))

			if key == "parameters" || hasMethod {

				if hasMethod {
					operationId := urlInfo[key].(map[string]interface{})["operationId"]
					paramList := []interface{}{}

					// http method level parameters
					methodParams := urlInfo[key].(map[string]interface{})["parameters"]
					if methodParams != nil {
						paramList = utils.SetArrayAny(paramList, methodParams.([]interface{}))
					}

					// path level parameters
					pathParams := urlInfo["parameters"]
					if pathParams != nil {
						paramList = utils.SetArrayAny(paramList, pathParams.([]interface{}))
					}

					// set mime values if it exist in fetched apilist
					requestBody := urlInfo[key].(map[string]interface{})["requestBody"]
					if requestBody != nil {

						content := requestBody.(map[string]interface{})["content"]
						if content != nil {

							for mime, _ := range content.(map[string]interface{}) {
								if slices.Contains(utils.GetMimeTypes(), strings.ToLower(mime)) {
									mimeType = mime
								}
							}

						}

					} else {
						mimeType = ""
					}

					// sets domino rest required data for request
					dro := new(DominoRestOperation)
					dro.Method = key
					dro.Url = url
					dro.Params = paramList
					dro.Mimetype = mimeType
					schema[operationId.(string)] = dro
				}

			}

		}

	}

	return nil
}

// checkMandatory private functions checks for url's mandatory properties.
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
