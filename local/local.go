package local

import (
	"errors"
	"fmt"

	tele "gopkg.in/telebot.v3"
)

var messageTemplateMap = map[string]map[string]string{
	"en": {
		"hello": `Welcome to the WhatWasIt — a telegram bot that helps you to save and remember your login data!
/set service_name login pass — set or update your login data for a service
/get service_name — get your login data for the service
/del service_name — delete your login data for a service 
/lang lang_tag — switch languages (available: en, ru)
/help, /doc, /start — show this message
Note: if you don't use or update your login data in 15 days, it's going to be deleted;
If you want to use spaces in your names, logins or passwords, enclose them in double quotes like this: "my service name"`,
		"set": "Login data for %s was set successfully",
		"err-few-args": "error: not enough arguments provided",
		"err-empty-arg": "error: empty arguments aren't allowed",
		"err-long-arg": "error: argument too long",
		"set-err-db-error": "server error: couldn't save credentials",
		"get": "Login:`%s`\nPassword:`%s`",
		"get-err-db-error": "error: couldn't find credentials",
		"del": "Your login data for %s was deleted successfully",
		"del-err-db-error": "server error: couldn't delete credentials",
		"lang-set": "switched language to: English",
		"lang-set-err": "error; the language isn't supported",
		"lang-err-no-message": "error; the response message can't be found",
	},
	"ru":  {
		"hello": `Привет, меня зовут WhatWasIt — я Телеграм-бот который поможет вам сохранять и помнить свои пароли!
/set имя_сервиса логин пароль — установить или перезаписать данные для входа в сервис
/get имя_сервиса — получить данные для входа в сервис
/del имя_сервиса — удалить данные для входа в сервис
/lang сокращение_языка — сменить язык (доступны: en, ru)
/help, /doc, /start — вывести эту справку
Внимание: если вы ни разу за 15 дней не запросите или не обновите данные для входа в сервис, то они будут удалены;
Если вы хотите использовать пробелы в ваших именах для сервисов, паролях или логинах, то окружите их двойными кавычками, вот так: "моё имя для сервиса"`,
		"set": "Данные входа для %s были сохранены успешно",
		"err-few-args": "ошибка: недостаточно аргументов",
		"err-empty-arg": "ошибка: пустые аргументы не допускаются",
		"err-long-arg": "ошибка: аргумент превышает максимально допустимое количество символов",
		"set-err-db-error": "ошибка сервера: не получилось сохранить данные входа",
		"get": "Логин:`%s`\nПароль:`%s`",
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
		lang = defaultLangTag
	}
	message, ok := messageTemplateMap[lang][msg]
	if !ok {
		message = messageTemplateMap[lang]["lang-err-no-message"]
	}
	
	return fmt.Sprintf(message, responseArgs...)  
}


