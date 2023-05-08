package controllers

import (
	"github.com/lalathealter/whatwasit/local"
	"github.com/lalathealter/whatwasit/postgre"
	tele "gopkg.in/telebot.v3"
)

func WelcomeHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "hello"))
}

func SetHandler(c tele.Context) error {
	args, err := parseArgs(c, 3)
	if err != nil {
		return c.Send(local.GetMessage(c, err.Error()))
	}
	serviceName := args[0]
	login := args[1]
	pass := args[2]

	accToken := generateAccessToken(c, serviceName)

	db := postgre.GetWrapper()
	if err := db.SetLogin(login, pass, accToken); err != nil {
		return c.Send(local.GetMessage(c, "set-err-db-error"))
	}
	return c.Send(local.GetMessage(c, "set", serviceName))
}


func GetHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "get"))
}

func DelHandler(c tele.Context) error {
	return c.Send(local.GetMessage(c, "del"))
}

