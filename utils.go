package groq

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"
)

func getTime() string {
	// Get the current UTC time
	currentUTCTime := time.Now().UTC()

	// Format the time to ISO 8601 with milliseconds and 'Z' suffix
	timestamp := currentUTCTime.Format("2006-01-02T15:04:05.000Z")
	return timestamp
}

func generateSdkClient() string {
	bas := map[string]interface{}{
		"event_id":       "event-id-ef768192-d460-46ee-b293-d7d42e4bec2e",
		"app_session_id": "app-session-id-cfc2e7d2-954f-4baf-8b71-a911c5d84c9c",
		"persistent_id":  "persistent-id-6327df6e-f4ca-4b5b-8ce8-24b710bb3d1b",
		"client_sent_at": getTime(),
		"timezone":       "Asia/Shanghai",
		"app":            map[string]string{"identifier": "groq.com"},
		"sdk":            map[string]string{"identifier": "Stytch.js Javascript SDK", "version": "4.6.0"},
	}
	marshal, err := json.Marshal(bas)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(marshal)
}

func RandomHexadecimalString(length int) string {
	if length < 1 {
		length = 4
	}
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789abcdef"
	// const length = 16 // The length of the string you want to generate
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
