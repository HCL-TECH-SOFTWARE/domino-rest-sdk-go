/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package documentscrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func PatchDocumentSample(session *gosdk.SessionMethods) {

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
	}

	doc := new(gosdk.DocumentJSON)
	doc.Fields = result
	unid := result["unid"].(string)

	patchOptions := new(gosdk.UpdateDocumentOptions)
	patchOptions.Revision = result["revision"].(string)

	patchResult, patchErr := session.PatchDocument("customersdb", unid, *doc, *patchOptions)
	if patchErr != nil {
		fmt.Println(patchErr)
	}

	fmt.Println(patchResult)
}
