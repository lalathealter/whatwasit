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

func parseSpacesAndQuotes(c tele.Context) []string {
	passedArgs := make([]string, 0)
	sb := strings.Builder{}
	skipSpaces := true
	for _, char := range c.Message().Payload {
		if char == '"' {
			if sb.Len() == 0 {
				skipSpaces = false
			} else {
				passedArgs = append(passedArgs, sb.String())
				sb.Reset()
				skipSpaces = true
			}
			continue
		}
		
		if char == ' ' && skipSpaces {
			if sb.Len() > 0 {
				passedArgs = append(passedArgs, sb.String())
				sb.Reset()
			}
			continue
		}
		sb.WriteRune(char)
	}
	passedArgs = append(passedArgs, sb.String())
	return passedArgs
}

func parseArgs(c tele.Context, argCount int) ([]string, error) {
	passedArgs := parseSpacesAndQuotes(c)
	if len(passedArgs) < argCount {
		return nil, errors.New("err-few-args")
	}
	passedArgs = passedArgs[:argCount]
	for _, val := range passedArgs {
		if len(val) < 1 {
			return nil, errors.New("err-empty-arg")
		}
		if len(val) > MAX_ARG_LEN {
			return nil, errors.New("err-long-arg")
		}
	}

	return passedArgs, nil
}
