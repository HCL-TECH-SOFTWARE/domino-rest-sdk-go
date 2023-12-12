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

	"github.com/stretchr/testify/assert"
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
		wantErr string
	}{
		{
			name: "FAIL: Missing BaseURL",
			fields: fields{
				BaseUrl: "",
			},
			wantErr: "URL should not be empty.",
		},
		{
			name: "FAIL: Missing Type",
			fields: fields{
				BaseUrl: "https://localhost:8880",
			},
			wantErr: "OAUTH needs appSecret, appId and refreshToken",
		},
		{
			name: "FAIL: Missing UserName for BASIC AUTH",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Type: "BASIC",
				},
			},
			wantErr: "BASIC authentication needs username and password.",
		},
		{
			name: "FAIL: Missing PassWord for BASIC AUTH",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Type:     "BASIC",
					UserName: "username",
				},
			},
			wantErr: "BASIC authentication needs username and password.",
		},
		{
			name: "FAIL: Missing AppId for OAUTH",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Type: "OAUTH",
				},
			},
			wantErr: "OAUTH needs appSecret, appId and refreshToken",
		},
		{
			name: "FAIL: Missing AppSecret for OAUTH",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Type:  "OAUTH",
					AppID: "random-string",
				},
			},
			wantErr: "OAUTH needs appSecret, appId and refreshToken",
		},
		{
			name: "FAIL: Missing refreshToken for OAUTH",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Type:      "OAUTH",
					AppID:     "random-string",
					AppSecret: "random-string",
				},
			},
			wantErr: "OAUTH needs appSecret, appId and refreshToken",
		},
		{
			name: "SUCCESS: Complete credentials for BASIC auth",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Scope:    "$DATA",
					Type:     "BASIC",
					UserName: "username",
					Password: "password",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			got, err := c.DominoAccess()
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}

			if !reflect.DeepEqual(got.GetBaseUrl(), tt.fields.BaseUrl) {
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
		{
			name: "SUCCESS: Get expiry with invalid credentials",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Scope:      "$DATA",
					Type:       "BASIC",
					UserName:   "username",
					Password:   "password",
					ExpiryTime: 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			assert.Equal(t, tt.want, c.getExpiry())
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
		{
			name: "SUCCESS: Get url with supplied config",
			fields: fields{
				BaseUrl: "https://localhost:8880",
				Credentials: Credentials{
					Scope:    "$DATA",
					Type:     "BASIC",
					UserName: "username",
					Password: "password",
				},
			},
			want: "https://localhost:8880",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:     tt.fields.BaseUrl,
				Credentials: tt.fields.Credentials,
			}
			assert.Equal(t, tt.want, c.getBaseUrl())
		})
	}
}
