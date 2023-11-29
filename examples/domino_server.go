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
)

// Expects to enable usage of DominoServer component functions.
// Below example shows usage of AvailableAPIs from DominoServer
func ExampleDominoServer() {

	server, err := sdk.DominoServer(Env.BaseURL)
	if err != nil {
		logger.Pretty("Some error occured! ", server)
	}

	apiList, err := server.AvailableAPIs()
	if err != nil {
		logger.Pretty("Some error occured!", err)
	}

	logger.Pretty("Prints list of available APIs from AvailableAPI() ", apiList)

}
