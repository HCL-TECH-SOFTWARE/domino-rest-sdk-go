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
	"io/ioutil"
	"os"

	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/examples/basis"
	log "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/logger"
)

var Env EnvironmentInfo

type EnvironmentInfo struct {
	BaseURL     string
	Credentials struct {
		Scope    string
		Type     string
		Username string
		Password string
	}
}

func init() {
	jsonFile, err := os.Open("env.json")
	if err != nil {
		log.Error(err.Error)
	}
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteData, &Env)
}

func main() {

	//Examples for Domino Access
	// ExampleDominoAccess()
	// sampleDAFetchExpiry()
	// sampleDAFetchValidExpiry()

	// Examples for Domino Server
	// sampleDSFetchAPI()
	// sampleGetOperation()
	// sampleGetUrl()
	// sampleRequest()
	//TestStructTranspose()
	basis.CreateDocument()
}
