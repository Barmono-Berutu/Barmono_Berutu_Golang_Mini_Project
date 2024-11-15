package helper

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ResponseAI(ctx context.Context, question string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDH_LO9LZulVt6yxB7j0pNel8N5U8Adp7o"))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	modelAI := client.GenerativeModel("gemini-1.5-flash")
	modelAI.SetTemperature(0)

	resp, err := modelAI.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	answer := resp.Candidates[0].Content.Parts[0]
	answerString := fmt.Sprintf("%v", answer)

	// Bersihkan simbol tambahan dari teks
	answerString = strings.ReplaceAll(answerString, "*", "")
	answerString = strings.ReplaceAll(answerString, "**", "")
	answerString = strings.ReplaceAll(answerString, "\n\n", " -")
	return answerString, nil
}
