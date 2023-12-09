package documentscrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func DeleteDocumentSample(session *gosdk.SessionMethods) {

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

	doc := new(gosdk.DocumentInfo)
	doc.Fields = result

	deleteResult, delErr := session.DeleteDocument("customersdb", *doc, "delete")
	if err != nil {
		fmt.Println(delErr)
	}

	fmt.Println(deleteResult)
}
