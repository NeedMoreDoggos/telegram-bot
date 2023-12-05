package telegram

import (
	"log"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Print("got new command '%s' from '%s'", text, username)

	//add page:
	//rnd page: /rnd
	//help: /help
	//start: /tart
}
