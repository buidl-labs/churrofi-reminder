package database

import "time"

type Reminder struct {
	ID        string    `pg:"default:gen_random_uuid()"`
	Action    string    `pg:",notnull"`
	Email     string    `pg:",notnull"`
	CreatedAt time.Time `pg:"default:now()::timestamp"`
	SendAt    time.Time `pg:",notnull"`
	Sent      bool      `pg:"default:false,use_zero"`
}
