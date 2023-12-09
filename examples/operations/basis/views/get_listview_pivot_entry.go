package views

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func GetListViewPivotEntrySample(session *gosdk.SessionMethods) {

	options := new(gosdk.GetListPivotViewEntryOptions)

	result, err := session.GetListViewPivotEntry("customersdb", "Customers", "name", *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
