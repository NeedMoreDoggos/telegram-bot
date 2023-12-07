package event_consumer

import (
	"log"
	"time"

	"github.com/NeedMoreDoggos/telegram-bot/events"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		//TODO: механизм ретрая
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERROR] consumer: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("[ERROR] consumer: %s", err.Error())
			continue
		}

	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)
		/*
			//TODO:
			1. механизм ретрая для события. Фоллбек. Подтверждение.
			2. Обработка всей пачки: останавливаться после первой ошибки, счетчик ошибок.
			3. Параллельная обработка
		*/
		if err := c.processor.Process(event); err != nil {
			log.Printf("cant handle event: %s", err.Error())
			continue
		}
	}

	return nil
}
