package content

import (
	"github.com/sundogrd/content-grpc/repositories/content"
	"time"
)

type Content struct {
	ContentID   int64
	AppID       string
	Title       string
	Description string
	AuthorID    int64
	Category    string
	Type        content.ContentType
	Body        string
	BodyType    content.ContentBodyType
	State       content.ContentState
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Extra       string
}