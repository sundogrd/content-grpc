package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func TestContentRepo_List(t *testing.T) {
	contentRepo, err := initTestRepo()
	if err != nil {
		t.Fatal(err)
	}

	res, err := contentRepo.List(context.Background(), &repo.ListRequest{

	})
	if err != nil {
		t.Fatalf("ListComment err: %+v", err)
	}

	t.Logf("ListComment: %+v", res)
}
