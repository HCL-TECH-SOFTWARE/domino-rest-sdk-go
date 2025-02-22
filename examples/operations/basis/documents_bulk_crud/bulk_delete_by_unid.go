/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package documentsbulkcrud

import (
	"encoding/json"
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func BulkDeleteDocumentByUNID(session *gosdk.SessionMethods) {

	data1 := gosdk.DocumentJSON{}
	data1.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley So",
		"Form":     "Customer",
	}

	data2 := gosdk.DocumentJSON{}
	data2.Fields = map[string]interface{}{
		"category": []string{"Movie", "Series", "Anime"},
		"name":     "Jujutsu Kaisen",
		"Form":     "Customer",
	}

	docList := []gosdk.DocumentJSON{}
	docList = append(docList, data1)
	docList = append(docList, data2)

	richTextAs := new(gosdk.RichTextRepresentation)

	result, err := session.BulkCreateDocument("customersdb", docList, *richTextAs)
	if err != nil {
		fmt.Println(err)
		return
	}

	UNIDList := []string{}
	for _, doc := range result {
		UNIDList = append(UNIDList, doc["@meta"].(map[string]interface{})["unid"].(string))
	}

	deleteResult, delErr := session.BulkDeleteDocumentByUnid("customersdb", UNIDList, "delete")
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
