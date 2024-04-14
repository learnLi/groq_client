package groq

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func ChatCompletions(client HTTPClient, api_request APIRequest, api_key string, organization string, proxy string) (*http.Response, error) {
	if proxy != "" {
		client.SetProxy(proxy)
	}
	body_json, _ := json.Marshal(api_request)
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	header.Set("groq-app", "chat")
	header.Set("groq-organization", organization)
	response, err := client.Request("POST", "https://api.groq.com/openai/v1/chat/completions", header, nil, bytes.NewBuffer(body_json))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {

		return nil, errors.New("response status code is not 200")
	}
	return response, nil
}

func GetSessionToken(client HTTPClient, api_key string, proxy string) (AuthenticateResponse, error) {
	if proxy != "" {
		client.SetProxy(proxy)
	}
	if api_key == "" {
		return AuthenticateResponse{}, errors.New("session token is empty")
	}
	authorization := generateRefreshToken(api_key)
	header := baseHeader()
	header.Set("authorization", "Basic "+authorization)
	header.Set("x-sdk-client", "eyJldmVudF9pZCI6ImV2ZW50LWlkLWY1ZDJiYTRkLWYwYzItNDI4OC05ZWIwLWE2ZmJiNTc5ZjIxYyIsImFwcF9zZXNzaW9uX2lkIjoiYXBwLXNlc3Npb24taWQtODc4MzY0OGQtYTg2ZC00ZDBkLTlmODMtMmIyZGE4ZDAyMjY0IiwicGVyc2lzdGVudF9pZCI6InBlcnNpc3RlbnQtaWQtOGE1MTg3NjEtZDc0Ni00NjY3LWE3OGEtNjIyM2Q1M2NkMTkzIiwiY2xpZW50X3NlbnRfYXQiOiIyMDI0LTA0LTEzVDExOjA4OjMwLjg4NVoiLCJ0aW1lem9uZSI6IkFzaWEvU2hhbmdoYWkiLCJzdHl0Y2hfdXNlcl9pZCI6InVzZXItbGl2ZS00ZWIzZTA1Mi0zMGFmLTQzZWItOGM0Yi02YmQzMmE4YzhlMWMiLCJzdHl0Y2hfc2Vzc2lvbl9pZCI6InNlc3Npb24tbGl2ZS1hMWYwYzdhYy01NDkwLTQ0YjItYmU3MS1kMjIwOTVjMDU5NmIiLCJhcHAiOnsiaWRlbnRpZmllciI6Imdyb3EuY29tIn0sInNkayI6eyJpZGVudGlmaWVyIjoiU3R5dGNoLmpzIEphdmFzY3JpcHQgU0RLIiwidmVyc2lvbiI6IjQuNS4zIn19")
	header.Set("x-sdk-parent-host", "https://groq.com")

	rawUrl := "https://web.stytch.com/sdk/v1/sessions/authenticate"
	req, err := client.Request("POST", rawUrl, header, nil, strings.NewReader(`{}`))
	if err != nil {
		return AuthenticateResponse{}, err
	}
	if req.StatusCode != 200 {
		return AuthenticateResponse{}, errors.New("authenticate failed")
	}
	var result AuthenticateResponse
	err = json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		return AuthenticateResponse{}, err
	}
	return result, nil
}

func GetModels(client HTTPClient, api_key string, proxy string) (*http.Response, error) {
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	if proxy != "" {
		client.SetProxy(proxy)
	}
	response, err := client.Request("GET", "https://api.groq.com/openai/v1/models", header, nil, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("response status code is not 200")
	}
	return response, nil
}

func GerOrganizationId(client HTTPClient, api_key string, proxy string) (string, error) {
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	if proxy != "" {
		client.SetProxy(proxy)
	}
	response, err := client.Request("GET", "https://api.groq.com/platform/v1/user/profile", header, nil, nil)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New("response status code is not 200")
	}
	var result Profile
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.User.Orgs.Data[0].Id, nil
}

func baseHeader() Headers {
	header := NewHeader()
	header.Set("accept", "*/*")
	header.Set("accept-language", "zh-CN,zh;q=0.9")
	header.Set("content-type", "application/json")
	header.Set("origin", "https://groq.com")
	header.Set("referer", "https://groq.com/")
	header.Set("sec-ch-ua", `"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	header.Set("sec-ch-ua-mobile", "?0")
	header.Set("sec-ch-ua-platform", `"Windows"`)
	header.Set("sec-fetch-dest", "empty")
	header.Set("sec-fetch-mode", "cors")
	header.Set("sec-fetch-site", "cross-site")
	header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	return header
}

func generateRefreshToken(api_key string) string {
	prefix := "public-token-live-26a89f59-09f8-48be-91ff-ce70e6000cb5:" + api_key
	return base64.StdEncoding.EncodeToString([]byte(prefix))
}
