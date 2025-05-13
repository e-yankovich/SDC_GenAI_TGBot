#!/bin/bash

# Download dependencies and generate go.sum
go mod download

echo "Dependencies downloaded and go.sum generated"
echo "To run the bot: BOT_TOKEN=your_telegram_token_here OPENAI_API_KEY=your_openai_api_key_here go run main.go" 