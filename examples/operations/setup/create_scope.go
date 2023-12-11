/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package setup

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateScopeSample(session *gosdk.SessionMethods) {

	scp := map[string]interface{}{
		"apiName":            "customersdb",
		"createSchema":       false,
		"description":        "The famous demo database",
		"icon":               "Base64 stuff, ie SVG",
		"iconName":           "beach",
		"isActive":           true,
		"maximumAccessLevel": "Manager",
		"nsfPath":            "customer.nsf",
		"schemaName":         "customers",
		"server":             "*",
	}

	result, err := session.CreateUpdateScope(scp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
