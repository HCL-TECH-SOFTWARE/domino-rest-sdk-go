/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/pkg/utils"
)

// Pertains with domino rest authentication type.
// BASIC requires username and password.
// OAUTH requires application ID, application secret and refresh token.
const (
	BASIC = "BASIC"
	OAUTH = "OAUTH"
	TOKEN = "TOKEN"
)

// Config structure used as required parameters for getting new access instance
type Config struct {
	BaseUrl     string `json:"baseUrl"`
	Credentials `json:"credentials"`
}

// Credentials needed to access Domino REST API server. Required properties changes depending
// on type. If type is basic then username and password is needed, otherwise, if type is
// oauth, then application ID, application secret and refresh token is required.
type Credentials struct {
	Scope        string `json:"scope"`
	Type         string `json:"type"`
	UserName     string `json:"userName"`
	Password     string `json:"passWord"`
	AppID        string `json:"appId"`
	AppSecret    string `json:"appSecret"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ExpiryTime   int    `json:"expiryTime"`
}

// AccessMethods are list of public functions available.
// GetBaseUrl return string of baseUrl.
// GetExpiry returns token expiration in nano milliseconds.
// GetAccessToken returns jwt token.
type AccessMethods struct {
	GetBaseUrl     func() string
	GetExpiry      func() int
	GetAccessToken func() (token string, err error)
}

// AccessResponse is a response object from domino server.
// Contains all information about token generated by domino server.
type AccessResponse struct {
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

func (c *Config) DominoAccess() (*AccessMethods, error) {

	if len(c.BaseUrl) == 0 {
		return nil, errors.New("URL should not be empty.")
	}

	if c.Credentials.Type == BASIC {
		if len(c.Credentials.UserName) == 0 || len(c.Credentials.Password) == 0 {
			return nil, errors.New("BASIC authentication needs username and password.")
		}
	}

	if c.Credentials.Type == TOKEN {
		if len(c.Credentials.Token) == 0 {
			return nil, errors.New("TOKEN must not be empty.")
		}
	}

	if c.Credentials.Type == OAUTH {
		if len(c.Credentials.AppSecret) == 0 ||
			len(c.Credentials.AppID) == 0 ||
			len(c.Credentials.RefreshToken) == 0 {
			return nil, errors.New("OAUTH needs appSecret, appId and refreshToken")
		}
	}

	accessMethods := new(AccessMethods)
	accessMethods.GetExpiry = c.getExpiry
	accessMethods.GetAccessToken = c.getAccessToken
	accessMethods.GetBaseUrl = c.getBaseUrl

	return accessMethods, nil

}

func (c *Config) getExpiry() int {
	return c.ExpiryTime
}

func (c *Config) getAccessToken() (token string, err error) {

	var (
		data        *bytes.Buffer
		endpoint    string
		contentType string
	)

	if len(c.Token) > 0 && !utils.IsExpired(c.ExpiryTime) {
		return c.Token, nil
	}

	if c.Type == BASIC {
		endpoint = c.BaseUrl + "/api/v1/auth"
		contentType = "application/json"
		credentials := map[string]interface{}{
			"username": c.UserName,
			"password": c.Password,
			"scope":    c.Scope,
		}

		jsonData, jsonErr := json.Marshal(credentials)
		if jsonErr != nil {
			return "", jsonErr
		}
		data = bytes.NewBuffer(jsonData)
	}

	if c.Type == TOKEN {
		endpoint = c.BaseUrl + "/api/v1/auth"
		contentType = "application/json"
		credentials := map[string]interface{}{
			"type":  c.Type,
			"token": c.Token,
		}

		jsonData, jsonErr := json.Marshal(credentials)
		if jsonErr != nil {
			return "", jsonErr
		}
		data = bytes.NewBuffer(jsonData)
	}

	if c.Type == OAUTH {
		endpoint = c.BaseUrl + "/oauth/token"
		contentType = "application/x-www-form-urlencoded"
		credentials := map[string]interface{}{
			"grant_type":    "refresh_token",
			"refresh_token": c.RefreshToken,
			"scope":         c.Scope,
			"client_id":     c.AppID,
			"client_secret": c.AppSecret,
		}

		jsonData, jsonErr := json.Marshal(credentials)
		if jsonErr != nil {
			return "", jsonErr
		}
		data = bytes.NewBuffer(jsonData)
	}

	reqParams := new(utils.RequestParameters)
	reqParams.Method = http.MethodPost
	reqParams.Url = endpoint
	reqParams.Body = data
	reqParams.Header = map[string]string{
		"Content-Type": contentType,
	}

	response, respError := reqParams.Request()
	if respError != nil {
		return "", respError
	}

	responseData := new(AccessResponse)
	respDataErr := json.Unmarshal(response, responseData)
	if respDataErr != nil {
		return "", respDataErr
	}
	c.ExpiryTime = responseData.Claims.Exp

	return responseData.Bearer, nil
}

func (c *Config) getBaseUrl() string {
	return c.BaseUrl
}
