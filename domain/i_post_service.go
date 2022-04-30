package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPostService interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(post *Post) (*Post, error)
	DeleteAll()
}
