package main

import (
	"fmt"
	"log"
	"os"
	"time"
	

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
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

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	fmt.Println("Starting the bot...")
	b.Start()
}


func MustGetEnv(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Value for the %s wasn't provided in the .env file", key)
	}
	return envVal
}
