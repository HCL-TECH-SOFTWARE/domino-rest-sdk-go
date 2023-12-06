/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer

// commons.go contains general and reusable functions, properties and struct information.
package gosdk

import (
	"encoding/json"

	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

// AccessConnectorConfig is a pointer struct that is supplied as required
// parameters to a session and operations.
type AccessConnectorConfig struct {
	AccessMethods    AccessMethods
	ConnectorMethods ConnectorMethods
}

// ExecuteConfig is a receiver pointer struct required for initiating execute
// function request to domino server.
type ExecuteConfig struct {
	AccessMethods    AccessMethods
	ConnectorMethods ConnectorMethods
}

// ApiLoaderParameters is a receiver pointer struct required for loading list
// of available api in domino server.
type ApiLoaderParameters struct {
	URL string
}

// ApiLoader function fetches list of api from domino.
func (a *ApiLoaderParameters) ApiLoader() (map[string]interface{}, error) {

	params := new(utils.RequestParameters)
	params.Url = a.URL
	params.Header = map[string]string{
		"content-type": "application/json",
	}

	response, err := params.Request()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Execute merges/grooms request parameters and executes api call to domino rest.
func (ec *ExecuteConfig) Execute(operationId string, options DominoRequestOptions) (map[string]interface{}, error) {

	reqParams := new(DominoRequestParameters)
	reqParams.OperationID = operationId
	reqParams.DominoRequestOptions = options
	reqParams.DominoAccess = ec.AccessMethods

	response, err := ec.ConnectorMethods.Request(reqParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
