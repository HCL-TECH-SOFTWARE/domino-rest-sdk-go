/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"errors"
	"strings"
)

type ScopeOperationsMethods struct {
	GetScope    func(scopeName string) (map[string]interface{}, error)
	GetScopes   func() (map[string]interface{}, error)
	DeleteScope func(scopeName string) (map[string]interface{}, error)
	// CreateUpdateScope func(scope map[string]interface{}, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error)
}

// ScopeOperations creates an instance of operations to make list of available operation
// methods functional.
func (ac *AccessConnectorConfig) ScopeOperations() *ScopeOperationsMethods {

	ec := new(ExecuteConfig)
	ec.AccessMethods = ac.AccessMethods
	ec.ConnectorMethods = ac.ConnectorMethods

	scopeOperation := new(ScopeOperationsMethods)
	scopeOperation.GetScope = ec.getScope
	scopeOperation.GetScopes = ec.getScopes
	scopeOperation.DeleteScope = ec.deleteScope

	return scopeOperation
}

// getScope retrieves scope from domino rest by specifying scopeName
func (ec *ExecuteConfig) getScope(scopeName string) (map[string]interface{}, error) {

	if len(strings.Trim(scopeName, "")) == 0 {
		return nil, errors.New("scopeName must not be empty.")
	}

	params := make(map[string]string)
	params["scopeName"] = scopeName

	reqOptions := new(DominoRequestOptions)
	reqOptions.Params = params

	response, err := ec.Execute("getScopeMapping", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getScopes retrieves list of scopes from domino rest.
func (ec *ExecuteConfig) getScopes() (map[string]interface{}, error) {

	params := make(map[string]string)

	reqOptions := new(DominoRequestOptions)
	reqOptions.Params = params

	response, err := ec.Execute("fetchScopeMapping", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// deleteScope delete scopes in domino rest using scopeName
func (ec *ExecuteConfig) deleteScope(scopeName string) (map[string]interface{}, error) {

	if len(strings.Trim(scopeName, "")) == 0 {
		return nil, errors.New("scopeName must not be empty")
	}

	params := make(map[string]string)
	reqOptions := new(DominoRequestOptions)
	reqOptions.Params = params

	response, err := ec.Execute("deleteScopeMapping", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// createUpdateScope updates or insert scope if not exist.
func (ec *ExecuteConfig) createUpdateScope(scope map[string]interface{}) (map[string]interface{}, error) {

	params := make(map[string]string)

	reqOptions := new(DominoRequestOptions)
	reqOptions.Params = params
	reqOptions.Body = scope

	response, err := ec.Execute("createUpdateScopeMapping", *reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}
