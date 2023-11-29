/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package types

import "time"

type DominoRestAccess struct {
	BaseUrl     string                 `json:"baseUrl"`
	Token       string                 `json:"token"`
	ExpiryTime  int                    `json:"expiryTime"`
	Credentials RestCredentials        `json:"updateCredentials"`
	AccessToken func() (string, error) `json:"accessToken"`
	Expiry      func() (int, error)    `json:"expiry"`
	Scope       func() string          `json:"scope"`
}

type RestCredentials struct {
	Scope        string `json:"scope"`
	Type         string `json:"type"`
	UserName     string `json:"userName"`
	Password     string `json:"passWord"`
	RefreshToken string `json:"refreshToken"`
	AppID        string `json:"appId"`
	AppSecret    string `json:"appSecret"`
}

type DominoRestAccessParams struct {
	BaseUrl     string          `json:"baseUrl"`
	Credentials RestCredentials `json:"credentials"`
}

type DominoRestAccessResponse struct {
	Bearer string `json:"bearer"`
	Claims struct {
		Iss   string   `json:"iss"`
		Sub   string   `json:"sub"`
		Iat   int      `json:"iat"`
		Exp   int      `json:"exp"`
		Aud   []string `json:"aud"`
		Cn    string   `json:"cn"`
		Scope string   `json:"scope"`
		Email string   `json:"email"`
	} `json:"claims"`
	Leeway     int       `json:"leeway"`
	ExpSeconds int       `json:"expSeconds"`
	IssueDate  time.Time `json:"issueDate"`
}
