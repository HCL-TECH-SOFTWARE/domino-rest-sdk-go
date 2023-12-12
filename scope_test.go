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

func TestDominoScope(t *testing.T) {
	type args struct {
		doc map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *ScopeInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DominoScope(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DominoScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScopeToJson(t *testing.T) {
	tests := []struct {
		name string
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScopeToJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScopeToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
