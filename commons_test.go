/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer

// commons.go contains general and reusable functions, properties and struct information.
package gosdk

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiLoaderParameters_ApiLoader(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr string
	}{
		{
			name: "ERROR: Returns invalid response for invalid URL.",
			fields: fields{
				URL: "http://localhost",
			},
			want:    map[string]interface{}(nil),
			wantErr: "invalid character 'p' after top-level value",
		},
		{
			name: "SUCCESS: Returns invalid response for invalid URL.",
			fields: fields{
				URL: "http://frascati.projectkeep.io",
			},
			want: map[string]interface{}(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApiLoaderParameters{
				URL: tt.fields.URL,
			}
			got, err := a.ApiLoader()
			if err != nil {
				if tt.fields.URL == "http://localhost" {
					assert.Equal(t, tt.wantErr, err.Error())
				}
			}
			if err != nil {
				assert.Equal(t, got, tt.want)
			}

		})
	}
}

func TestAccessConnectorConfig_Execute(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		operationId string
		options     DominoRequestOptions
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
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.Execute(tt.args.operationId, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
