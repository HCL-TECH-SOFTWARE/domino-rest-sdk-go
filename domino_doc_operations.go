/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	"encoding/json"
	"errors"
	"strings"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseDocOperation = new(BaseDominoDocumentOperation)

type BaseDominoDocumentOperation struct {
	ExecuteOperation                    func(dc t.DominoRestConnector, da t.DominoRestAccess, operationId string, options t.DominoRequestOptions) (map[string]interface{}, error)
	GetDocumentOperation                func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, options t.GetDocumentOptions) (map[string]interface{}, error)
	CreateDocumentOperation             func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc t.DocumentJSON, options t.CreateDocumentOptions) (map[string]interface{}, error)
	UpdateDocumentOperation             func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc BaseDominoDocument, options t.UpdateDocumentOptions) (map[string]interface{}, error)
	PatchDocumentOperation              func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, docJsonPatch t.DocumentJSON, options t.UpdateDocumentOptions) (map[string]interface{}, error)
	DeleteDocumentOperation             func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc BaseDominoDocument, mode string) (map[string]interface{}, error)
	DeleteDocumentByUNIDOperation       func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, mode string) (map[string]interface{}, error)
	BulkGetDocumentsOperation           func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unids []string, options t.BulkGetDocumentsOptions) (map[string]interface{}, error)
	GetDocumentsByQueryOperation        func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, request t.GetDocumentsByQueryRequest, qaction t.QueryActions, options t.GetDocumentsByQueryOptions) (map[string]interface{}, error)
	BulkCreateDocumentsOperation        func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, docs []t.DocumentJSON, richTextAs t.RichTextRepresentation) (map[string]interface{}, error)
	BulkUpdateDocumentsByQueryOperation func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, request t.BulkUpdateDocumentsByQueryRequest, richTextAs t.RichTextRepresentation) (map[string]interface{}, error)
	BulkDeleteDocumentsByQueryOperation func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, docs []BaseDominoDocument, mode string) (map[string]interface{}, error)
	BulkDeleteDocumentsByUNIDOperation  func(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unids []string, mode string) (map[string]interface{}, error)
}

func DominoDocumentOperation() *BaseDominoDocumentOperation {
	baseDocOperation.ExecuteOperation = executeOperation
	baseDocOperation.GetDocumentOperation = getDocumentOperation
	baseDocOperation.CreateDocumentOperation = createDocumentOperation
	baseDocOperation.UpdateDocumentOperation = updateDocumentOperation
	baseDocOperation.PatchDocumentOperation = patchDocumentOperation
	baseDocOperation.DeleteDocumentOperation = deleteDocumentOperation
	baseDocOperation.DeleteDocumentByUNIDOperation = deleteDocumentByUNIDOperation
	baseDocOperation.BulkGetDocumentsOperation = bulkGetDocumentsOperation
	baseDocOperation.GetDocumentsByQueryOperation = getDocumentsByQueryOperation
	baseDocOperation.BulkCreateDocumentsOperation = bulkCreateDocumentsOperation
	baseDocOperation.BulkUpdateDocumentsByQueryOperation = bulkUpdateDocumentsByQueryOperation
	baseDocOperation.BulkDeleteDocumentsByQueryOperation = bulkDeleteDocumentsByQueryOperation
	baseDocOperation.BulkDeleteDocumentsByUNIDOperation = bulkDeleteDocumentsByUNIDOperation

	return baseDocOperation
}

