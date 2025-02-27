/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package views

import (
	"encoding/json"
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateUpdateListView(session *gosdk.SessionMethods) {

	listViewData := new(gosdk.ListViewBody)
	listViewData.Columns = []gosdk.DesignColumnSimple{}

	data1 := gosdk.DesignColumnSimple{
		Formula:                "email",
		Name:                   "email",
		SeparateMultipleValues: false,
		Sort:                   "ascending",
		Title:                  "email",
	}

	data2 := gosdk.DesignColumnSimple{
		Formula:                "name",
		Name:                   "name",
		SeparateMultipleValues: false,
		Sort:                   "ascending",
		Title:                  "name",
	}

	listViewData.Name = "testgo"
	listViewData.SelectionFormula = "SELECT Form = \"Customers\""
	listViewData.Columns = append(listViewData.Columns, data1)
	listViewData.Columns = append(listViewData.Columns, data2)

	options := new(gosdk.CreateUpdateDesignOptions)

	result, err := session.CreateUpdateListView("customersdb", *listViewData, "testgo", *options)
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
