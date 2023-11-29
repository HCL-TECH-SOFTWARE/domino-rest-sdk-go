/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package main

import (
	sdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/logger"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

// Expects to enable usage of DominoConnector component functions
// Below example shows the usage of GetOperation and GetUrl from DominoConnector
func ExampleDominoConnector() {
	server, err := sdk.DominoServer(Env.BaseURL)
	if err != nil {
		logger.Pretty("Some error occured! ", server)
	}

	connector, err := server.GetDominoConnector("basis")
	if err != nil {
		logger.Pretty("Some error occured! ", err)
	}

	operation, err := connector.GetOperation("getOdataItem")
	if err != nil {
		logger.Pretty("Some error occured! ", err)
	}

	opts := t.DominoRequestOptions{}
	opts.DataSource = "demo"
	opts.Params = map[string]string{
		"unid":    "ABCD1234567890BCABCD1234567890BC",
		"name":    "customer",
		"$select": "name,age,hobbies",
	}

	url, err := connector.GetUrl(operation, "demo", opts.Params)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Pretty("Prints operation from GetOperation() ", operation)
	logger.Pretty("Prints url result from GetUrl() ", url)
}
