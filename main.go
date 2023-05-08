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

	b.Handle("/start", welcomeHandler)
	b.Handle("/set", setHandler)
	b.Handle("/get", getHandler)
	b.Handle("/del", delHandler)

	fmt.Println("Starting the bot...")
	b.Start()
}

var (
	msgWelcome = `
		HELLO
	`
)

var messageTemplateMap = map[string]map[string]string{
	"en": {
		"hello": "Welcome to the WhatWasIt — a telegram bot that helps you to save and remember your login data!",
		"set": "Login data for %s was set successfully",
		"get": "Login: %s\nPassword: %s",
		"del": "Your login data for %s was deleted successfully",
	},
}

func getMessage(c tele.Context, msg string) string {
	lang, ok := c.Get("lang").(string)
	if !ok {
		lang = "en"
		c.Set("lang", lang)
	}
	message, ok := messageTemplateMap[lang][msg]
	if !ok {
		if lang == "" {
			message = "error; the set language isn't supported"
		} 
		if msg == "" {
			message = "error; the message can't be found"
		} 
	}
	return message  
}

func welcomeHandler(c tele.Context) error {
	return c.Send(getMessage(c, "hello"))
}

func setHandler(c tele.Context) error {
	return c.Send(getMessage(c, "set"))
}

func getHandler(c tele.Context) error {
	return c.Send(getMessage(c, "get"))
}

func delHandler(c tele.Context) error {
	return c.Send(getMessage(c, "del"))
}

func MustGetEnv(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Value for the %s wasn't provided in the .env file", key)
	}
	return envVal
}
