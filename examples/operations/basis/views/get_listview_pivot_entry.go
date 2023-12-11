/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package views

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetListViewPivotEntrySample(session *gosdk.SessionMethods) {

	options := new(gosdk.GetListPivotViewEntryOptions)

	result, err := session.GetListViewPivotEntry("customersdb", "Customers", "name", *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
