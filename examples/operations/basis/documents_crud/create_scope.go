package documentscrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateScopeSample(session gosdk.SessionMethods) {

	scp := map[string]interface{}{
		"apiName":            "customersdb",
		"createSchema":       false,
		"description":        "The famous demo database",
		"icon":               "Base64 stuff, ie SVG",
		"iconName":           "beach",
		"isActive":           true,
		"maximumAccessLevel": "Manager",
		"nsfPath":            "customer.nsf",
		"schemaName":         "customers",
		"server":             "*",
	}

	result, err := session.CreateUpdateScope(scp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
