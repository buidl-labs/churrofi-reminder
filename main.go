package main

import (
	"context"
	"log"
	"os"
	"reminders/database"
	"reminders/handlers"
	"reminders/mail"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DB_URL := os.Getenv("DB_URL")
	log.Println(DB_URL)
	if DB_URL == "" {
		log.Fatal("Please provide a DB url.")
	}

	DB := database.New(DB_URL)
	defer DB.Close()
	// DB.AddQueryHook(pgdebug.DebugHook{
	// 	Verbose: true,
	// })

	ctx := context.Background()
	if err := DB.Ping(ctx); err != nil {
		log.Println(err)
	}

	// dropTable(DB)
	createTable(DB, false)

	// SCHEDULER TO PERIODICALLY SEND EMAILS
	scheduler := cron.New()
	scheduler.AddFunc("@every 2h", func() {
		MAIL_DOMAIN := os.Getenv("MAILGUN_DOMAIN")
		MAIL_APIKEY := os.Getenv("MAILGUN_API_KEY")
		log.Println("Fetching reminders to send")
		remindersToSend, err := database.GetRemindersToSend(DB)
		if err != nil {
			log.Fatal("Can't fetch reminders from DB")
		}

		log.Println("Amount of reminders to send", len(remindersToSend))
		for _, reminder := range remindersToSend {
			if _, err = mail.SendReminder(*reminder, MAIL_DOMAIN, MAIL_APIKEY); err == nil {
				log.Println(database.MarkReminderDone(DB, reminder))
			}

		}
	})
	scheduler.Start()

	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Routes
	r.GET("/", handlers.IndexGet())
	r.GET("/reminder", handlers.ReminderGet(DB))
	r.POST("/reminder", handlers.ReminderPost(DB))

	// Running the app
	r.Run()
}

func dropTable(DB *pg.DB) {

	_, err := DB.Exec("drop table reminders")
	if err != nil {
		panic(err)
	}
}

func createTable(DB *pg.DB, temp bool) {
	err := DB.Model((*database.Reminder)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: temp,
	})
	if err != nil {
		log.Print(err)
	}
}
