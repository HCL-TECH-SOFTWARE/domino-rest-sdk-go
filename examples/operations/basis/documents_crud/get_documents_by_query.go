package documentscrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetDocumentsByQuerySample(session *gosdk.SessionMethods) {

	request := new(gosdk.GetDocumentByQueryRequest)
	request.Count = 1

	options := new(gosdk.GetDocumentByQueryOptions)
	options.Count = 1

	result, err := session.GetDocumentByQuery("customersdb", *request, gosdk.EXECUTE, *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
