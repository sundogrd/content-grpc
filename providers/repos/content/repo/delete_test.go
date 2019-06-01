package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func TestCommentProvider_Delete(t *testing.T) {
	gormDB, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := gormDB.Delete(context.Background(), &repo.DeleteRequest{
		ContentId: 343193762765221888,
	})
	if err != nil {
		t.Fatalf("DeleteComment err: %+v", err)
	}
	t.Logf("DeleteComment: %+v", res)
}
