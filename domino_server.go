/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseServer BaseDominoServer

type BaseDominoServer struct {
	baseUrl      string
	apiMap       map[string]interface{}
	connectorMap map[string]t.DominoRestConnector
	t.DominoRestServer
}

func DominoServer(baseUrl string) (t.DominoRestServer, error) {
	baseServer.baseUrl = baseUrl
	err := dominoServerApiLoader()
	if err != nil {
		return t.DominoRestServer{}, err
	}
	baseServer.AvailableAPIs = availableApis
	baseServer.GetDominoConnector = getDominoConnector
	baseServer.Basis = basis
	baseServer.Setup = setup
	baseServer.Admin = admin
	return baseServer.DominoRestServer, nil
}

func dominoServerApiLoader() error {
	req := h.RequestParams{}
	req.Method = http.MethodGet
	req.Url = baseServer.baseUrl + "/api"
	req.Body = nil
	req.Header = map[string]string{"Content-type": "application/json"}
	response, err := h.Request(req)
	if err != nil {
		return errors.New("API loader failed")
	}

	err = json.Unmarshal(response, &baseServer.apiMap)
	if err != nil {
		return err
	}

	return nil
}

func availableApis() (map[string]interface{}, error) {

	if len(baseServer.apiMap) > 0 {
		return baseServer.apiMap, nil
	}

	err := dominoServerApiLoader()
	if err != nil {
		return nil, err
	}

	return baseServer.apiMap, nil
}

func getDominoConnector(apiName string) (t.DominoRestConnector, error) {

	if len(baseServer.apiMap) == 0 {
		err := dominoServerApiLoader()
		if err != nil {
			return t.DominoRestConnector{}, err
		}
	}

	cMapValue, cMap := baseServer.connectorMap[apiName]
	if cMap {
		return cMapValue, nil
	}

	apiMapVal, hasKey := baseServer.apiMap[apiName].(map[string]interface{})
	data, _ := json.Marshal(apiMapVal)

	var apiMeta t.DominoApiMeta
	json.Unmarshal(data, &apiMeta)

	if hasKey {
		dc := DominoConnector(baseServer.baseUrl, apiMeta)
		baseServer.connectorMap = map[string]t.DominoRestConnector{}
		baseServer.connectorMap[apiName] = dc
		return dc, nil
	}

	return t.DominoRestConnector{}, errors.New("API " + apiName + " is not available on this server.")
}

func basis() (t.DominoRestConnector, error) {
	return getDominoConnector("basis")
}

func setup() (t.DominoRestConnector, error) {
	return getDominoConnector("setup")
}

func admin() (t.DominoRestConnector, error) {
	return getDominoConnector("admin")
}
