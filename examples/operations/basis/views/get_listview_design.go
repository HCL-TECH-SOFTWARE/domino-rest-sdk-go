package views

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetListViewDesignSample(session *gosdk.SessionMethods) {
	options := new(gosdk.GetDesignOptions)
	result, err := session.GetListView("customersdb", "customers", *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
