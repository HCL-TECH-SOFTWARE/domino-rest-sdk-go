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

func TestDominoServer(t *testing.T) {
	type args struct {
		baseUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    *ServerMethods
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DominoServer(tt.args.baseUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DominoServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_availableAPIs(t *testing.T) {
	tests := []struct {
		name string
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := availableAPIs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("availableAPIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConnector(t *testing.T) {
	type args struct {
		apiName string
	}
	tests := []struct {
		name    string
		args    args
		want    *ConnectorMethods
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConnector(tt.args.apiName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}
