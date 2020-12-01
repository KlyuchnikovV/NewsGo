package database

import (
	"time"
)

func init() {

}

type RssUrls struct {
	Url      string        `gorm:"PRIMARY_KEY;TYPE:text;"`
	Duration time.Duration `gorm:"TYPE:text"`
}

// TODO: check link if appears in db
type Feed struct {
	ID          uint              `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string            `gorm:"TYPE:text"`
	Description string            `gorm:"TYPE:text"`
	Link        string            `gorm:"TYPE:text"`
	FeedLink    string            `gorm:"TYPE:text"`
	Updated     string            `gorm:"TYPE:text"`
	Published   string            `gorm:"TYPE:text"`
	Language    string            `gorm:"TYPE:text"`
	AuthorID    uint              `gorm:"REFERENCES:author(id)"`
	Author      Author            `gorm:"FOREIGNKEY:AuthorID"`
	Categories  []FeedsCategories `gorm:"ASSOCIATION_AUTOCREATE:true;ASSOCIATION_AUTOUPDATE:true;"`
}

type Author struct {
	ID    uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name  string `gorm:"TYPE:text"`
	Email string `gorm:"TYPE:text"`
}

type FeedsCategories struct {
	ID         uint     `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	FeedID     uint     `gorm:"REFERENCES:feed(id)"`
	CategoryID uint     `gorm:"REFERENCES:category(id)"`
	Category   Category `gorm:"ASSOCIATION_AUTOCREATE:true;ASSOCIATION_AUTOUPDATE:true;FOREIGNKEY:CategoryID"`
}

type Category struct {
	ID   uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string `gorm:"PRIMARY_KEY;TYPE:text;"`
}
