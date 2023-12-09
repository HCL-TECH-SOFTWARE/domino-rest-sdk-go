package documentscrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetDocumentSample(session *gosdk.SessionMethods) {
	options := new(gosdk.GetDocumentOptions)
	options.Mode = "default"
	options.RichTextAs = "mime"

	result, err := session.GetDocument("customersdb", "10BA715F2EB3B02885258A7B005D6DA2", *options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
