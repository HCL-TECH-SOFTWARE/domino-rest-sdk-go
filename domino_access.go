package keep_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	c "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/constants"
	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseAccess = new(BaseDominoAccess)

type BaseDominoAccess struct {
	t.DominoRestAccess
}

func DominoAccess(args t.DominoRestAccessParams) (t.DominoRestAccess, error) {

	if args.BaseUrl == "" {
		return t.DominoRestAccess{}, errors.New("URL should not be empty")
	}

	if args.Credentials.Type == c.BASIC {
		if args.Credentials.UserName == "" || args.Credentials.Password == "" {
			return t.DominoRestAccess{}, errors.New("BASIC authentication needs username and password")
		}
	} else {
		if args.Credentials.AppSecret == "" || args.Credentials.AppID == "" || args.Credentials.RefreshToken == "" {
			return t.DominoRestAccess{}, errors.New("OAUTH needs appSecret, appId and refreshToken")
		}
	}

	baseAccess.BaseUrl = args.BaseUrl
	baseAccess.Credentials = args.Credentials
	baseAccess.AccessToken = accessToken

	token, err := baseAccess.AccessToken()
	if err != nil {
		return t.DominoRestAccess{}, err
	}

	baseAccess.Token = token
	baseAccess.Expiry = expiry

	return baseAccess.DominoRestAccess, nil
}

func accessToken() (string, error) {
	var (
		data        *bytes.Buffer
		endpoint    string
		contentType string
	)

	if baseAccess.Token != "" && !h.IsJWTExpired(baseAccess.ExpiryTime) {
		return baseAccess.Token, nil
	}

	if baseAccess.Credentials.Type == c.BASIC {
		endpoint = baseAccess.BaseUrl + "/api/v1/auth"
		contentType = "application/json"
		jsonData, _ := json.Marshal(map[string]interface{}{
			"username": baseAccess.Credentials.UserName,
			"password": baseAccess.Credentials.Password,
			"scope":    baseAccess.Credentials.Scope,
		})
		data = bytes.NewBuffer(jsonData)
	} else {
		endpoint = baseAccess.BaseUrl + "/oauth/token"
		contentType = "application/x-www-form-urlencoded"
		jsonData, _ := json.Marshal(map[string]interface{}{
			"grant_type":    "refresh_token",
			"refresh_token": baseAccess.Credentials.RefreshToken,
			"scope":         baseAccess.Credentials.Scope,
			"client_id":     baseAccess.Credentials.AppID,
			"client_secret": baseAccess.Credentials.AppSecret,
		})
		data = bytes.NewBuffer(jsonData)
	}
	req := h.RequestParams{}
	req.Method = http.MethodPost
	req.Url = endpoint
	req.Body = data
	req.Header = map[string]string{"Content-type": contentType}

	response, err := h.Request(req)
	if err != nil {
		return "", err
	}

	result := t.DominoRestAccessResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	baseAccess.Token = result.Bearer
	baseAccess.ExpiryTime = result.Claims.Exp

	return baseAccess.Token, nil
}

func expiry() (int, error) {
	if baseAccess.ExpiryTime <= 0 {
		return 0, errors.New("No token with expiry found")
	}

	return baseAccess.ExpiryTime, nil
}
