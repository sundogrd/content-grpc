package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func TestCommentProvider_Delete(t *testing.T) {
	contentRepo, err := initTestRepo()
	if err != nil {
		t.Fatal(err)
	}

	res, err := contentRepo.Delete(context.Background(), &repo.DeleteRequest{
		ContentId: 343193762765221888,
	})
	if err != nil {
		t.Fatalf("DeleteComment err: %+v", err)
	}
	t.Logf("DeleteComment: %+v", res)
}
