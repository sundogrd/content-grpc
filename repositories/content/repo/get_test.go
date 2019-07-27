package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func TestContentRepo_Get(t *testing.T) {
	contentRepo, err := initTestRepo()
	if err != nil {
		t.Fatal(err)
	}

	res, err := contentRepo.Get(context.Background(), &repo.GetRequest{
		ContentID: 370631075367497728,
	})
	if err != nil {
		t.Fatalf("GetContent err: %+v", err)
	}
	t.Logf("GetContent: %+v", res.Content)
}
