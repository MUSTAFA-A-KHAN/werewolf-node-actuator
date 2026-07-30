package main

import (
	"fmt"
	"os/exec"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const adminID int64 = 123456789 // your Telegram user ID

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Only allow your Telegram account
		if update.Message.From.ID != adminID {
			continue
		}

		switch update.Message.Text {

		case "/restart_node":
			cmd := exec.Command(
				"bash",
				"-c",
				"cd ~/Server/'Node 1' && nohup ./WerewolfNode > node.log 2>&1 &",
			)

			err := cmd.Run()

			msg := "Node started."
			if err != nil {
				msg = fmt.Sprintf("Failed: %v", err)
			}

			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				msg,
			))

		case "/status":
			cmd := exec.Command(
				"bash",
				"-c",
				"pgrep -af WerewolfNode",
			)

			out, _ := cmd.Output()

			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				string(out),
			))
		}
	}
}
