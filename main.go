package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/lalathealter/whatwasit/local"
	"github.com/lalathealter/whatwasit/postgre"
)

func main() {
	TGTOKEN := postgre.GetEnv("TGTOKEN")

	pref := tele.Settings{
		Token: TGTOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", welcomeHandler)
	b.Handle("/docs", welcomeHandler)
	b.Handle("/help", welcomeHandler)

	b.Handle("/set", setHandler)
	b.Handle("/get", getHandler)
	b.Handle("/del", delHandler)

	fmt.Println("Starting the bot...")
	b.Start()
}

func welcomeHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "hello"))
}

func setHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "set"))
}

func getHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "get"))
}

func delHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "del"))
}

