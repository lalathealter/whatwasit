package local

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

var messageTemplateMap = map[string]map[string]string{
	"en": {
		"hello": "Welcome to the WhatWasIt — a telegram bot that helps you to save and remember your login data!",
		"set": "Login data for %s was set successfully",
		"err-few-args": "error: not enough arguments provided",
		"err-empty-arg": "error: empty arguments aren't allowed",
		"err-long-arg": "error: argument too long",
		"set-err-db-error": "server error: couldn't save credentials",
		"get": "Login: `%s`\nPassword: `%s`",
		"get-err-db-error": "error: couldn't find credentials",
		"del": "Your login data for %s was deleted successfully",
		"del-err-db-error": "server error: couldn't delete credentials",
	},
}

func GetMessage(c tele.Context, msg string, responseArgs ...any) string {
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
	
	return fmt.Sprintf(message, responseArgs...)  
}


