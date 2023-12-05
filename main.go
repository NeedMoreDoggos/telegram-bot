package main

import (
	"flag"
	"log"

	"github.com/NeedMoreDoggos/telegram-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	token := mustToken()
	tgClient = telegram.New(tgBotHost, token)
	fetcher = fetcher.New()
	processor = processor.New()
	//consumer.Start(fetcher, processor)
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
