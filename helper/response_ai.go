package helper

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ResponseAI(ctx context.Context, question string) (string, error) {
	// Get the API key from environment variable
	apiKey := os.Getenv("AI_API_KEY")
	if apiKey == "" {
		log.Println("API Key is missing") // Ganti log.Fatal dengan log.Println untuk debugging
		return "", fmt.Errorf("API Key is missing")
	}

	log.Println("API Key successfully retrieved")

	// Create default HTTP client
	httpClient := &http.Client{}

	// Create a new GenAI client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey), option.WithHTTPClient(httpClient))
	if err != nil {
		log.Printf("Error creating client: %v", err)
		return "", err
	}
	defer client.Close()

	log.Println("Client created successfully, calling the model")

	// Call the model
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		return "", err
	}

	log.Println("Model response received")

	// Process the response
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		log.Println("No response from AI model")
		return "", fmt.Errorf("no response from AI model")
	}

	var answerString string
	for _, part := range resp.Candidates[0].Content.Parts {
		answerString += fmt.Sprintf("%v", part)
	}

	// Clean the response string
	answerString = strings.ReplaceAll(answerString, "*", "")
	answerString = strings.ReplaceAll(answerString, "**", "")
	answerString = strings.ReplaceAll(answerString, "\n\n", " -")

	return answerString, nil
}
