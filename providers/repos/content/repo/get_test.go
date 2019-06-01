package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func TestCommentProvider_Get(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Get(context.Background(), &repo.GetRequest{
		ContentId: 343087411107999744,
	})
	if err != nil {
		t.Fatalf("GetComment err: %+v", err)
	}
	t.Logf("GetComment: %+v", res)
}