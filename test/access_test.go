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
)

func TestAccessCredentials(t *testing.T) {

	var testCases = []struct {
		title    string
		input    *gosdk.Config
		expected string
	}{
		{
			title:    "TestCase 1: Should fail if BaseURL is empty.",
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
			title:    "TestCase 2: Should fail if config type is empty.",
			expected: "OAUTH needs appSecret, appId and refreshToken",
			input: &gosdk.Config{
				BaseUrl: "https://frascati.projectkeep.io",
				Credentials: gosdk.Credentials{
					Scope:        "",
					Type:         "",
					UserName:     "",
					Password:     "",
					AppID:        "",
					AppSecret:    "",
					Token:        "",
					RefreshToken: "",
					ExpiryTime:   0,
				},
			},
		},
		{
			title:    "TestCase 3: Should fail if type is BASIC and no username",
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
			title:    "TestCase 4: Should fail if type is BASIC and no password",
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			_, err := testCase.input.DominoAccess()
			if err.Error() != testCase.expected {
				t.Errorf("%s, Result: %s, Expected: %s", testCase.title, err.Error(), testCase.expected)
			}
		})
	}

}
