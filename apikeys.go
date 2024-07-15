package groq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type APIKEYS struct {
	Object string   `json:"object"`
	Data   []APIKEY `json:"data"`
}

type APIKEY struct {
	Object           string        `json:"object"`
	Id               string        `json:"id"`
	OrgId            string        `json:"org_id"`
	Name             string        `json:"name"`
	SecretKey        string        `json:"secret_key"`
	Created          int64         `json:"created"`
	LastUse          int           `json:"last_use"`
	CreatedBy        string        `json:"created_by"`
	Scopes           []interface{} `json:"scopes"`
	ExposedSecretKey string        `json:"exposed_secret_key"`
}

func GetAPIKEYSLIST(client HTTPClient, api_key string, organization string, proxy string) (*APIKEYS, error) {
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	header.Set("referer", "https://console.groq.com/")
	header.Set("origin", "https://console.groq.com")
	header.Set("groq-organization", organization)
	if proxy != "" {
		client.SetProxy(proxy)
	}

	resp, err := client.Request("GET", fmt.Sprintf("https://api.groq.com/platform/v1/organizations/%s/api_keys", organization), header, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("response status code is not 200")
	}
	result := new(APIKEYS)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteAPIKEY(client HTTPClient, api_key string, organization string, api_key_id string, proxy string) error {
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	header.Set("referer", "https://console.groq.com/")
	header.Set("origin", "https://console.groq.com")
	header.Set("groq-organization", organization)
	if proxy != "" {
		client.SetProxy(proxy)
	}
	if api_key_id == "" {
		return errors.New("api_key_id is empty")
	}
	resp, err := client.Request("DELETE", fmt.Sprintf("https://api.groq.com/platform/v1/organizations/%s/api_keys/%s", organization, api_key_id), header, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}
	return nil
}

func GenerateAPIKEY(client HTTPClient, api_key string, organization string, name string, proxy string) (*APIKEY, error) {
	header := baseHeader()
	header.Set("authorization", "Bearer "+api_key)
	header.Set("referer", "https://console.groq.com/")
	header.Set("origin", "https://console.groq.com")
	header.Set("groq-organization", organization)
	if proxy != "" {
		client.SetProxy(proxy)
	}

	newName := RandomHexadecimalString(10) + name
	jsonData := map[string]string{
		"name": newName,
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	resp, err := client.Request("POST", fmt.Sprintf("https://api.groq.com/platform/v1/organizations/%s/api_keys", organization), header, nil, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("response status code is not 200")
	}
	var result APIKEY
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
