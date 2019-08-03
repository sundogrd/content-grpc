package content_test

import (
	"context"
	"fmt"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
	"testing"
)

func TestContentServer_ListContents(t *testing.T) {
	server, err := newTestContentServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	res, err := server.ListContents(context.Background(), &content.ListContentsRequest{
		AppId:     "lwio",
		Page: 1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ListContents: %+v", res)
}
