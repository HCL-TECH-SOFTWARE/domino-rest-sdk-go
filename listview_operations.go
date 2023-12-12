/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

type ListViewOperationsMethods struct {
	GetListViewPivotEntry func(dataSource string, listViewName string, pivotColumn string, options GetListPivotViewEntryOptions) (map[string]interface{}, error)
	GetListView           func(dataSource string, designName string, options GetDesignOptions) (map[string]interface{}, error)
	GetListViews          func(dataSource string, options GetListViewOptions) (map[string]interface{}, error)
	CreateUpdateListView  func(dataSource string, listView ListViewBody, designName string, options CreateUpdateDesignOptions) (map[string]interface{}, error)
}

func (ac *AccessConnectorConfig) DominoListViewOperations() *ListViewOperationsMethods {
	listView := new(ListViewOperationsMethods)
	listView.GetListViewPivotEntry = ac.getListViewPivotEntry
	listView.GetListView = ac.getListView
	listView.GetListViews = ac.getListViews
	listView.CreateUpdateListView = ac.createUpdateListView
	return listView
}

func (ac *AccessConnectorConfig) getListViewEntry(dataSource string, listViewName string, options map[string]interface{}) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	if len(strings.Trim(listViewName, "")) == 0 {
		return nil, errors.New("name must not be empty.")
	}

	var subscriber interface{}
	params := make(map[string]string)
	params["name"] = listViewName

	for key, _ := range options {
		if key == "subscriber" {
			subscriber = options[key]
		}
		params[key] = fmt.Sprintf("%v", options[key])
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = subscriber

	response, err := ac.Execute("fetchViewEntries", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ac *AccessConnectorConfig) getListViewPivotEntry(dataSource string, listViewName string, pivotColumn string, options GetListPivotViewEntryOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	if len(strings.Trim(listViewName, "")) == 0 {
		return nil, errors.New("name must not be empty.")
	}

	if len(strings.Trim(pivotColumn, "")) == 0 {
		return nil, errors.New("pivotColumn must not be empty.")
	}

	params := make(map[string]string)
	params["name"] = listViewName
	params["pivotColumn"] = pivotColumn

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := ac.Execute("pivotViewEntries", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ac *AccessConnectorConfig) getListViews(dataSource string, options GetListViewOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	params := make(map[string]string)

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := ac.Execute("fetchViews", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ac *AccessConnectorConfig) createUpdateListView(dataSource string, listView ListViewBody, designName string, options CreateUpdateDesignOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(designName, "")) == 0 {
		return nil, errors.New("designName must not be empty")
	}

	params := make(map[string]string)
	params["designName"] = designName
	params["designType"] = "views"

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	listViewObj, listErr := DominoListView(listView)
	if listErr != nil {
		return nil, listErr
	}

	mappedListView, listErr := listViewObj.ToListViewJson()
	if listErr != nil {
		return nil, listErr
	}

	jsonData, err := json.Marshal(mappedListView)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = string(jsonData)

	response, err := ac.Execute("updateCreateDesign", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ac *AccessConnectorConfig) getListView(dataSource string, designName string, options GetDesignOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(designName, "")) == 0 {
		return nil, errors.New("designName must not be empty")
	}

	params := make(map[string]string)
	params["designName"] = designName
	params["designType"] = "views"

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := ac.Execute("getDesign", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}
