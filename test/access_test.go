/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package test

import (
	"testing"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestDominoAccess(t *testing.T) {

	var testCases = []struct {
		title    string
		input    *gosdk.Config
		expected string
		TCType   string
	}{
		{
			title:    "FAIL: Missing BaseURL",
			expected: "URL should not be empty.",
			input: &gosdk.Config{
				BaseUrl: "",
				Credentials: gosdk.Credentials{
					Scope:     "",
					Type:      "",
					UserName:  "",
					Password:  "",
					AppID:     "",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing Type",
			expected: "OAUTH needs appSecret, appId and refreshToken",
			input: &gosdk.Config{
				BaseUrl: "http://localhost:8880",
				Credentials: gosdk.Credentials{
					Scope:     "",
					Type:      "",
					UserName:  "",
					Password:  "",
					AppID:     "",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing UserName for BASIC AUTH",
			expected: "BASIC authentication needs username and password.",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:     "",
					Type:      "BASIC",
					UserName:  "",
					Password:  "",
					AppID:     "",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing PassWord on BASIC AUTH",
			expected: "BASIC authentication needs username and password.",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:     "",
					Type:      "BASIC",
					UserName:  "username",
					Password:  "",
					AppID:     "",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing AppID on OAUTH",
			expected: "OAUTH needs appSecret, appId and refreshToken",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:     "$DATA",
					Type:      "OAUTH",
					UserName:  "",
					Password:  "",
					AppID:     "",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing AppSecret on OAUTH",
			expected: "OAUTH needs appSecret, appId and refreshToken",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:     "$DATA",
					Type:      "OAUTH",
					UserName:  "",
					Password:  "",
					AppID:     "random-string",
					AppSecret: "",
				},
			},
		},
		{
			title:    "FAIL: Missing RefreshToken on OAUTH",
			expected: "OAUTH needs appSecret, appId and refreshToken",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:        "$DATA",
					Type:         "OAUTH",
					UserName:     "",
					Password:     "",
					AppID:        "random-string",
					AppSecret:    "random-string",
					RefreshToken: "",
				},
			},
		},
		{
			title:    "FAIL: Invalid user account",
			TCType:   "INVALID_CREDENTIALS",
			expected: "",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:    "$DATA",
					Type:     "BASIC",
					UserName: "testuser",
					Password: "testpassword",
				},
			},
		},
		{
			title:    "SUCCESS: Complete credentials for BASIC AUTH",
			TCType:   "BASEURL_NOTEMPTY",
			expected: "https://frascati.projectkeep.io",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:    "$DATA",
					Type:     "BASIC",
					UserName: "username",
					Password: "password",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			access, err := testCase.input.DominoAccess()
			if err != nil {
				assert.Equal(t, testCase.expected, err.Error())

			} else {
				if testCase.TCType == "INVALID_CREDENTIALS" {
					token, _ := access.GetAccessToken()
					assert.Empty(t, token)
					assert.Equal(t, testCase.expected, token)
					assert.Equal(t, 0, access.GetExpiry())
				}
				if testCase.TCType == "BASEURL_NOTEMPTY" {
					assert.Equal(t, testCase.expected, access.GetBaseUrl())
				}
			}
		})
	}

}
