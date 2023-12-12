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

func TestConfig_DominoAccess(t *testing.T) {
	type fields struct {
		BaseUrl     string
		Credentials Credentials
	}
	tests := []struct {
		name    string
		fields  fields
		want    *AccessMethods
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			got, err := c.DominoAccess()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.DominoAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.DominoAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_getExpiry(t *testing.T) {
	type fields struct {
		BaseUrl     string
		Credentials Credentials
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			if got := c.getExpiry(); got != tt.want {
				t.Errorf("Config.getExpiry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_getAccessToken(t *testing.T) {
	type fields struct {
		BaseUrl     string
		Credentials Credentials
	}
	tests := []struct {
		name      string
		fields    fields
		wantToken string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			gotToken, err := c.getAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.getAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("Config.getAccessToken() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestConfig_getBaseUrl(t *testing.T) {
	type fields struct {
		BaseUrl     string
		Credentials Credentials
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			if got := c.getBaseUrl(); got != tt.want {
				t.Errorf("Config.getBaseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
