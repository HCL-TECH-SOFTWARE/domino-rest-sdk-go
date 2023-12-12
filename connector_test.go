/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"reflect"
	"testing"
)

func TestConnectorConfig_DominoConnector(t *testing.T) {
	type fields struct {
		BaseUrl string
		ApiMeta ApiMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   *ConnectorMethods
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConnectorConfig{
				BaseUrl: tt.fields.BaseUrl,
				ApiMeta: tt.fields.ApiMeta,
			}
			if got := c.DominoConnector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectorConfig.DominoConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDominoRequestParameters_dominoRequest(t *testing.T) {
	type fields struct {
		DominoAccess         AccessMethods
		OperationID          string
		DominoRequestOptions DominoRequestOptions
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drp := &DominoRequestParameters{
				DominoAccess:         tt.fields.DominoAccess,
				OperationID:          tt.fields.OperationID,
				DominoRequestOptions: tt.fields.DominoRequestOptions,
			}
			got, err := drp.dominoRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoRequestParameters.dominoRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DominoRequestParameters.dominoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDominoRestOperation_getUrl(t *testing.T) {
	type fields struct {
		Method   string
		Url      string
		Params   []interface{}
		Mimetype string
	}
	type args struct {
		scope  string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dro := &DominoRestOperation{
				Method:   tt.fields.Method,
				Url:      tt.fields.Url,
				Params:   tt.fields.Params,
				Mimetype: tt.fields.Mimetype,
			}
			got, err := dro.getUrl(tt.args.scope, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoRestOperation.getUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DominoRestOperation.getUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOperation(t *testing.T) {
	type args struct {
		operationId string
	}
	tests := []struct {
		name    string
		args    args
		want    *DominoRestOperation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getOperation(tt.args.operationId)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchOptionsParameters_getFetchOptions(t *testing.T) {
	type fields struct {
		AccessMethods        AccessMethods
		DominoRestOperations *DominoRestOperation
		DominoRequestOptions DominoRequestOptions
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fo := &FetchOptionsParameters{
				AccessMethods:        tt.fields.AccessMethods,
				DominoRestOperations: tt.fields.DominoRestOperations,
				DominoRequestOptions: tt.fields.DominoRequestOptions,
			}
			got, err := fo.getFetchOptions()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchOptionsParameters.getFetchOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchOptionsParameters.getFetchOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operationLoader(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := operationLoader(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("operationLoader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkMandatory(t *testing.T) {
	type args struct {
		required  bool
		params    map[string]string
		paramName string
		scope     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMandatory(tt.args.required, tt.args.params, tt.args.paramName, tt.args.scope); (err != nil) != tt.wantErr {
				t.Errorf("checkMandatory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
