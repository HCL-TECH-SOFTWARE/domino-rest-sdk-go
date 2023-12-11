package setup

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func UpdateScopeSample(session *gosdk.SessionMethods) {

	result, err := session.GetScope("customersdb2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result["description"] = "Updated description!"

	resultScope, errScope := session.CreateUpdateScope(result)
	if errScope != nil {
		fmt.Println(errScope)
	}
	fmt.Println(resultScope)
}
