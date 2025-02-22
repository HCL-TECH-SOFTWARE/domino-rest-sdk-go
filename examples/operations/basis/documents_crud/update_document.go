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

func UpdateDocumentSample(session *gosdk.SessionMethods) {

	options := new(gosdk.UpdateDocumentOptions)
	options.Mode = "default"
	options.ParentUnid = ""
	options.RichTextAs = "mime"
	options.Revision = ""
	options.MarkUnread = false

	doc := new(gosdk.DocumentInfo)
	doc.Form = "Customer"
	doc.Meta.UNID = "10BA715F2EB3B02885258A7B005D6DA2"
	doc.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley Sopas",
		"Form":     "Customer",
	}

	result, err := session.UpdateDocument("customersdb", *doc, *options)
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
