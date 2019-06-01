package content

import (
	"github.com/jinzhu/gorm"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
)

type ContentServiceServer struct {
	GormDB              *gorm.DB
	content.ContentServiceServer
}