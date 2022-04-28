package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId primitive.ObjectID `bson:"userId"`
	Text   string             `bson:"text"`
}
