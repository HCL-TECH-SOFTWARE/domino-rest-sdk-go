/* ========================================================================== *
 * Copyright (C) 2025 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

package views

import (
	"encoding/json"
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetListViewEntrySample(session *gosdk.SessionMethods) {

	options := new(gosdk.GetListViewEntryOptions)

	result, err := session.GetListViewEntry("customersdb", "Customers", *options)
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
