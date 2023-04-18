package telegram

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type TelegramBot struct {
	Token string
}

var bot *TelegramBot

func GetBot() *TelegramBot {
	if bot == nil {
		bot = &TelegramBot{
			Token: os.Getenv("TELEGRAM_BOT_API_KEY"),
		}
	}

	return bot
}

func (bot TelegramBot) SendMessage(chatId string, text string) error {
	request := "https://api.telegram.org/bot" + bot.Token + "/sendMessage" +
		"?chat_id=" + chatId +
		"&text=" + text +
		"&parse_mode=html"

	response, err := http.Get(request)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		return errors.New(bodyString)
	}

	return nil
}
