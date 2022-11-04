package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go/types"
)

func TestLogout(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	client := autoconfirmClient

	// Create logged in user.
	email := randomEmail()
	password := randomString(10)
	session, err := client.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(err)

	// Logout.
	err = client.WithToken(session.AccessToken).Logout()
	require.NoError(err)

	// Attempt refresh - should fail.
	_, err = client.RefreshToken(session.RefreshToken)
	assert.Error(err)
}
