/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package documentsbulkcrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func BulkUpdateDocumentByQuery(session *gosdk.SessionMethods) {

	request := new(gosdk.BulkUpdateDocumentsByQueryRequest)
	request.Query = "form = 'Customer' and name = 'Alien'"
	request.ReplaceItems = map[string]interface{}{
		"category": []string{"Friendly"},
	}

	richTextAs := new(gosdk.RichTextRepresentation)

	result, err := session.BulkUpdateDocumentByQuery("customersdb", *request, *richTextAs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
