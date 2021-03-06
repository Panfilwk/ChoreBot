package main

import (
	"log"
  "os"
	"time"

  tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
  b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_API_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
}