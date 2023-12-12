/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
	docsCrud "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/examples/operations/basis/documents_crud"
)

var config = new(gosdk.Config)

func init() {
	file, err := os.Open("../../pkg/resources/env.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(byteData, &config)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

func main() {

	access, err := config.DominoAccess()
	if err != nil {
		fmt.Println(err.Error())
	}

	server, err := gosdk.DominoServer(access.GetBaseUrl())
	if err != nil {
		fmt.Println(err.Error())
	}

	connector, err := server.GetConnector("basis")
	if err != nil {
		fmt.Println(err.Error())
	}

	sessionCfg := new(gosdk.SessionConfig)
	sessionCfg.AccessMethods = access
	sessionCfg.ConnectorMethods = connector

	session := sessionCfg.DominoUserSession()

	docsCrud.GetDocumentSample(session)

}
