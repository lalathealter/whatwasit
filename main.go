package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"

	"github.com/lalathealter/whatwasit/local"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}

func MustGetEnv(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Value for the %s wasn't provided in the .env file", key)
	}
	return envVal
}

func main() {
	TGTOKEN := MustGetEnv("TGTOKEN")

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

