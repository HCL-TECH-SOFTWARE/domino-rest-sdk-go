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

func TestSessionConfig_DominoUserSession(t *testing.T) {
	type fields struct {
		AccessMethods    *AccessMethods
		ConnectorMethods *ConnectorMethods
	}
	tests := []struct {
		name   string
		fields fields
		want   *SessionMethods
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SessionConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			if got := s.DominoUserSession(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionConfig.DominoUserSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
