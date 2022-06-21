package blog

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateBlog(ctx context.Context, b Blog) (string, error)
	GetBlog(ctx context.Context, id string) (*Blog, error)
	ListBlogs(ctx context.Context) ([]Blog, error)
	PublishBlog(ctx context.Context, id string) (*Blog, error)
}

var (
	Published = "PUBLISHED"
	Draft     = "DRAFT"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id",bson:"_id"`
	Topic     string             `json:"topic",bson:"topic"`
	Content   string             `json:"content",bson:"content"`
	Status    string             `json:"status",bson:"status"`
	Published time.Time          `json:"published",bson:"published"`
	Author    string             `json:"author",bson:"author"`
	Likes     uint64             `json:"likes",bson:"likes"`
}

var (
	ErrBlogNotFound = errors.New("Blog is not found")
)

type mongoBlogService struct {
	collection mongo.Collection
}

func NewMongoBlogService(c mongo.Collection) Service {
	return &mongoBlogService{
		collection: c,
	}
}

func (s *mongoBlogService) CreateBlog(ctx context.Context, b Blog) (string, error) {
	res, err := s.collection.InsertOne(ctx, b)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	} else {
		return "", errors.New("Unable to get object ID")
	}
}

func (s *mongoBlogService) GetBlog(ctx context.Context, id string) (*Blog, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := s.collection.FindOne(ctx, bson.M{"_id": oid})
	var blog Blog
	if err = res.Decode(&blog); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (s *mongoBlogService) ListBlogs(ctx context.Context) ([]Blog, error) {
	cur, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []Blog
	for cur.Next(ctx) {
		var item Blog
		err = cur.Decode(&item)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	return results, nil
}

func (s *mongoBlogService) PublishBlog(ctx context.Context, id string) (*Blog, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := s.collection.FindOne(ctx, bson.M{"_id": oid})
	err = res.Err()
	if err != nil {
		return nil, err
	}

	var b Blog
	err = res.Decode(&b)
	if err != nil {
		return nil, err
	}

	b.Published = time.Now()
	b.Status = Published

	update, err := s.collection.UpdateOne(ctx, bson.M{"_id": oid}, b)
	if err != nil {
		return nil, err
	}

	if update.ModifiedCount > 0 {
		return &b, nil
	} else {
		return nil, errors.New("Unable to update document")
	}
}
