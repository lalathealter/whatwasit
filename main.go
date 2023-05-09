package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	tele "gopkg.in/telebot.v3"

	"github.com/lalathealter/whatwasit/controllers"
	"github.com/lalathealter/whatwasit/postgre"
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
	b.Handle("/lang", controllers.LocalHandler)

	scheduleAutoDelete()
	
	fmt.Println("Starting the bot...")
	go http.ListenAndServe("/", nil)
	b.Start()
}

func scheduleAutoDelete() {
	s := gocron.NewScheduler(time.UTC)
	CRONSTR := postgre.GetEnv("CLEANSCHEDULE")
	s.Cron(CRONSTR).Do(controllers.ScheduledCleanHandler)
	s.StartAsync()
}

