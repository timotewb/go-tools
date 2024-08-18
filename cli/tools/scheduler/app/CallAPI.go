package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func CallAPI() {
	url := "http://10.0.0.200:80/api"

	// Prepare the JSON payload
	jsonData := map[string]interface{}{
		"name": "hello_world",
	}
	jsonPayload, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Err: %v", err)
	}

	// Create a new request using http.Post
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error creating request. Err: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "text/plain")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request. Err: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body. Err: %v", err)
	}

	fmt.Println(string(body))
}