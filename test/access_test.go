package test

import (
	"testing"

	sdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
	"github.com/stretchr/testify/assert"
)

// "baseUrl": "https://frascati.projectkeep.io",
//
//	"credentials": {
//	  "scope": "$DATA",
//	  "type": "BASIC",
//	  "username": "Doctor Notes",
//	  "password": "lotusnotes"
//	}

func TestDominoAccess(t *testing.T) {

	Env.BaseURL = "https://frascati.projectkeep.io"
	Env.Credentials.Type = "BASIC"
	Env.Credentials.Scope = "$DATA"
	Env.Credentials.Username = "Doctor Notes"
	Env.Credentials.Password = "lotusnotes"

	drp := types.DominoRestAccessParams{}
	drp.BaseUrl = Env.BaseURL
	drp.Credentials.Type = Env.Credentials.Type
	drp.Credentials.UserName = Env.Credentials.Username
	drp.Credentials.Password = Env.Credentials.Password

	access, err := sdk.DominoAccess(drp)
	token, tokenErr := access.AccessToken()
	expiry, expiryErr := access.Expiry()

	t.Run("OnSuccess: DominoAccess error should return nil", func(t *testing.T) {
		assert.Nil(t, err)
	})

	t.Run("OnSuccess: Token error should return nil", func(t *testing.T) {
		assert.Nil(t, tokenErr)
	})

	t.Run("OnSuccess: Token should not be empty", func(t *testing.T) {
		assert.NotEmpty(t, token)
	})

	t.Run("OnSuccess: Expiry error should return nil", func(t *testing.T) {
		assert.Nil(t, expiryErr)
	})

	t.Run("OnSuccess: Expiry should return non zero value", func(t *testing.T) {

		isZero := func() bool {
			if expiry > 0 {
				return true
			}
			return false
		}
		assert.Condition(t, isZero)

	})

}
