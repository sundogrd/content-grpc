package content_test

import (
	"github.com/sundogrd/content-grpc/servers/content"
)

func newTestContentServer() (*content.ContentServiceServer, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &content.ContentServiceServer{
		GormDB: client,
	}, nil
}
