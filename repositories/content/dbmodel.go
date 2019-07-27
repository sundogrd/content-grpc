package content

import (
	"time"
)

//BodyType ...
type ContentBodyType int16

// ...
const (
	BODY_TYPE_UNKNOWN_STATE ContentBodyType = 0
	BODY_TYPE_TEXT          ContentBodyType = 1
	BODY_TYPE_HTML          ContentBodyType = 2
	BODY_TYPE_MARKDOWN      ContentBodyType = 3
	BODY_TYPE_LATEX         ContentBodyType = 4
)

// ContentType ...
type ContentType int16

// ...
var (
	TYPE_UNKNOWN_STATE ContentType = 0
	TYPE_RICHTEXT      ContentType = 1
)

type ContentState int16

var (
	STATE_UNKNOWN_STATE ContentState = 0
	STATE_DRAFT         ContentState = 1
	STATE_CHECKING      ContentState = 2
	STATE_PUBLISHED     ContentState = 3
	STATE_DELETED       ContentState = 255
)

// SDContent http://gorm.io/docs/models.html
type Content struct {
	// gorm.Model
	ID          int64           `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID   int64           `gorm:"not null;"`
	AppID       string          `gorm:"not null;"`
	Title       string          `gorm:"type:varchar(60);not null"`
	Description string          `gorm:"type:varchar(300);not null"`
	AuthorID    int64           `gorm:"not null;"`
	Category    string          `gorm:"type:varchar(60)"`
	Type        ContentType     `gorm:"type:INT;NOT NULL"`
	Body        string          `gorm:"type:TEXT;NOT NULL"`
	BodyType    ContentBodyType `gorm:"type:INT;NOT NULL;DEFAULT:1"`
	State       ContentState    `gorm:"type:INT;NOT NULL;DEFAULT:10"`
	CreatedAt   time.Time       `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time       `gorm:""`
	Extra       string          `gorm:"type:TEXT;"`
}

func (*Content) TableName() string {
	return "sd_contents"
}
