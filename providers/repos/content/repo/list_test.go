package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func TestCommentProvider_List(t *testing.T) {
	gormDB, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := gormDB.List(context.Background(), &repo.ListRequest{

	})
	if err != nil {
		t.Fatalf("ListComment err: %+v", err)
	}

	t.Logf("ListComment: %+v", res)
}
