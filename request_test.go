package groq

import "testing"

func TestGerOrganizationId(t *testing.T) {
	client := NewBasicClient()
	sessionToken := "3WbHxt-tzOVjM5jhMCJA-c60WR9CdndB0ar2uuz_4oI5"
	account := NewAccount(sessionToken, "")

	token, err := GetSessionToken(client, account.SessionToken, "")
	if err != nil {
		return
	}

	organizationId, err := GerOrganizationId(client, token.Data.SessionJwt, "")
	if err != nil {
		return
	}
	t.Log(organizationId)
}
