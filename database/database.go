package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitConnection(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&RssUrls{},
		&Author{},
		&Item{},
		&FeedItems{},
		&Feed{},
	)

	return db, err
}

func GetUrls(db *gorm.DB) ([]RssUrls, error) {
	var urls []RssUrls
	return urls, db.Find(&urls).Error
}

func AddRss(db *gorm.DB, url string, duration time.Duration, rule *string) error {
	return db.Create(&RssUrls{
		Url:         url,
		Duration:    duration,
		ParsingRule: rule,
	}).Error
}

func ListNews(db *gorm.DB) ([]Item, error) {
	var items []Item
	return items, db.Find(&items).Error
}

func GetNews(db *gorm.DB, request string) ([]Item, error) {
	var items []Item
	return items, db.Find(&items, "title LIKE ?", "%"+request+"%").Error
}

func CreateFeed(db *gorm.DB, feed Feed) error {
	return db.Create(&feed).Error
}

func CreateItem(db *gorm.DB, item Item) error {
	return db.Create(&item).Error
}
