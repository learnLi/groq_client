package groq

import "testing"

func TestGerOrganizationId(t *testing.T) {
	client := NewBasicClient()
	sessionToken := "ywN9r5lAbI9HchNN6JsT2qQIfLrRodv4HT_4GjzlFJO-"
	account := NewAccount(sessionToken, "")

	token, err := GetSessionToken(client, account.SessionToken, "http://127.0.0.1:7990")
	if err != nil {
		return
	}

	organizationId, err := GerOrganizationId(client, token.Data.SessionJwt, "http://127.0.0.1:7990")
	if err != nil {
		return
	}
	t.Log(organizationId)
}
