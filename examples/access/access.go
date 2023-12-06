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
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

var config = new(gosdk.Config)

func init() {
	file, err := os.Open("../pkg/resources/env.json")
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

	token, err := access.GetAccessToken()
	if err != nil {
		fmt.Println(err.Error())
	}

	expiry := access.GetExpiry()

	fmt.Println(token)
	fmt.Println(expiry)
	fmt.Println(utils.IsExpired(expiry))
}
