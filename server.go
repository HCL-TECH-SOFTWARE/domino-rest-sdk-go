/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"encoding/json"
)

var (
	apiMap    = make(map[string]interface{})
	serverURL string
)

type ServerMethods struct {
	GetApiList   func() map[string]interface{}
	GetConnector func(apiName string) (*ConnectorMethods, error)
}

// DominoServer function creates new instace of server.
// provides connector to domino.
func DominoServer(baseUrl string) (*ServerMethods, error) {

	loaderParams := new(ApiLoaderParameters)
	loaderParams.URL = baseUrl + "/api"
	data, err := loaderParams.ApiLoader()
	if err != nil {
		return nil, err
	}
	apiMap = data
	serverURL = baseUrl

	server := new(ServerMethods)
	server.GetApiList = availableAPIs
	server.GetConnector = getConnector

	return server, nil
}

// availableAPIs private function for fetching list of available api endpoints
// from domino.
func availableAPIs() map[string]interface{} {
	return apiMap
}

// getConnector gets new instance of connector to domino.
func getConnector(apiName string) (*ConnectorMethods, error) {
	var newData = make(map[string]interface{})
	apiMapVal, hasKey := apiMap[apiName].(map[string]interface{})
	if hasKey {
		data, err := json.Marshal(apiMapVal)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &newData)
		if err != nil {
			return nil, err
		}

	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return nil, err
	}

	var apiMeta = new(ApiMeta)
	err = json.Unmarshal(jsonData, &apiMeta)
	if err != nil {
		return nil, err
	}

	cc := new(ConnectorConfig)
	cc.BaseUrl = serverURL
	cc.ApiMeta = *apiMeta
	connector := cc.DominoConnector()

	return connector, nil

}
