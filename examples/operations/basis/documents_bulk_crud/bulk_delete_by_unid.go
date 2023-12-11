package documentsbulkcrud

import (
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
	}

	UNIDList := []string{}
	for key, value := range result {
		if key == "unid" {
			UNIDList = append(UNIDList, value.(string))
		}
	}

	deleteResult, delErr := session.BulkDeleteDocumentByUnid("customersdb", UNIDList, "delete")
	if delErr != nil {
		fmt.Println(delErr)
	}
	fmt.Println(deleteResult)
}
