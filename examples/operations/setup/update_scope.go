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

func UpdateScopeSample(session *gosdk.SessionMethods) {

	result, err := session.GetScope("customersdb2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result["description"] = "Updated description!"

	resultScope, errScope := session.CreateUpdateScope(result)
	if errScope != nil {
		fmt.Println(errScope)
	}
	fmt.Println(resultScope)
}
