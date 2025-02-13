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

func GetScopesSample(session *gosdk.SessionMethods) {
	result, err := session.GetScopes()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(prettyJSON))
}
