package basis

import (
	"fmt"

	sdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/logger"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

func CreateDocument() {

	// Setup access credentials
	drp := t.DominoRestAccessParams{}
	drp.BaseUrl = ""
	drp.Credentials.Type = ""
	drp.Credentials.Scope = ""
	drp.Credentials.UserName = ""
	drp.Credentials.Password = ""

	// Get domino access instance
	access, accessErr := sdk.DominoAccess(drp)
	if accessErr != nil {
		fmt.Println(accessErr.Error())
	}

	// Get domino server instance
	server, serverErr := sdk.DominoServer(access.BaseUrl)
	if serverErr != nil {
		fmt.Println(serverErr.Error())
	}

	// Get domino connector instance
	connector, connectorErr := server.Basis()
	if connectorErr != nil {
		fmt.Println(connectorErr.Error())
	}

	session := sdk.DominoUserSession(access, connector)

	operation := session.(*sdk.BaseDominoDocumentOperation)

	formData := t.DocumentJSON{}
	formData.Form = "Customer"

	formData.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley So",
		"Form":     "Customer",
	}

	options := t.CreateDocumentOptions{}
	options.RichTextAs = "mime"

	result, err := operation.CreateDocumentOperation("customersdb", access, connector, formData, options)
	if err != nil {
		fmt.Println(err.Error())
	}

	logger.Pretty(result)
}
