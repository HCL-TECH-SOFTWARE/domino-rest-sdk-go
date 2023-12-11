package setup

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetScopeSample(session *gosdk.SessionMethods) {

	result, err := session.GetScope("customersdb2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