func executeOperation(dc t.DominoRestConnector, da t.DominoRestAccess, operationId string, options t.DominoRequestOptions) (map[string]interface{}, error) {
	response, err := dc.Request(da, operationId, options)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func getDocumentOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, options t.GetDocumentOptions) (map[string]interface{}, error) {

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

	var mapOps map[string]string
	data, _ := json.Marshal(options)
	json.Unmarshal(data, &mapOps)

	for keyOps, keyVal := range mapOps {
		params[keyOps] = keyVal
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	response, err := baseDocOperation.ExecuteOperation(dc, da, "getDocument", reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func createDocumentOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc t.DocumentJSON, options t.CreateDocumentOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	body := t.DocumentBody{}
	body.Form = doc.Form
	body.Fields = doc.Fields

	dominoDoc, err := DominoDocument(body)
	if err != nil {
		return nil, err
	}

	ops, err := h.StructTranspose(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)

	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = dominoDoc.ToDocJson()

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "createDocument", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func updateDocumentOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc BaseDominoDocument, options t.UpdateDocumentOptions) (map[string]interface{}, error) {

	unid := doc.GetUNID()

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

	ops, err := h.StructTranspose(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = doc.ToDocJson()

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "updateDocument", reqOptions)
	if err != nil {
		return nil, err
	}
	return docResponse, nil
}

func patchDocumentOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, docJsonPatch t.DocumentJSON, options t.UpdateDocumentOptions) (map[string]interface{}, error) {

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

	ops, err := h.StructTranspose(options)
	if err != nil {
		return nil, err
	}

	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = docJsonPatch

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "patchDocument", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func deleteDocumentOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, doc BaseDominoDocument, mode string) (map[string]interface{}, error) {

	unid := doc.GetUNID()

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

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "deleteDocument", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func deleteDocumentByUNIDOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unid string, mode string) (map[string]interface{}, error) {

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

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "deleteDocument", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func bulkGetDocumentsOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unids []string, options t.BulkGetDocumentsOptions) (map[string]interface{}, error) {

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

	ops, err := h.StructTranspose(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = unids

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "bulkGetDocumentsByUnid", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func getDocumentsByQueryOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, request t.GetDocumentsByQueryRequest, qaction t.QueryActions, options t.GetDocumentsByQueryOptions) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(request.Query, "")) == 0 {
		return nil, errors.New("query inside request body should not be empty")
	}

	ops, err := h.StructTranspose(options)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	for key, value := range ops {
		params[key] = value.(string)
	}

	reqOptions := t.DominoRequestOptions{}
	reqOptions.DataSource = dataSource
	reqOptions.Params = params
	reqOptions.Body = request

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "query", reqOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func bulkCreateDocumentsOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, docs []t.DocumentJSON, richTextAs t.RichTextRepresentation) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	doc, err := h.StructTranspose(docs)
	if err != nil {
		return nil, err
	}

	if len(doc) == 0 {
		return nil, errors.New("document array should not be empty")
	}

	params := make(map[string]string)
	if richTextAs == (t.RichTextRepresentation{}) {
		jsonByte, _ := json.Marshal(richTextAs)
		params["richTextAs"] = string(jsonByte)
	}

	requestOptions := t.DominoRequestOptions{}
	requestOptions.DataSource = dataSource
	requestOptions.Params = params
	requestOptions.Body = map[string]interface{}{
		"documents": docs,
	}

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "bulkCreateDocuments", requestOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func bulkUpdateDocumentsByQueryOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, request t.BulkUpdateDocumentsByQueryRequest, richTextAs t.RichTextRepresentation) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	if len(strings.Trim(request.Query, "")) == 0 {
		return nil, errors.New("query inside request body should not be empty")
	}

	if request.ReplaceItems == nil || len(request.ReplaceItems) == 0 {
		return nil, errors.New("request replaceItems shoud not be empty")
	}

	params := make(map[string]string)
	if richTextAs == (t.RichTextRepresentation{}) {
		jsonByte, _ := json.Marshal(richTextAs)
		params["richTextAs"] = string(jsonByte)
	}

	requestOptions := t.DominoRequestOptions{}
	requestOptions.DataSource = dataSource
	requestOptions.Params = params
	requestOptions.Body = request

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "bulkUpdateDocumentByQuery", requestOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func bulkDeleteDocumentsByQueryOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, docs []BaseDominoDocument, mode string) (map[string]interface{}, error) {

	if len(strings.Trim(dataSource, "")) == 0 {
		return nil, errors.New("dataSource must not be empty")
	}

	doc, err := h.StructTranspose(docs)
	if err != nil {
		return nil, err
	}

	if len(doc) == 0 {
		return nil, errors.New("documents array should not be empty")
	}

	unids := []string{}

	for _, value := range docs {
		unid := value.GetUNID()
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

	requestOptions := t.DominoRequestOptions{}
	requestOptions.DataSource = dataSource
	requestOptions.Params = make(map[string]string)
	requestOptions.Body = body

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "bulkDeleteDocuments", requestOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}

func bulkDeleteDocumentsByUNIDOperation(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, unids []string, mode string) (map[string]interface{}, error) {

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

	requestOptions := t.DominoRequestOptions{}
	requestOptions.DataSource = dataSource
	requestOptions.Params = make(map[string]string)
	requestOptions.Body = body

	docResponse, err := baseDocOperation.ExecuteOperation(dc, da, "bulkDeleteDocuments", requestOptions)
	if err != nil {
		return nil, err
	}

	return docResponse, nil
}
