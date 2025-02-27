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
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

func DeleteScopeSample(session *gosdk.SessionMethods) {

	scopeData := map[string]interface{}{
		"apiName":            "sdkcpdlt",
		"createSchema":       false,
		"description":        "Scope to be deleted by Go SDK",
		"icon":               "Base64 stuff, preferably SVG",
		"iconName":           "beach",
		"isActive":           true,
		"maximumAccessLevel": "Manager",
		"nsfPath":            "Demo.nsf",
		"schemaName":         "demoapi",
		"server":             "*",
	}

	scp, err := gosdk.DominoScope(scopeData)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := utils.StructToMap(scp)

	result, errScp := session.CreateUpdateScope(data, false)
	if errScp != nil {
		fmt.Println(errScp)
		return
	}

	deleteResult, delErr := session.DeleteScope(result["apiName"].(string))
	if delErr != nil {
		fmt.Println(delErr)
		return
	}

	prettyJSON, err := json.MarshalIndent(deleteResult, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(prettyJSON))
}
