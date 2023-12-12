/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package main

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func main() {

	server, err := gosdk.DominoServer("")
	if err != nil {
		fmt.Println(err)
	}

	connector, err := server.GetConnector("basis")
	if err != nil {
		fmt.Println(err)
	}

	operation, err := connector.GetOperation("getOdataItem")
	if err != nil {
		fmt.Println(err)
	}

	opts := new(gosdk.DominoRequestOptions)
	opts.DataSource = "demo"
	opts.Params = map[string]string{
		"unid":    "ABCD1234567890BCABCD1234567890BC",
		"name":    "customer",
		"$select": "name,age,hobbies",
	}

	url, err := connector.GetUrl(operation, "demo", opts.Params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)

}
