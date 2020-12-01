package main

import (
	"context"
	"main/database"
	"main/server"

	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// var serverLogFilePath = "./logs/server.log"

	// TODO: create logs path?
	// if file, err := os.OpenFile(serverLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766); err != nil {
	// 	logrus.Errorf("can't open file '%s' (cause: %s) - logging into STDOUT", serverLogFilePath, err.Error())
	// } else {
	// 	logrus.SetOutput(file)
	// 	logrus.Info("logging intited")
	// }
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(
		&database.RssUrls{},
		&database.Author{},
		&database.Category{},
		&database.FeedsCategories{},
		&database.Feed{}); err != nil {
		panic(err)
	}

	db.Create(&database.RssUrls{
		Url:      "http://feeds.twit.tv/twit.xml",
		Duration: time.Second * 10,
	})

	s := server.New(context.Background(), db, 30*time.Second)
	s.Start()
	defer s.Stop()
	// Defines 5 minute server
	time.Sleep(5 * time.Minute)

	// TODO: start rpc server
}
