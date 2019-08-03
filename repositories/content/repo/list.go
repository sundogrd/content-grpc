package repo

import (
	"context"
	"github.com/lwyj123/qsearch"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-grpc/repositories"
	"strconv"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)


type listContentQuery struct {
	AppID         *string
	ContentIDs    *[]int64
	Title         *string
	State         *repositories.ContentState
	AuthorID      *int64
}

func parseQuery(qs string) (*listContentQuery, error) {
	if qs == "" {
		return &listContentQuery{}, nil
	}
	var result listContentQuery
	contentIDs := make([]int64, 0)
	queryString := qsearch.NewQueryStringQuery(qs)
	q, err := queryString.Parse()
	if err != nil {
		return nil, err
	}
	for _, queryItem := range q.(*qsearch.BooleanQuery).Should.(*qsearch.DisjunctionQuery).Disjuncts {
		switch queryType := queryItem.(type) {
		case *qsearch.MatchQuery:
			logrus.Debugf("match MatchQuery %#v", queryType)
			if queryType.Field() == "" {
				result.Title = &queryType.Match
			} else if queryType.Field() == "content_id" {
				contentIDNum, err := strconv.ParseInt(queryType.Match, 10, 64)
				if err != nil {
					return nil, err
				}
				contentIDs = append(contentIDs, contentIDNum)
			} else if queryType.Field() == "state" {
				state, err := strconv.ParseInt(queryType.Match, 10, 16)
				if err != nil {
					return nil, err
				}
				contentState := repositories.ContentState(state)
				result.State = &contentState
			} else if queryType.Field() == "author_id" {
				authorIDNum, err := strconv.ParseInt(queryType.Match, 10, 64)
				if err != nil {
					return nil, err
				}
				result.AuthorID = &authorIDNum
			}
		default:
			logrus.Debugf("not match %#v", queryType)
		}
	}
	if len(contentIDs) != 0 {
		result.ContentIDs = &contentIDs
	}
	return &result, nil
}

func (r contentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {
	//queryString := qsearch.NewQueryStringQuery("title:科科 content_id:30214120512 state:3")
	query, err := parseQuery(req.Query)
	if err != nil {
		return nil, err
	}

	var page int32 = 1
	var pageSize int32 = 10
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		pageSize = *req.PageSize
	}
	//
	contents := make([]*repositories.Content, 0)
	count := int64(0)
	//
	db := r.gormDB
	if query.ContentIDs != nil && len(*query.ContentIDs) != 0 {
		db = db.Where("content_id in (?)", *query.ContentIDs)
	}
	if query.Title != nil {
		db = db.Where("title LIKE ?", "%" + *query.Title + "%")
	}
	if query.AuthorID != nil {
		db = db.Where("author_id = ", *query.AuthorID)
	}
	//if req.Type != nil {
	//	db = db.Where("type = ?", *req.Type)
	//}
	if query.State != nil {
		db = db.Where("state = ?", *query.State)
	}
	db.Limit(pageSize).Offset((page - 1) * (pageSize))
	if err := db.Find(&contents).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	} else {
		return &repo.ListResponse{
			List: contents,
			Page: page,
			PageSize: pageSize,
			Total: count,
		}, nil
	}
}
