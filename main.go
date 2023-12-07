package main

import (
	"flag"
	"log"

	"github.com/NeedMoreDoggos/telegram-bot/clients/telegram"
	event_consumer "github.com/NeedMoreDoggos/telegram-bot/consumer/event-consumer"
	telegram2 "github.com/NeedMoreDoggos/telegram-bot/events/telegram"
	"github.com/NeedMoreDoggos/telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram2.New(
		telegram.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"token-bot",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
