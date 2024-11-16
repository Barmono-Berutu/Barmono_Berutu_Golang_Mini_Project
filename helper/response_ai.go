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
	apiKey := os.Getenv("AI_API_KEY")
	if apiKey == "" {
		log.Fatal("API Key is missing")
		return "", fmt.Errorf("API Key is missing")
	}

	httpClient := &http.Client{}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey), option.WithHTTPClient(httpClient))
	if err != nil {
		log.Printf("Error creating client: %v", err)
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		return "", err
	}

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
