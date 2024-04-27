package groq

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type MagicLinkRequest struct {
	SignupMagicLinkURL      string `json:"signup_magic_link_url"`
	SignupExpirationMinutes int    `json:"signup_expiration_minutes"`
	LoginMagicLinkURL       string `json:"login_magic_link_url"`
	LoginExpirationMinutes  int    `json:"login_expiration_minutes"`
	Email                   string `json:"email"`
	DfpTelemetryID          string `json:"dfp_telemetry_id"`
}

func LoginOrCreate(client HTTPClient, email string, proxy string) (*http.Response, error) {
	rawUrl := "https://web.stytch.com/sdk/v1/magic_links/email/login_or_create"
	if proxy != "" {
		client.SetProxy(proxy)
	}
	dfpTelemetryId, err := Submit(client, proxy)
	if err != nil {
		return nil, err
	}
	header := baseHeader()
	header.Set("authorization", "Basic "+BasicAuth)
	header.Set("x-sdk-client", generateSdkClient())
	header.Set("x-sdk-parent-host", "https://groq.com")
	requestPayload := MagicLinkRequest{
		SignupMagicLinkURL:      "https://groq.com/authenticate",
		SignupExpirationMinutes: 60,
		LoginMagicLinkURL:       "https://groq.com/authenticate",
		LoginExpirationMinutes:  60,
		Email:                   email,
		DfpTelemetryID:          dfpTelemetryId,
	}
	data, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	req, err := client.Request("POST", rawUrl, header, nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if req.StatusCode != 200 {
		return nil, errors.New("login or create failed")
	}
	return req, nil
}

func LoginOrCreateCallback(client HTTPClient, token string, proxy string) (*http.Response, error) {
	if proxy != "" {
		client.SetProxy(proxy)
	}
	header := baseHeader()
	header.Set("authorization", "Basic "+BasicAuth)
	header.Set("x-sdk-client", generateSdkClient())
	header.Set("x-sdk-parent-host", "https://groq.com")
	rawUrl := "https://web.stytch.com/sdk/v1/magic_links/authenticate"
	jsonData := map[string]interface{}{
		"session_duration_minutes": 43200,
		"token":                    token,
	}
	data, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	req, err := client.Request("POST", rawUrl, header, nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if req.StatusCode != 200 {
		return nil, errors.New("login or create callback failed")
	}
	return req, nil
}

type CallbackResponse struct {
	Data struct {
		RequestId    string `json:"request_id"`
		SessionJwt   string `json:"session_jwt"`
		SessionToken string `json:"session_token"`
		StatusCode   int    `json:"status_code"`
	} `json:"data"`
}
