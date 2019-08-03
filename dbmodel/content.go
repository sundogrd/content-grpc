package dbmodel

import "time"

// SDContent http://gorm.io/docs/models.html
type DBContent struct {
	// gorm.Model
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID   int64     `gorm:"not null;"`
	AppID       string    `gorm:"not null;"`
	Title       string    `gorm:"type:varchar(60);not null"`
	Description string    `gorm:"type:varchar(300);not null"`
	AuthorID    int64     `gorm:"not null;"`
	Category    string    `gorm:"type:varchar(60)"`
	Type        int16     `gorm:"type:INT;NOT NULL"`
	Body        string    `gorm:"type:TEXT;NOT NULL"`
	BodyType    int16     `gorm:"type:INT;NOT NULL;DEFAULT:1"`
	State       int16     `gorm:"type:INT;NOT NULL;DEFAULT:10"`
	CreatedAt   time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time `gorm:""`
	Extra       string    `gorm:"type:TEXT;"`
}

func (*DBContent) TableName() string {
	return "sd_contents"
}
