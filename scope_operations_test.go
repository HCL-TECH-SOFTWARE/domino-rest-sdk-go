/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
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

func TestAccessConnectorConfig_ScopeOperations(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	tests := []struct {
		name   string
		fields fields
		want   *ScopeOperationsMethods
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			if got := ac.ScopeOperations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.ScopeOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getScope(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		scopeName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ec.getScope(tt.args.scopeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getScopes(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
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
			ec := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ec.getScopes()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getScopes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getScopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_deleteScope(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		scopeName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ec.deleteScope(tt.args.scopeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.deleteScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.deleteScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_createUpdateScope(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		scope        map[string]interface{}
		createSchema bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ec.createUpdateScope(tt.args.scope, tt.args.createSchema)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.createUpdateScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.createUpdateScope() = %v, want %v", got, tt.want)
			}
		})
	}
}
