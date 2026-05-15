package main

import (
	"LangBot/Internal/config"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.MustLoad()
	pref := tele.Settings{
		Token:  os.Getenv(cfg.Token),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет я бот, трекер расходов, я помогаю подсчитывать ваши расходы на разные категории ")
	})
	b.Start()
}
