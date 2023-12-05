package telegram

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/NeedMoreDoggos/telegram-bot/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
	AddCmd   = "/add"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Print("got new command '%s' from '%s'", text, username)

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case AddCmd:
		return p.savePage(chatID, text, username)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.unknownCommand(chatID)
	}

}

func (p *Processor) savePage(chatId int, pageURL string, username string) error {
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return fmt.Errorf("cant do command: save page. %w", err)
	}

	if isExists {
		return p.tg.SendMessage(chatId, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return fmt.Errorf("cant save page message: %w", err)
	}

	if err := p.tg.SendMessage(chatId, msgSaved); err != nil {
		return fmt.Errorf("cant save page message: %w", err)
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) error {
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return fmt.Errorf("cant do command: \"send random\": %w", err)
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return fmt.Errorf("cant do command: \"send random\": %w", err)
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func (p *Processor) unknownCommand(chatId int) error {
	return p.tg.SendMessage(chatId, msgUnknownCommand)
}
