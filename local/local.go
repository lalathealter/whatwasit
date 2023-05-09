package local

import (
	"errors"
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
		"lang-set": "switched language to: English",
		"lang-set-err": "error; the language isn't supported",
		"lang-err-no-message": "error; the response message can't be found",
	},
}

const (
	langKey = "lang"
	defaultLangTag = "en"
)

func SetLocalLanguage(c tele.Context, langTag string) (error) {
	_, ok := messageTemplateMap[langTag]
	if !ok {
		return errors.New("lang-set-err")
	}
	c.Set(langKey, langTag)
	return nil
}

func GetMessage(c tele.Context, msg string, responseArgs ...any) string {
	lang, ok := c.Get(langKey).(string)
	if !ok {
		lang = defaultLangTag
		c.Set(langKey, lang)
	}
	message, ok := messageTemplateMap[lang][msg]
	if !ok {
		message = messageTemplateMap[lang]["lang-err-no-message"]
	}
	
	return fmt.Sprintf(message, responseArgs...)  
}


