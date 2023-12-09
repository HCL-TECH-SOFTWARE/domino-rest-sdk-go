package documentscrud

import (
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
	}

	fmt.Println(result)
}
