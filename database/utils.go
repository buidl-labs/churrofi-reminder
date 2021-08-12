package database

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
)

func AddReminder(DB *pg.DB, action string, email string) (Reminder, error) {
	reminder := &Reminder{
		Action: action,
		Email:  email,
	}
	now := time.Now()

	var sendAt time.Time
	switch action {
	case "withdraw":
		sendAt = now.AddDate(0, 0, 3)
		// sendAt = now.AddDate(0, 0, 0) : Uncomment when testing

	case "activate":
		sendAt = now.AddDate(0, 0, 1)
		// sendAt = now.AddDate(0, 0, 0) : Uncomment when testing

	default:
		return *reminder, errors.New("action not allowed")
	}

	reminder.SendAt = sendAt

	if _, err := DB.Model(reminder).Insert(); err != nil {
		return *reminder, err
	}

	return *reminder, nil
}

func GetAllReminders(DB *pg.DB) ([]*Reminder, error) {
	var reminders []*Reminder
	if err := DB.Model(&reminders).Select(); err != nil {
		return reminders, err
	}

	return reminders, nil
}

func GetRemindersToSend(DB *pg.DB) ([]*Reminder, error) {
	var reminders []*Reminder

	QUERY := "send_at < ? and sent is false"
	if err := DB.Model(&reminders).Where(QUERY, time.Now()).Select(); err != nil {
		return reminders, err
	}

	return reminders, nil
}

func MarkReminderDone(DB *pg.DB, reminder *Reminder) error {
	reminder.Sent = true
	if _, err := DB.Model(reminder).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
