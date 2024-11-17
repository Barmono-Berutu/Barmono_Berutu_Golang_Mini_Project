package helper

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// ResponseAI generates a response using the generative AI API
func ResponseAI(ctx context.Context, question string) (string, error) {
	// Create the client using the API key from environment variables
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("AI_API_KEY")))
	if err != nil {
		return "", fmt.Errorf("failed to create AI client: %v", err)
	}
	defer client.Close()

	// Select the generative model
	model := client.GenerativeModel("gemini-1.5-flash")

	// Generate content using the AI model
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	var result strings.Builder

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result.WriteString(fmt.Sprintf("%s", part))
			}
		}
	}

	// Clean up the result (remove newlines)
	finalResult := strings.ReplaceAll(result.String(), "\n", "")

	return finalResult, nil
}
