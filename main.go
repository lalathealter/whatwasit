package main

import (
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"

	"github.com/lalathealter/whatwasit/postgre"
	"github.com/lalathealter/whatwasit/controllers"
)

func main() {
	TGTOKEN := postgre.GetEnv("TGTOKEN")

	pref := tele.Settings{
		Token: TGTOKEN,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", controllers.WelcomeHandler)
	b.Handle("/docs", controllers.WelcomeHandler)
	b.Handle("/help", controllers.WelcomeHandler)

	b.Handle("/set", controllers.SetHandler)
	b.Handle("/get", controllers.GetHandler)
	b.Handle("/del", controllers.DelHandler)

	fmt.Println("Starting the bot...")
	b.Start()
}

