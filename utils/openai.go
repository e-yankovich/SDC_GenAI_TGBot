package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// GenerateSciFiStory generates a short sci-fi story using OpenAI API
func GenerateSciFiStory() (string, error) {
	// Get OpenAI API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Println("OPENAI_API_KEY environment variable not set")
		return getFallbackStory(), nil
	}
	
	log.Printf("Found OpenAI API key: %s...%s", apiKey[:10], apiKey[len(apiKey)-5:])
	
	// Define the request payload
	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a creative sci-fi writer. Keep responses under 400 characters.",
			},
			{
				"role":    "user",
				"content": "Write a creative sci-fi micro-story in exactly 400 characters. Make it engaging with a beginning, middle and end.",
			},
		},
		"max_tokens": 150,
	}
	
	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return getFallbackStory(), nil
	}
	
	log.Printf("Request payload: %s", string(jsonPayload))
	
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return getFallbackStory(), nil
	}
	
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
	// Create a client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// Send the request
	log.Println("Sending request to OpenAI API...")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return getFallbackStory(), nil
	}
	defer resp.Body.Close()
	
	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return getFallbackStory(), nil
	}
	
	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", string(body))
	
	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		log.Printf("OpenAI API error: %s", resp.Status)
		return getFallbackStory(), nil
	}
	
	// Parse the response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		return getFallbackStory(), nil
	}
	
	// Extract the story from the response
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		log.Println("No choices in response")
		return getFallbackStory(), nil
	}
	
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		log.Println("Invalid choice format")
		return getFallbackStory(), nil
	}
	
	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		log.Println("Invalid message format")
		return getFallbackStory(), nil
	}
	
	content, ok := message["content"].(string)
	if !ok {
		log.Println("Invalid content format")
		return getFallbackStory(), nil
	}
	
	log.Printf("Generated story: %s", content)
	return content, nil
}

// getFallbackStory returns a hardcoded story when OpenAI API fails
func getFallbackStory() string {
	log.Println("Using fallback story")
	return "THE LAST MESSAGE: In 2157, Earth received a cryptic signal from deep space. Scientists scrambled to decode it as strange phenomena plagued the planet. When they finally understood, it was too late. The message was a timer, counting down to something inevitable. Humanity had received its eviction notice."
} 