package blog

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateBlogEndpoint  endpoint.Endpoint
	GetBlogEndpoint     endpoint.Endpoint
	ListBlogEndpoint    endpoint.Endpoint
	PublishBlogEndpoint endpoint.Endpoint
}

func MakeServerEndpoint(s Service) Endpoints {
	return Endpoints{
		CreateBlogEndpoint:  makeCreateBlogEndpoint(s),
		GetBlogEndpoint:     makeGetBlogEndpoint(s),
		ListBlogEndpoint:    makeListBlogsEndpoint(s),
		PublishBlogEndpoint: makePublishBlogEndpoint(s),
	}
}

func makeCreateBlogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createBlogRequest)
		b := Blog{
			Topic:   req.Topic,
			Content: req.Content,
			Author:  req.Author,
			Status:  Draft,
		}
		id, err := s.CreateBlog(ctx, b)
		return createBlogResponse{ID: id}, err
	}
}

func makeGetBlogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBlogRequest)
		b, err := s.GetBlog(ctx, req.ID)
		return b, err
	}
}

func makeListBlogsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listBlogRequest)
		list, err := s.ListBlogs(ctx)
		return list, err
	}
}

func makePublishBlogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(publishBlogRequest)
		b, err := s.PublishBlog(ctx, req.ID)
		return b, err
	}
}

type createBlogRequest struct {
	Topic   string `json:"topic"`
	Content string `json:"content"`
	Author  string `json:"content"`
}

type createBlogResponse struct {
	ID string `json: "id"`
}

type getBlogRequest struct {
	ID string
}

type listBlogRequest struct {
}

type publishBlogRequest struct {
	ID string
}
