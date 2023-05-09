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
	"ru":  {
		"hello": "Привет, меня зовут WhatWasIt — я Телеграм-бот который поможет вам сохранять и помнить свои пароли!",
		"set": "Данные входа для %s были сохранены успешно",
		"err-few-args": "ошибка: недостаточно аргументов",
		"err-empty-arg": "ошибка: пустые аргументы не допускаются",
		"err-long-arg": "ошибка: аргумент превышает максимально допустимое количество символов",
		"set-err-db-error": "ошибка сервера: не получилось сохранить данные входа",
		"get": "Логин: [`%s`]\nПароль: [`%s`]",
		"get-err-db-error": "ошибка: не получилось найти данные входа",
		"del": "Ваши данные входа для %s были успешно удалены",
		"del-err-db-error": "ошибка сервера: не получилось удалить данные входа",
		"lang-set": "язык переключён на: Русский",
		"lang-set-err": "ошибка: язык не поддерживается",
		"lang-err-no-message": "ошибка: шаблон ответа не найден",
	},
}

const (
	langKey = "lang"
	defaultLangTag = "en"
)

var usersLangs = map[int64]string{}

func SetLocalLanguage(c tele.Context, langTag string) (error) {
	_, ok := messageTemplateMap[langTag]
	if !ok {
		return errors.New("lang-set-err")
	}
	usersLangs[c.Chat().ID] = langTag
	return nil
}

func GetMessage(c tele.Context, msg string, responseArgs ...any) string {
	lang, ok := usersLangs[c.Chat().ID]
	if !ok {
		SetLocalLanguage(c, defaultLangTag)
	}
	message, ok := messageTemplateMap[lang][msg]
	if !ok {
		message = messageTemplateMap[lang]["lang-err-no-message"]
	}
	
	return fmt.Sprintf(message, responseArgs...)  
}


