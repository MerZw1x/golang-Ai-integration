# golang-Ai-integration

Клиент для работы с OpenRouter API на языке Go.


### 1. Установка зависимостей

```bash
go get github.com/joho/godotenv
```

### 2. Создай файл `.env` в корне проекта

```env
OPENROUTER_API_KEY=sk-or-...
OPENROUTER_MODEL=google/gemma-4-31b-it:free
OPENROUTER_BASE_URL=https://openrouter.ai/api/v1
```

### 3.1 Модель для запроса в OpenRouter

```go
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}
```
```Stream == true``` -> если мы хотим получать ответ раздельными словами(а не целиком)
```Stream == false``` -> если мы готовы немножко подождать и получить 

Temperature -> про креативность ответа(чем больше, тем креативнее ответ)

### 3.2 Необходимый формат POST запроса в OpenRouter

```bash
curl https://openrouter.ai/api/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENROUTER_API_KEY" \
  -d '{
  "model": "google/gemma-4-31b-it:free",
  "messages": [
    {
      "role": "user",
      "content": "How many r`s are in the word `strawberry?`"
    }
  ],
  "reasoning": {
    "enabled": true
  }
}
```

### 3.3 Пример кода

```go
	// Подготовка запросов
	messages := []models.Message{
		{
			Role:    "user",
			Content: "скажи что я молодец",
		},
	}

	// Отправляем запрос к модели
	response, err := orClient.ChatCompletion(model, messages)
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}

	// Выводим ответ
	if len(response.Choices) > 0 {
		fmt.Println("Ответ модели:")
		fmt.Println(response.Choices[0].Message.Content)
		fmt.Println(strings.Repeat("-", 60))
		fmt.Printf("Использовано токенов: %d\n", response.Usage.TotalTokens)
		fmt.Printf("Модель: %s\n", model)
	}
```

### 3.4 Пример с несколькими запросами

```go
messages := []models.Message{
	{Role: "system", Content: "Ты полезный и весёлый помощник."},
	{Role: "user",   Content: "Привет!"},
	{Role: "user",   Content: "Расскажи шутку про программистов"},
}
```

## Конфигурация

| Переменная                | Описание                           | Обязательно | Пример значения                     |
|---------------------------|------------------------------------|-------------|-------------------------------------|
| `OPENROUTER_API_KEY`      | API-ключ OpenRouter                | Да          | `sk-or-v1-...`                      |
| `OPENROUTER_MODEL`        | Модель для использования           | Да          | `google/gemma-4-31b-it:free`        |
| `OPENROUTER_BASE_URL`     | Базовый URL API                    | Нет         | `https://openrouter.ai/api/v1`      |


Базовый URL API можно вынести в const в самом коде

### Лучшие бесплатные модели на данный момент:

- **`google/gemma-4-31b-it:free`**
- `meta-llama/llama-4-maverick:free`
- `meta-llama/llama-4-scout:free`
- `deepseek/deepseek-chat-v3-0324:free`
- `qwen/qwen3-235b-a22b:free`


**Я использую:**
```env
OPENROUTER_MODEL=google/gemma-4-31b-it:free
```
