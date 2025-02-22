/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
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

// DocumentOperationsMethods contains list of available function for CRUD operation to domino.
type DocumentOperationsMethods struct {
	CreateDocument            func(dataSource string, doc DocumentJSON, options CreateDocumentOptions) (map[string]interface{}, error)
	GetDocument               func(dataSource string, unid string, options GetDocumentOptions) (map[string]interface{}, error)
	UpdateDocument            func(dataSource string, doc DocumentInfo, options UpdateDocumentOptions) (map[string]interface{}, error)
	PatchDocument             func(dataSource string, doc DocumentInfo, options UpdateDocumentOptions) (map[string]interface{}, error)
	DeleteDocument            func(dataSource string, doc DocumentInfo, mode string) (map[string]interface{}, error)
	DeleteDocumentByUnid      func(dataSource string, unid string, mode string) (map[string]interface{}, error)
	BulkGetDocument           func(dataSource string, unids []string, mode string, options BulkGetDocumentOptions) ([]map[string]interface{}, error)
	GetDocumentByQuery        func(dataSource string, getRequest GetDocumentByQueryRequest, action string, options GetDocumentByQueryOptions) ([]map[string]interface{}, error)
	BulkCreateDocument        func(dataSource string, docs []DocumentJSON, richTextAs RichTextRepresentation) ([]map[string]interface{}, error)
	BulkUpdateDocumentByQuery func(dataSource string, bulkUpdateReq BulkUpdateDocumentsByQueryRequest, richTextAs RichTextRepresentation) ([]map[string]interface{}, error)
	BulkDeleteDocuments       func(dataSource string, docs []DocumentInfo, mode string) ([]map[string]interface{}, error)
	BulkDeleteDocumentByUnid  func(dataSource string, unids []string, mode string) ([]map[string]interface{}, error)
}

// DocumentOperations creates an instance of operations to make list of available operation
// methods functional.
func (ac *AccessConnectorConfig) DocumentOperations() *DocumentOperationsMethods {

	operation := new(DocumentOperationsMethods)
	operation.CreateDocument = ac.createDocument
	operation.GetDocument = ac.getDocument
	operation.UpdateDocument = ac.updateDocument
	operation.PatchDocument = ac.patchDocument
	operation.DeleteDocument = ac.deleteDocument
	operation.DeleteDocumentByUnid = ac.deleteDocumentByUnid
	operation.BulkGetDocument = ac.bulkGetDocument
	operation.GetDocumentByQuery = ac.getDocumentByQuery
	operation.BulkCreateDocument = ac.bulkCreateDocument
	operation.BulkUpdateDocumentByQuery = ac.bulkUpdateDocumentByQuery
	operation.BulkDeleteDocumentByUnid = ac.bulkDeleteDocumentByUnid
	operation.BulkDeleteDocuments = ac.BulkDeleteDocuments

	return operation
}

