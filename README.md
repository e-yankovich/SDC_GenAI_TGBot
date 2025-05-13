# Telegram Message Inverter Bot

A simple Telegram bot that inverts any text message sent to it and generates sci-fi stories.

## Features

- Replies to any text message with its reversed version
- For example, if you send "hello", it will reply with "olleh"
- Generate sci-fi stories with the `/story` command (uses OpenAI API)

## Setup

1. Install Go (version 1.21 or later recommended)
2. Clone this repository
3. Get a Telegram Bot token from [@BotFather](https://t.me/BotFather)
4. Get an OpenAI API key from [OpenAI Platform](https://platform.openai.com/)
5. Run the setup script to download dependencies and generate go.sum:
   ```bash
   ./setup.sh
   ```
6. Set the environment variables:
   ```bash
   export BOT_TOKEN=your_telegram_token_here
   export OPENAI_API_KEY=your_openai_api_key_here
   ```
7. Run the bot:
   ```bash
   go run main.go
   ```

## Building and Running

To build the executable:

```bash
go build -o telegram-inverter-bot
```

To run the compiled executable:

```bash
BOT_TOKEN=your_telegram_token_here OPENAI_API_KEY=your_openai_api_key_here ./telegram-inverter-bot
```

## Project Structure

- `main.go` - Main application file with bot initialization and message handling
- `utils/string_utils.go` - Utility functions for string manipulation, including the message inversion logic
- `utils/openai.go` - Integration with OpenAI API for generating sci-fi stories

## Usage

1. Start a chat with your bot on Telegram
2. Send any text message to get its reversed version
3. Send `/story` command to get a generated sci-fi micro-story 