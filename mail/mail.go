package mail

import (
	"context"
	"fmt"
	"reminders/database"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

func SendReminder(reminder database.Reminder, domain string, apiKey string) (string, error) {

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		fmt.Sprintf("manan@%s", domain),
		"Reminder from ChurroFi!",
		buildMessage(reminder.Action),
		reminder.Email,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

func buildMessage(action string) string {
	var msg string
	switch action {
	case "withdraw":
		msg = "Hey! Your CELO is waiting withdrawal. Make sure to withdraw it to your account, cheers."

	case "activate":
		msg = "Hey! Your CELO is waiting to be activated. You need to activate your investment for it to start earning for you. Make sure to activate your CELO soon, cheers."
	}

	return msg
}
