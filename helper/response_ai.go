package helper

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ResponseAI(ctx context.Context, question string) (string, error) {
	// Create a custom HTTP client with insecure transport
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDAsFVvWobfb1sWT-fZE1FJVyLyq-9kh0k"), option.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		log.Println("No response from AI model")
		return "", fmt.Errorf("no response from AI model")
	}

	answer := resp.Candidates[0].Content.Parts[0]
	answerString := fmt.Sprintf("%v", answer)

	answerString = strings.ReplaceAll(answerString, "*", "")
	answerString = strings.ReplaceAll(answerString, "**", "")
	answerString = strings.ReplaceAll(answerString, "\n\n", " -")

	return answerString, nil
}
