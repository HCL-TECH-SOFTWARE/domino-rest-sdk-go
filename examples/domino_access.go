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

// Expects to enable usage of DominoAccess functions.
// Below example shows usage of AccessToken, Expiry and other
// and other properties of DominoAccess component
func ExampleDominoAccess() {
	drp := t.DominoRestAccessParams{}
	drp.BaseUrl = Env.BaseURL
	drp.Credentials.Type = Env.Credentials.Type
	drp.Credentials.UserName = Env.Credentials.Username
	drp.Credentials.Password = Env.Credentials.Password

	dominoAccess, err := sdk.DominoAccess(drp)
	if err != nil {
		logger.Pretty("Some error occurred! ", err.Error())
	}

	token, err := dominoAccess.AccessToken()
	if err != nil {
		logger.Pretty("Some error occurred! ", err.Error())
	}
	expiry, err := dominoAccess.Expiry()
	if err != nil {
		logger.Pretty("Some error occured! ", err)
	}

	logger.Pretty("Prints token from AcceesToken() ", token)
	logger.Pretty("Prints token from DominoRestAccess.Token ", dominoAccess.Token)
	logger.Pretty("Prints expiration time from DominoRestAccess.ExpiryTime ", dominoAccess.ExpiryTime)
	logger.Pretty("Prints expiration time from Expiry() ", expiry)
}
