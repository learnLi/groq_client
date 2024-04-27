package groq

import (
	"fmt"
	"testing"
)

func TestSdkClient(t *testing.T) {
	str := generateSdkClient()
	fmt.Println(str)
}
