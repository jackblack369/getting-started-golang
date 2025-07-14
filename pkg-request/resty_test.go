package pkgrequest

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestResty(t *testing.T) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("https://github.com/dingodb/dingoadm/releases/download/latest/commit_id")

	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode())
	fmt.Printf("Response Body: %s\n", resp.Body())
	respString := string(resp.Body())
	if respString == "" {
		t.Fatalf("Response body is empty")
	}
	fmt.Printf("Response string: %s\n", respString)
}
