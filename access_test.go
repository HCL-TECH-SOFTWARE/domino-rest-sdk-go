/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func getTokenizer(user string) map[string]interface{} {

	file, err := os.Open("./pkg/resources/.testcerts/private.key")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(byteData)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	claims := map[string]interface{}{
		"sub":   user,
		"CN":    user,
		"iss":   "Domino REST API Mocha Tests",
		"scope": "$DATA",
		"aud":   "Domino",
		"exp":   3000,
	}

	bearer := jwt.New(jwt.SigningMethodRS256)
	signedKey, err := bearer.SignedString(key)
	if err != nil {
		fmt.Println(err)
	}

	result := map[string]interface{}{
		"bearer":    bearer,
		"claims":    claims,
		"signedKey": signedKey,
	}

	return result
}

func getCredentials() *Config {

	var config = new(Config)

	file, err := os.Open("./pkg/resources/env.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(byteData, &config)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return config
}

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

	token := getTokenizer("John Doe")
	creds := getCredentials()

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
		{
			name: "SUCCESS: Returns token if exist and not expired.",
			fields: fields{
				BaseUrl: creds.BaseUrl,
				Credentials: Credentials{
					Scope:      creds.Credentials.Scope,
					Type:       creds.Credentials.Type,
					UserName:   creds.Credentials.UserName,
					Password:   creds.Credentials.Password,
					Token:      token["signedKey"].(string),
					ExpiryTime: token["claims"].(map[string]interface{})["exp"].(int),
				},
			},
			wantToken: token["signedKey"].(string),
		},
		{
			name: "SUCCESS: Returns token if no existing token that is valid",
			fields: fields{
				BaseUrl: creds.BaseUrl,
				Credentials: Credentials{
					Scope:    creds.Credentials.Scope,
					Type:     creds.Credentials.Type,
					UserName: creds.Credentials.UserName,
					Password: creds.Credentials.Password,
				},
			},
		},
		{
			name: "SUCCESS: Returns token if no existing token that is valid",
			fields: fields{
				BaseUrl: creds.BaseUrl,
				Credentials: Credentials{
					Scope:        creds.Credentials.Scope,
					Type:         OAUTH,
					AppID:        "sample-id",
					AppSecret:    "sample-secret",
					RefreshToken: "sample-refresh-token",
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
			gotToken, err := c.getAccessToken()
			if err != nil {
				t.Error(err)
			}
			if tt.wantToken != "" {
				assert.Equal(t, gotToken, tt.wantToken)
				assert.Greater(t, len(gotToken), 0)
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
