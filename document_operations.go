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

// DocumentOperationsMethods contains list of available function for CRUD operation to domino.
type DocumentOperationsMethods struct {
	CreateDocument            func(dataSource string, doc DocumentJSON, options CreateDocumentOptions) (map[string]interface{}, error)
	GetDocument               func(dataSource string, unid string, options GetDocumentOptions) (map[string]interface{}, error)
	UpdateDocument            func(dataSource string, doc DocumentInfo, options UpdateDocumentOptions) (map[string]interface{}, error)
	PatchDocument             func(dataSource string, unid string, doc DocumentJSON, options UpdateDocumentOptions) (map[string]interface{}, error)
	DeleteDocument            func(dataSource string, doc DocumentInfo, mode string) (map[string]interface{}, error)
	DeleteDocumentByUnid      func(dataSource string, unid string, mode string) (map[string]interface{}, error)
	BulkGetDocument           func(dataSource string, unids []string, options BulkGetDocumentOptions) (map[string]interface{}, error)
	GetDocumentByQuery        func(dataSource string, getRequest GetDocumentByQueryRequest, action string, options GetDocumentByQueryOptions) (map[string]interface{}, error)
	BulkCreateDocument        func(dataSource string, docs []DocumentJSON, richTextAs RichTextRepresentation) (map[string]interface{}, error)
	BulkUpdateDocumentByQuery func(dataSource string, bulkUpdateReq BulkUpdateDocumentsByQueryRequest, richTextAs RichTextRepresentation) (map[string]interface{}, error)
	BulkDeleteDocumentByQuery func(dataSource string, docs []DocumentInfo, mode string) (map[string]interface{}, error)
	BulkDeleteDocumentByUnid  func(dataSource string, unids []string, mode string) (map[string]interface{}, error)
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
	operation.BulkDeleteDocumentByQuery = ac.bulkDeleteDocumentsByQuery
	operation.BulkDeleteDocumentByUnid = ac.bulkDeleteDocumentByUnid

	return operation
}

// createDocument inserts new document in domino rest.
func (ac *AccessConnectorConfig) createDocument(dataSource string, doc DocumentJSON, options CreateDocumentOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty.")
	}

	body := make(map[string]interface{})
	body["Form"] = doc.Form
	body["fields"] = doc.Fields

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
	reqOptions.Body = docMethods.ToDocumentJSON()

	response, err := ac.Execute("createDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
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

	return response, nil
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

	return response, nil
}

// patchDocument patches specific field in a document in rest domino.
func (ac *AccessConnectorConfig) patchDocument(dataSource string, unid string, doc DocumentJSON, options UpdateDocumentOptions) (map[string]interface{}, error) {

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
	reqOptions.Body = doc

	response, err := ac.Execute("patchDocument", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
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

	return response, nil
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

	return response, nil
}

// bulkGetDocument gets list of documents by unids.
func (ac *AccessConnectorConfig) bulkGetDocument(dataSource string, unids []string, options BulkGetDocumentOptions) (map[string]interface{}, error) {

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

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = unids

	response, err := ac.Execute("bulkGetDocumentsByUnid", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getDocumentByQuery queries documents in domino rest with specific actions.
func (ac *AccessConnectorConfig) getDocumentByQuery(dataSource string, getRequest GetDocumentByQueryRequest, actions string, options GetDocumentByQueryOptions) (map[string]interface{}, error) {

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
	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = getRequest

	response, err := ac.Execute("query", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkCreateDocument inserts multiple documents in domino rest.
func (ac *AccessConnectorConfig) bulkCreateDocument(dataSource string, docs []DocumentJSON, richTextAs RichTextRepresentation) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	docList, err := utils.StructToMap(docs)
	if err != nil {
		return nil, err
	}

	if len(docList) == 0 {
		return nil, errors.New("document array should not be empty.")
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
		"documents": docList,
	}

	response, err := ac.Execute("bulkCreateDocuments", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkUpdateDocumentByQuery update multiple documents based on query.
func (ac *AccessConnectorConfig) bulkUpdateDocumentByQuery(dataSource string, bulkUpdateReq BulkUpdateDocumentsByQueryRequest, richTextAs RichTextRepresentation) (map[string]interface{}, error) {

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

	reqOptions := new(DominoRequestOptions)
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = bulkUpdateReq

	response, err := ac.Execute("bulkUpdateDocumentByQuery", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// bulkDeleteDocumentsByQuery deletes multiple documents based on query.
func (ac *AccessConnectorConfig) bulkDeleteDocumentsByQuery(dataSource string, docs []DocumentInfo, mode string) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	docList, err := utils.StructToMap(docs)
	if err != nil {
		return nil, err
	}

	if len(docList) == 0 {
		return nil, errors.New("documents array should not be empty")
	}

	unids := []string{}
	for _, value := range docs {
		unid := value.getUNID()
		if len(strings.Trim(unid, "")) == 0 || unid == "" {
			return nil, errors.New("one of the given documents has empty UNID")
		}

		if len(unid) != 32 {
			return nil, errors.New("one of the given document has invalid UNID")
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

// bulkDeleteDocumentByUnid deletes multiple documents based on unid list.
func (ac *AccessConnectorConfig) bulkDeleteDocumentByUnid(dataSource string, unids []string, mode string) (map[string]interface{}, error) {

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
