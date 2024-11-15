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
		log.Println("Error creating AI client:", err)
		return "", err
	}

	// Inisialisasi model generatif
	modelAI := client.GenerativeModel("gemini-1.5-flash")
	modelAI.SetTemperature(0) // Atur suhu ke 0 untuk respons deterministik

	// Panggil API untuk menghasilkan konten
	resp, err := modelAI.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Println("Error generating content:", err)
		return "", err
	}

	// Pastikan respons memiliki kandidat
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		log.Println("No response from AI model")
		return "", fmt.Errorf("no response from AI model")
	}

	// Ambil respons pertama dari kandidat
	answer := resp.Candidates[0].Content.Parts[0]
	answerString := fmt.Sprintf("%v", answer)

	// Bersihkan simbol tambahan dari teks
	answerString = strings.ReplaceAll(answerString, "*", "")
	answerString = strings.ReplaceAll(answerString, "**", "")
	answerString = strings.ReplaceAll(answerString, "\n\n", " -")

	return answerString, nil
}
