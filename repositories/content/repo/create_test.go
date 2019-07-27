package repo_test

import (
	"context"
	"github.com/sundogrd/gopkg/pointer"
	"testing"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func TestContentProvider_Create(t *testing.T) {
	contentRepo, err := initTestRepo()
	if err != nil {
		t.Fatal(err)
	}

	res, err := contentRepo.Create(context.Background(), &repo.CreateRequest{
		AppID: "lwio",
		Title: "test",
		Description: nil,
		AuthorID: 443474713,
		Category: nil,
		Type: &repo.TYPE_RICHTEXT,
		Body: "#worinima",
		BodyType: (*repo.ContentBodyType)(pointer.Int16(int16(repo.BODY_TYPE_MARKDOWN))),
		Extra:    map[string]interface{}{
			"keke": true,
		},
	})
	if err != nil {
		t.Fatalf("CreateComment err: %+v", err)
	}
	t.Logf("CreateComment: %+v", res)
}
