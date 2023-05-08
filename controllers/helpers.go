package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

const MAX_ARG_LEN = 256

func generateAccessToken(c tele.Context, servName string) string {
	userID := c.Chat().ID
	privateID := []byte(strconv.FormatInt(userID, 10))
	loweredServName := strings.ToLower(servName)

	hasher := sha256.New()
	hasher.Write([]byte(loweredServName))
	hashSum := hasher.Sum(privateID)

	hexHashString := hex.EncodeToString(hashSum)
	return hexHashString
}

func readArgConsideringQuotes(input string) string {
	if len(input) < 1 {
		return ""
	}

	fch := input[0]
	if fch == '\'' || fch == '"' {
		if len(input) <= 2 {
			return ""
		}
		return input[1 : len(input)-1]
	}
	return input
}


func parseArgs(c tele.Context, argCount int) ([]string, error) {
	passedArgs := c.Args()
	if len(passedArgs) < argCount {
		return nil, errors.New("err-few-args")
	}

	argsSlice := passedArgs[:argCount]
	for i, val := range argsSlice {
		nextVal := readArgConsideringQuotes(val)
		if nextVal == "" {
			return nil, errors.New("err-empty-arg")
		}
		if len(nextVal) > MAX_ARG_LEN {
			return nil, errors.New("err-long-arg")
		}
		argsSlice[i] = nextVal
	}

	return argsSlice, nil
}
