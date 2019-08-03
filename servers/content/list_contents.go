package content

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
	contentrepo "github.com/sundogrd/content-grpc/repositories/content"
)

func packContent(repoContent *contentrepo.Content) (*content.Content, error) {
	return &content.Content{
		ContentId: repoContent.ContentID,
		AppId: repoContent.AppID,
		Title: repoContent.Title,
		Description: repoContent.Description,
		AuthorId: repoContent.AuthorID,
		Category: repoContent.Category,
		Type: content.EContentType(repoContent.Type),
		Body: repoContent.Body,
		BodyType: content.EContentBodyType(repoContent.BodyType),
		State: content.EContentState(repoContent.State),
		Extra: repoContent.Extra,
		CreatedAt: uint32(repoContent.CreatedAt.Unix()),
		UpdatedAt: uint32(repoContent.UpdatedAt.Unix()),
	}, nil
}

func (server ContentServiceServer) ListContents(ctx context.Context, req *content.ListContentsRequest) (*content.ListContentsResponse, error) {
	listRes, err := server.ContentRepo.List(ctx, &contentrepo.ListRequest{
		//AppId: req.AppId,
		Page: &req.Page,
		PageSize: &req.PageSize,
	})
	if err != nil {
		logrus.Errorf("ListContents err=%s", err.Error())
		return nil, err
	}
	resultContents := make([]*content.Content, 0)
	for _, repoItem := range listRes.List {
		resultItem, err := packContent(repoItem)
		if err != nil {
			return nil, err
		}
		resultContents = append(resultContents, resultItem)
	}
	return &content.ListContentsResponse{
		AppId: req.AppId,
		Contents: resultContents,
		Total: listRes.Total,
		Page: listRes.Page,
		PageSize: listRes.PageSize,
	}, nil
}
