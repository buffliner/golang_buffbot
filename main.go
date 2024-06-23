package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "7319106212:AAH_pdu7Q3TMW6iWrw9qNrVa2EsklE6ndPM"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			keyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton("btn"),
					tu.KeyboardButton("btn").WithRequestContact(),
					tu.KeyboardButton("btn").WithRequestLocation(),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("btn"),
					tu.KeyboardButton("btn"),
					tu.KeyboardButton("btn"),
				),
			)

			message := tu.Message(
				chatID,
				"Some text",
			).WithReplyMarkup(keyboard)

			_, _ = bot.SendMessage(message)
		}
	}
}