// createDocument inserts new document in domino rest.
func (ac *AccessConnectorConfig) createDocument(dataSource string, doc DocumentJSON, options CreateDocumentOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	body := make(map[string]interface{})
	body["Form"] = doc.Form
	for k, v := range doc.Fields {
		body[k] = v
	}

	docMethods, err := DominoDocument(body)
	if err != nil {
		return nil, err
	}

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)

	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = docMethods.toDocumentJSON()

	response, err := ac.Execute("createDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// getDocument retrieves specific document in domino rest by unid.
func (ac *AccessConnectorConfig) getDocument(dataSource string, unid string, options GetDocumentOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	if len(strings.Trim(unid, "")) == 0 {
		return nil, errors.New("UNID must not be empty.")
	}

	if len(unid) != 32 {
		return nil, errors.New("UNID has an invalid value.")
	}

	params := make(map[string]string)
	params["unid"] = unid

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

	response, err := ac.Execute("getDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// updateDocument updates specific document in domino rest.
func (ac *AccessConnectorConfig) updateDocument(dataSource string, doc DocumentInfo, options UpdateDocumentOptions) (map[string]interface{}, error) {

	unid := doc.getUNID()

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	if len(strings.Trim(unid, "")) == 0 {
		return nil, errors.New("UNID must not be empty.")
	}

	if len(unid) != 32 {
		return nil, errors.New("UNID has an invalid value.")
	}

	params := make(map[string]string)
	params["unid"] = unid

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
	reqOptions.Body = doc.toDocumentJSON()

	response, err := ac.Execute("updateDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// patchDocument patches specific field in a document in rest domino.
func (ac *AccessConnectorConfig) patchDocument(dataSource string, doc DocumentInfo, options UpdateDocumentOptions) (map[string]interface{}, error) {

	unid := doc.getUNID()

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(unid, "")) == 0 {
		return nil, errors.New("UNID must not be empty")
	}

	if len(unid) != 32 {
		return nil, errors.New("UNID has an invalid value")
	}

	params := make(map[string]string)
	params["unid"] = unid

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
	reqOptions.Body = doc.toDocumentJSON()

	response, err := ac.Execute("patchDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// deleteDocument deletes specific document in rest domino.
func (ac *AccessConnectorConfig) deleteDocument(dataSource string, doc DocumentInfo, mode string) (map[string]interface{}, error) {

	unid := doc.getUNID()

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(unid, "")) == 0 {
		return nil, errors.New("UNID must not be empty")
	}

	if len(unid) != 32 {
		return nil, errors.New("UNID has an invalid value")
	}

	params := make(map[string]string)
	params["unid"] = unid
	if mode != "" {
		params["mode"] = mode
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := ac.Execute("deleteDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// deleteDocumentByUnid deletes specific document queried using unid.
func (ac *AccessConnectorConfig) deleteDocumentByUnid(dataSource string, unid string, mode string) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(unid, "")) == 0 {
		return nil, errors.New("UNID must not be empty")
	}

	if len(unid) != 32 {
		return nil, errors.New("UNID has an invalid value")
	}

	params := make(map[string]string)
	params["unid"] = unid
	if mode != "" {
		params["mode"] = mode
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := ac.Execute("deleteDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

// bulkGetDocument gets list of documents by unids.
func (ac *AccessConnectorConfig) bulkGetDocument(dataSource string, unids []string, mode string, options BulkGetDocumentOptions) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(unids) == 0 {
		return nil, errors.New("UNIDs array should not be empty")
	}

	for _, unid := range unids {
		if len(strings.Trim(unid, "")) == 0 {
			return nil, errors.New("one of the given UNIDs is empty")
		}
		if len(unid) != 32 {
			return nil, errors.New("one of the given UNIDs is invalid")
		}
	}

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	for key, value := range ops {
		params[key] = fmt.Sprintf("%v", value)
	}

	body := map[string]interface{}{
		"mode":  mode,
		"unids": unids,
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = body

	response, err := ac.Execute("bulkGetDocumentsByUnid", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getDocumentByQuery queries documents in domino rest with specific actions.
func (ac *AccessConnectorConfig) getDocumentByQuery(dataSource string, getRequest GetDocumentByQueryRequest, actions string, options GetDocumentByQueryOptions) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(getRequest.Query, "")) == 0 {
		return nil, errors.New("query inside request body should not be empty")
	}

	ops, err := utils.StructToMap(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	params["action"] = actions
	for key, value := range ops {
		params[key] = value.(string)
	}

	requestBody, err := utils.StructToMap(getRequest)
	if err != nil {
		return nil, err
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = requestBody

	response, err := ac.Execute("query", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkCreateDocument inserts multiple documents in domino rest.
func (ac *AccessConnectorConfig) bulkCreateDocument(dataSource string, docs []DocumentJSON, richTextAs RichTextRepresentation) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(docs) == 0 {
		return nil, errors.New("document array should not be empty")
	}

	var body []map[string]interface{}
	for _, doc := range docs {
		json := make(map[string]interface{})
		json["Form"] = doc.Form

		for key, value := range doc.Fields {
			json[key] = value
		}

		body = append(body, json)
	}

	params := make(map[string]string)
	if richTextAs == (RichTextRepresentation{}) {
		byteData, err := json.Marshal(richTextAs)
		if err != nil {
			return nil, err
		}
		params["richTextAs"] = string(byteData)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = map[string]interface{}{
		"documents": body,
	}

	response, err := ac.Execute("bulkCreateDocuments", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkUpdateDocumentByQuery update multiple documents based on query.
func (ac *AccessConnectorConfig) bulkUpdateDocumentByQuery(dataSource string, bulkUpdateReq BulkUpdateDocumentsByQueryRequest, richTextAs RichTextRepresentation) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(bulkUpdateReq.Query, "")) == 0 {
		return nil, errors.New("query inside request body should not be empty")
	}

	if bulkUpdateReq.ReplaceItems == nil || len(bulkUpdateReq.ReplaceItems) == 0 {
		return nil, errors.New("request replaceItems shoud not be empty")
	}

	params := make(map[string]string)
	if richTextAs == (RichTextRepresentation{}) {
		jsonByte, _ := json.Marshal(richTextAs)
		params["richTextAs"] = string(jsonByte)
	}

	requestBody, err := utils.StructToMap(bulkUpdateReq)
	if err != nil {
		return nil, err
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = requestBody

	response, err := ac.Execute("bulkUpdateDocumentsByQuery", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkDeleteDocumentByUnid deletes multiple documents based on unid list.
func (ac *AccessConnectorConfig) bulkDeleteDocumentByUnid(dataSource string, unids []string, mode string) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(unids) == 0 {
		return nil, errors.New("UNIDs array should not be empty")
	}

	for _, value := range unids {
		unid := value
		if len(strings.Trim(unid, "")) == 0 {
			return nil, errors.New("one of the given UNIDs is empty")
		}

		if len(unid) != 32 {
			return nil, errors.New("one of the given UNIDs is invalid")
		}
	}

	body := map[string]interface{}{
		"unids": unids,
	}

	if mode != "" {
		body["mode"] = mode
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = make(map[string]string)
	reqOptions.Body = body

	response, err := ac.Execute("bulkDeleteDocuments", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// BulkDeleteDocuments deletes multiple documents based on unids of given documents.
func (ac *AccessConnectorConfig) BulkDeleteDocuments(dataSource string, docs []DocumentInfo, mode string) ([]map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(docs) == 0 {
		return nil, errors.New("docs array should not be empty")
	}

	unids := []string{}
	for _, doc := range docs {
		unid := doc.getUNID()
		if len(strings.Trim(unid, "")) == 0 {
			return nil, errors.New("one of the given UNIDs is empty")
		}

		if len(unid) != 32 {
			return nil, errors.New("one of the given UNIDs is invalid")
		}

		unids = append(unids, unid)
	}

	body := map[string]interface{}{
		"unids": unids,
	}

	if mode != "" {
		body["mode"] = mode
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = make(map[string]string)
	reqOptions.Body = body

	response, err := ac.Execute("bulkDeleteDocuments", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}
