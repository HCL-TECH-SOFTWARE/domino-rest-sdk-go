/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var (
	baseUserSession = new(BaseDominoUserSession)
	access          = baseUserSession.DominoRestAccess
	connector       = baseUserSession.DominoRestConnector
	docOperation    = new(BaseDominoDocumentOperation)
	scopeOperation  = new(BaseDominoScopeOperations)
)

type BaseDominoUserSession struct {
	t.DominoRestAccess
	t.DominoRestConnector
}

func DominoUserSession(da t.DominoRestAccess, dc t.DominoRestConnector) interface{} {
	baseUserSession.DominoRestAccess = da
	baseUserSession.DominoRestConnector = dc
	docOperation = DominoDocumentOperation()

	return docOperation
}

func Request(operationId string, options t.DominoRequestOptions) (map[string]interface{}, error) {

	response, err := baseUserSession.DominoRestConnector.Request(baseUserSession.DominoRestAccess, operationId, options)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func GetDocument(dataSource string, unid string, options t.GetDocumentOptions) (map[string]interface{}, error) {
	return docOperation.GetDocumentOperation(dataSource, access, connector, unid, options)
}

func CreateDocument(dataSource string, doc t.DocumentJSON, options t.CreateDocumentOptions) (map[string]interface{}, error) {
	return docOperation.CreateDocumentOperation(dataSource, access, connector, doc, options)
}

func UpdateDocument(dataSource string, doc BaseDominoDocument, options t.UpdateDocumentOptions) (map[string]interface{}, error) {
	return docOperation.UpdateDocumentOperation(dataSource, access, connector, doc, options)
}

func PatchDocument(dataSource string, unid string, docJsonPatch t.DocumentBody, options t.UpdateDocumentOptions) (map[string]interface{}, error) {
	json := t.DocumentJSON{}
	json.Fields = docJsonPatch.Fields
	return docOperation.PatchDocumentOperation(dataSource, access, connector, unid, json, options)
}

func DeleteDocument(dataSource string, doc BaseDominoDocument, mode string) (map[string]interface{}, error) {
	return docOperation.DeleteDocumentOperation(dataSource, access, connector, doc, mode)
}

func DeleteDocumentByUNID(dataSource string, unid string, mode string) (map[string]interface{}, error) {
	return docOperation.DeleteDocumentByUNIDOperation(dataSource, access, connector, unid, mode)
}

func BulkGetDocuments(dataSource string, unids []string, options t.BulkGetDocumentsOptions) (map[string]interface{}, error) {
	return docOperation.BulkGetDocumentsOperation(dataSource, access, connector, unids, options)
}

func BulkCreateDocuments(dataSource string, docs []t.DocumentJSON, richTextAs t.RichTextRepresentation) (map[string]interface{}, error) {
	return docOperation.BulkCreateDocumentsOperation(dataSource, access, connector, docs, richTextAs)
}

func BulkUpdateDocumentsByQuery(dataSource string, request t.BulkUpdateDocumentsByQueryRequest, richTextAs t.RichTextRepresentation) (map[string]interface{}, error) {
	return docOperation.BulkUpdateDocumentsByQueryOperation(dataSource, access, connector, request, richTextAs)
}

func CreateUpdateScope(scope map[string]interface{}) (map[string]interface{}, error) {
	return scopeOperation.CreateUpdateScope(scope, access, connector)
}

func GetScope(scopeName string) (map[string]interface{}, error) {
	return scopeOperation.GetScope(scopeName, access, connector)
}

func GetScopes() (map[string]interface{}, error) {
	return scopeOperation.GetScopes(access, connector)
}

func DeleteScope(scopeName string) (map[string]interface{}, error) {
	return scopeOperation.DeleteScope(scopeName, access, connector)
}
