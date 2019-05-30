package content

import (
	"github.com/jinzhu/gorm"
)

type ContentServiceServer struct {
	GormDB              *gorm.DB
}
