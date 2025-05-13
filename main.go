package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	
	"github.com/evgeniya/tgbot/utils"
)

func main() {
	// Get bot token from environment variable
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("No BOT_TOKEN provided")
	}

	// Create bot instance
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("Failed to create new bot: ", err)
	}

	// Create dispatcher and updater
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// Log errors when panics happen in handlers
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Printf("Error handling update: %s", err)
			return ext.DispatcherActionNoop
		},
	})
	
	updater := ext.NewUpdater(dispatcher, &ext.UpdaterOpts{})
	
	// Simplified handler approach - just two handlers
	
	// 1. Command handler for /story - this must be first
	dispatcher.AddHandler(handlers.NewMessage(
		func(msg *gotgbot.Message) bool {
			if msg.Text == "" {
				return false
			}
			
			isStoryCommand := msg.Text == "/story" || strings.HasPrefix(msg.Text, "/story ")
			if isStoryCommand {
				log.Printf("Story command detected: %s", msg.Text)
				return true
			}
			return false
		},
		storyCommandHandler,
	))
	log.Println("Registered /story command handler")
	
	// 2. Handler for all other text messages (non-commands)
	dispatcher.AddHandler(handlers.NewMessage(
		func(msg *gotgbot.Message) bool {
			if msg.Text == "" || strings.HasPrefix(msg.Text, "/") {
				return false
			}
			log.Printf("Text message detected: %s", msg.Text)
			return true
		},
		invertMessageHandler,
	))
	log.Println("Registered text message handler")

	// Start the updater
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
	})
	if err != nil {
		log.Fatal("Failed to start polling: ", err)
	}

	log.Printf("Bot started as @%s", bot.User.Username)

	// Keep the program running
	updater.Idle()
}

// invertMessageHandler handles text messages and replies with inverted text
func invertMessageHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	// Get the message text
	text := ctx.EffectiveMessage.Text
	log.Printf("Inverting message: '%s'", text)

	// Invert the text using the utility function
	invertedText := utils.InvertString(text)

	// Reply to the message with the inverted text
	_, err := ctx.EffectiveMessage.Reply(bot, invertedText, nil)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// storyCommandHandler handles the /story command and generates a sci-fi story
func storyCommandHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	log.Printf("Story command triggered by user %d (%s)", ctx.EffectiveUser.Id, ctx.EffectiveUser.Username)
	
	// Send a temporary message while generating the story
	tempMsg, err := ctx.EffectiveMessage.Reply(bot, "Generating a sci-fi story for you...", nil)
	if err != nil {
		log.Printf("Failed to send temporary message: %v", err)
		// Try to send a direct message if reply fails
		_, err = ctx.EffectiveMessage.Reply(bot, "Error: Unable to generate story at this time.", nil)
		return fmt.Errorf("failed to send temporary message: %w", err)
	}
	
	log.Println("Calling GenerateSciFiStory function...")
	// Generate the sci-fi story - this now has a fallback mechanism
	story, err := utils.GenerateSciFiStory()
	if err != nil {
		log.Printf("Failed to generate story: %v", err)
		// Update the temporary message with the error
		_, _, err2 := bot.EditMessageText("Sorry, failed to generate story. Please try again later.", &gotgbot.EditMessageTextOpts{
			ChatId:    tempMsg.Chat.Id,
			MessageId: tempMsg.MessageId,
		})
		if err2 != nil {
			log.Printf("Failed to edit message: %v", err2)
			// Try to send a direct message if editing fails
			_, _ = ctx.EffectiveMessage.Reply(bot, "Sorry, failed to generate story. Please try again later.", nil)
		}
		return fmt.Errorf("failed to generate story: %w", err)
	}
	
	log.Printf("Story generated successfully, length: %d chars", len(story))
	// Update the temporary message with the generated story
	_, _, err = bot.EditMessageText(story, &gotgbot.EditMessageTextOpts{
		ChatId:    tempMsg.Chat.Id,
		MessageId: tempMsg.MessageId,
	})
	if err != nil {
		log.Printf("Failed to edit message: %v", err)
		// Try to send a direct message if editing fails
		_, _ = ctx.EffectiveMessage.Reply(bot, story, nil)
		return fmt.Errorf("failed to edit message: %w", err)
	}
	
	log.Println("Story command completed successfully")
	return nil
} 