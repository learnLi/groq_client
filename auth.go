package groq

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func LoginOrCreate(client HTTPClient, email string, proxy string) (*http.Response, error) {
	rawUrl := "https://web.stytch.com/sdk/v1/magic_links/email/login_or_create"
	if proxy != "" {
		client.SetProxy(proxy)
	}
	dfpTelemetryId, err := Submit(client, "")
	if err != nil {
		return nil, err
	}
	header := baseHeader()
	header.Set("authorization", "Basic "+BasicAuth)
	header.Set("x-sdk-client", generateSdkClient())
	header.Set("x-sdk-parent-host", "https://groq.com")
	jsonData := map[string]interface{}{
		"signup_magic_link_url":     "https://groq.com/authenticate",
		"signup_expiration_minutes": 60,
		"login_magic_link_url":      "https://groq.com/authenticate",
		"login_expiration_minutes":  60,
		"email":                     email,
		"dfp_telemetry_id":          dfpTelemetryId,
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
	defer req.Body.Close()
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
