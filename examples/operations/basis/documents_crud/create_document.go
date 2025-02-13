/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package documentscrud

import (
	"encoding/json"
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateDocumentSample(session *gosdk.SessionMethods) {
	formData := new(gosdk.DocumentJSON)
	formData.Form = "Customer"

	formData.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley So",
		"Form":     "Customer",
	}

	options := new(gosdk.CreateDocumentOptions)
	options.RichTextAs = "mime"

	result, err := session.CreateDocument("customersdb", *formData, *options)
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
