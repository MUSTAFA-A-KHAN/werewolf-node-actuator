package main

import (
	"fmt"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const adminID int64 = 1006461736 // Telegram user ID

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
		fmt.Println(update.Message.Text)
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
			fmt.Println("Checking node status...")

			cmd := exec.Command(
				"bash",
				"-c",
				"pgrep -af WerewolfNode",
			)

			out, err := cmd.Output()

			var status string

			if err != nil || len(out) == 0 {
				status = "🔴 WerewolfNode is not running."
			} else {
				status = "🟢 WerewolfNode is running:\n" + string(out)
			}

			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				status,
			))

		}
	}
}
