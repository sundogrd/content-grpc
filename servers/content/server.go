package content

import (
	"github.com/jinzhu/gorm"
	contentRepo "github.com/sundogrd/content-grpc/repositories/content"
)

type ContentServiceServer struct {
	GormDB      *gorm.DB
	ContentRepo contentRepo.Repo
	//content.ContentServiceServer
}
