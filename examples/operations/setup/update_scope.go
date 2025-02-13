/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package setup

import (
	"encoding/json"
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func UpdateScopeSample(session *gosdk.SessionMethods) {

	scopeName := "customersdb"
	result, err := session.GetScope(scopeName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if result["errorId"] != nil {
		fmt.Println("scope ", scopeName, " not found")
		return
	}

	result["description"] = "Updated description!"

	resultScope, errScope := session.CreateUpdateScope(result, false)
	if errScope != nil {
		fmt.Println(errScope)
		return
	}

	prettyJSON, err := json.MarshalIndent(resultScope, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(prettyJSON))
}
