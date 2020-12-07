package database

import "time"

type RssUrls struct {
	Url      string        `gorm:"PRIMARY_KEY;TYPE:text;"`
	Duration time.Duration `gorm:"TYPE:text"`
	ParsingRule *string `gorm:"TYPE:text"`
}

type Feed struct {
	ID          uint        `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string      `gorm:"TYPE:text"`
	Description string      `gorm:"TYPE:text"`
	Link        string      `gorm:"TYPE:text"`
	FeedLink    string      `gorm:"TYPE:text"`
	Items       []FeedItems `gorm:"ASSOCIATION_AUTOCREATE:true;ASSOCIATION_AUTOUPDATE:true;"`
}

type FeedItems struct {
	ID     uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	FeedID uint `gorm:"REFERENCES:feed(id)"`
	ItemID uint `gorm:"REFERENCES:item(id)"`
	Item   Item `gorm:"ASSOCIATION_AUTOCREATE:true;ASSOCIATION_AUTOUPDATE:true;FOREIGNKEY:ItemID"`
}

type Author struct {
	ID    uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name  string `gorm:"TYPE:text"`
	Email string `gorm:"TYPE:text"`
}

type Item struct {
	ID          uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string `gorm:"TYPE:text"`
	Description string `gorm:"TYPE:text"`
	Content     string `gorm:"TYPE:text"`
	Link        string `gorm:"TYPE:text"`
	AuthorID    uint   `gorm:"REFERENCES:author(id)"`
	Author      Author `gorm:"FOREIGNKEY:AuthorID"`
}
