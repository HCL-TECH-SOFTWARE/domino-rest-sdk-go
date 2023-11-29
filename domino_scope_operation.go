package keep_sdk

import (
	"errors"
	"strings"

	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseScopeOperation = new(BaseDominoScopeOperations)

type BaseDominoScopeOperations struct {
	GetScope          func(scopeName string, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error)
	GetScopes         func(da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error)
	DeleteScope       func(scopeName string, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error)
	CreateUpdateScope func(scope map[string]interface{}, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error)
}

func DominoScopeOperations() {
	baseScopeOperation.GetScope = getScope
	baseScopeOperation.GetScopes = getScopes
	baseScopeOperation.DeleteScope = deleteScope
	baseScopeOperation.CreateUpdateScope = createUpdateScope
}

func getScope(scopeName string, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error) {

	if len(strings.Trim(scopeName, "")) == 0 {
		return nil, errors.New("scopeName must not be empty")
	}

	params := make(map[string]string)
	params["scopeName"] = scopeName

	reqOptions := t.DominoRequestOptions{}
	reqOptions.Params = params

	response, err := baseDocOperation.ExecuteOperation(dc, da, "getScopeMapping", reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func getScopes(da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error) {

	params := make(map[string]string)

	reqOptions := t.DominoRequestOptions{}
	reqOptions.Params = params

	response, err := baseDocOperation.ExecuteOperation(dc, da, "fetchScopeMapping", reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func deleteScope(scopeName string, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error) {

	if len(strings.Trim(scopeName, "")) == 0 {
		return nil, errors.New("scopeName must not be empty")
	}

	params := make(map[string]string)
	params["scopeName"] = scopeName

	reqOptions := t.DominoRequestOptions{}
	reqOptions.Params = params

	response, err := baseDocOperation.ExecuteOperation(dc, da, "deleteScopeMapping", reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func createUpdateScope(scope map[string]interface{}, da t.DominoRestAccess, dc t.DominoRestConnector) (map[string]interface{}, error) {

	params := make(map[string]string)

	reqOptions := t.DominoRequestOptions{}
	reqOptions.Params = params
	reqOptions.Body = scope

	response, err := baseDocOperation.ExecuteOperation(dc, da, "createUpdateScopeMapping", reqOptions)
	if err != nil {
		return nil, err
	}

	return response, nil
}
