package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nmyshenkov/private-telegram-bot/config"
)

func main() {
	configPath := flag.String("c", "config_local.json", "config file")
	flag.Parse()

	cfg := new(config.Config)

	if err := cfg.ParseFromFile(*configPath); err != nil {
		log.Fatalf("parse config file: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("validate config file: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = cfg.Debug
	if cfg.Debug {
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			if cfg.Debug {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			}

			msg := tgbotapi.NewMessage(cfg.TargetChatID, update.Message.Text)

			if _, err := bot.Send(msg); err != nil {
				log.Printf("error send message to channel: %v", err)
			}
		}
	}
}
