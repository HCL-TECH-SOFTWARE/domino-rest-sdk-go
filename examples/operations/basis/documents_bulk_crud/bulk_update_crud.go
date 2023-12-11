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
