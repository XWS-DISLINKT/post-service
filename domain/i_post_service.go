package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPostService interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	GetByUser(id primitive.ObjectID) ([]*Post, error)
	Insert(post *Post) error
	DeleteAll()
	InsertReaction(reaction *PostReaction) error
	DeleteReaction(postId primitive.ObjectID, userId primitive.ObjectID)
	InsertComment(comment *Comment) error
}
