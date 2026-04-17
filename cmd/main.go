package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"openrouter-integration/internal/client"
	"openrouter-integration/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Api-key
	openrouterKey := os.Getenv("OPENROUTER_API_KEY")

	// LLM-model
	model := os.Getenv("OPENROUTER_MODEL")

	// URL of openRouter website
	baseURL := os.Getenv("OPENROUTER_BASE_URL")

	orClient := client.NewOpenRouterClient(openrouterKey, baseURL)

	messages := []models.Message{
		{
			Role:    "user",
			Content: "Что будет через 10 лет в мире?",
		},
	}

	response, err := orClient.ChatCompletion(model, messages)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if len(response.Choices) > 0 {
		fmt.Print("Ответ: ")
		fmt.Println(response.Choices[0].Message.Content)
		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("Статистика: %d токенов использовано\n", response.Usage.TotalTokens)
	}
}
