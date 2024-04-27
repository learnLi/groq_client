package groq

import (
	"fmt"
	"testing"
)

func TestTelemetry(t *testing.T) {
	client := NewBasicClient()
	submit, err := Submit(client, "http://127.0.0.1:7990")
	if err != nil {
		return
	}

	fmt.Println(submit)
}
