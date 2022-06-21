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
		id, err := s.CreateBlog(ctx, req.Blog)
		return createBlogResponse{ID: id, Err: err}, err
	}
}

func makeGetBlogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBlogRequest)
		b, err := s.GetBlog(ctx, req.ID)
		return getBlogResponse{Item: b, Err: err}, err
	}
}

func makeListBlogsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listBlogRequest)
		list, err := s.ListBlogs(ctx)
		return listBlogResponse{Items: list, Err: err}, err
	}
}

func makePublishBlogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(publishBlogRequest)
		b, err := s.PublishBlog(ctx, req.ID)
		return publishBlogResponse{Item: b, Err: err}, err
	}
}

type createBlogRequest struct {
	Blog Blog
}

type createBlogResponse struct {
	ID  string
	Err error
}

type getBlogRequest struct {
	ID string
}

type getBlogResponse struct {
	Item *Blog
	Err  error
}

type listBlogRequest struct {
}

type listBlogResponse struct {
	Items []Blog
	Err   error
}

type publishBlogRequest struct {
	ID string
}

type publishBlogResponse struct {
	Item *Blog
	Err  error
}
