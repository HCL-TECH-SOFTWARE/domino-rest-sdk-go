package views

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func CreateUpdateListView(session *gosdk.SessionMethods) {

	listViewData := new(gosdk.ListViewBody)
	listViewData.Columns = []gosdk.DesignColumnSimple{}

	data1 := gosdk.DesignColumnSimple{
		Formula:                "email",
		Name:                   "email",
		SeparateMultipleValues: false,
		Sort:                   "ascending",
		Title:                  "email",
	}

	data2 := gosdk.DesignColumnSimple{
		Formula:                "name",
		Name:                   "name",
		SeparateMultipleValues: false,
		Sort:                   "ascending",
		Title:                  "name",
	}

	listViewData.Columns = append(listViewData.Columns, data1)
	listViewData.Columns = append(listViewData.Columns, data2)

	options := new(gosdk.CreateUpdateDesignOptions)

	result, err := session.CreateUpdateListView("customersdb", *listViewData, "designName", *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
