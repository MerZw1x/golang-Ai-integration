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

### 3.1 Выделяем модели

### 3.2 Пример кода

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

### 3.3 Пример с несколькими запросами

```go
messages := []models.Message{
	{Role: "system", Content: "Ты полезный и весёлый помощник."},
	{Role: "user",   Content: "Привет!"},
	{Role: "user",   Content: "Расскажи шутку про программистов"},
}
```

## Конфигурация (.env)

| Переменная                | Описание                           | Обязательно | Пример значения                     |
|---------------------------|------------------------------------|-------------|-------------------------------------|
| `OPENROUTER_API_KEY`      | API-ключ OpenRouter                | Да          | `sk-or-v1-...`                      |
| `OPENROUTER_MODEL`        | Модель для использования           | Да          | `google/gemma-4-31b-it:free`        |
| `OPENROUTER_BASE_URL`     | Базовый URL API                    | Нет         | `https://openrouter.ai/api/v1`      |


Базовый URL API можно вынести в const в самом коде

### Лучшие бесплатные модели на данный момент:

- **`google/gemma-4-31b-it:free`** ← **твоя текущая модель** (рекомендуется)
- `meta-llama/llama-4-maverick:free`
- `meta-llama/llama-4-scout:free`
- `deepseek/deepseek-chat-v3-0324:free`
- `qwen/qwen3-235b-a22b:free`


**Я использую:**
```env
OPENROUTER_MODEL=google/gemma-4-31b-it:free
```
