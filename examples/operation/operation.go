/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

var config = new(gosdk.Config)

func init() {
	file, err := os.Open("../../pkg/resources/env.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(byteData, &config)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

func main() {

	access, err := config.DominoAccess()
	if err != nil {
		fmt.Println(err.Error())
	}

	server, err := gosdk.Server(access.GetBaseUrl())
	if err != nil {
		fmt.Println(err.Error())
	}

	connector, err := server.GetConnector("basis")
	if err != nil {
		fmt.Println(err.Error())
	}

	sessionCfg := new(gosdk.SessionConfig)
	sessionCfg.AccessMethods = access
	sessionCfg.ConnectorMethods = connector

	session := sessionCfg.DominoUserSession()

	getDocumentExample(session)

}

func createDocumentExample(session *gosdk.SessionMethods) {
	formData := new(gosdk.DocumentJSON)
	formData.Form = "Customer"

	formData.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley So",
		"Form":     "Customer",
	}

	options := new(gosdk.CreateDocumentOptions)
	options.RichTextAs = "mime"

	result, err := session.CreateDocument("customersdb", *formData, *options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func getDocumentExample(session *gosdk.SessionMethods) {

	options := new(gosdk.GetDocumentOptions)
	options.Mode = "default"
	options.RichTextAs = "mime"

	result, err := session.GetDocument("customersdb", "10BA715F2EB3B02885258A7B005D6DA2", *options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func updateDocumentExample(session *gosdk.SessionMethods) {

	options := new(gosdk.UpdateDocumentOptions)
	options.Mode = "default"
	options.ParentUnid = ""
	options.RichTextAs = "mime"
	options.Revision = ""
	options.MarkUnread = false

	doc := new(gosdk.DocumentInfo)
	doc.Form = "Customer"
	doc.Meta.UNID = "10BA715F2EB3B02885258A7B005D6DA2"
	doc.Fields = map[string]interface{}{
		"category": []string{"Super", "Grand", "Master"},
		"name":     "Wesley Sopas",
		"Form":     "Customer",
	}

	result, err := session.UpdateDocument("customersdb", *doc, *options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
