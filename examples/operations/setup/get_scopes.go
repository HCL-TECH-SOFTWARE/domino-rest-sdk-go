package setup

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateUpdateScopesSample(session *gosdk.SessionMethods) {
	result, err := session.GetScopes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
