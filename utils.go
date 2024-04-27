package groq

import (
	"encoding/json"
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
		"event_id":       "event-id-d65b2524-ab2c-4cda-8259-e50231fc4a4d",
		"app_session_id": "app-session-id-9b98a045-49f6-4a32-8cf5-7850018b7148",
		"persistent_id":  "persistent-id-27df4aea-b5c2-4806-892d-c9280abbd7c8",
		"client_sent_at": getTime(),
		"timezone":       "Asia/Shanghai",
		"app":            map[string]string{"identifier": "groq.com"},
		"sdk":            map[string]string{"identifier": "Stytch.js Javascript SDK", "version": "4.5.3"},
	}
	marshal, err := json.Marshal(bas)
	if err != nil {
		return ""
	}
	return string(marshal)
}
