package groq

import (
	"errors"
	"io"
	"strings"
)

func Submit(client HTTPClient, proxy string) (string, error) {
	if proxy != "" {
		client.SetProxy(proxy)
	}
	header := baseHeader()
	header.Set("content-type", "application/x-www-form-urlencoded")
	rawUrl := "https://telemetry.stytch.com/submit"
	response, err := client.Request("POST", rawUrl, header, nil, strings.NewReader(SubmitPayload))
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New("fetch telemetry id failed")
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}
