package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id      primitive.ObjectID `bson:"_id"`
	UserId  primitive.ObjectID `bson:"userId"`
	Text    string             `bson:"text"`
	Picture []byte             `bson:"picture"`
	Links   []string           `bson:"links"`
}

type PostReaction struct {
	Id       primitive.ObjectID `bson:"_id"`
	PostId   primitive.ObjectID `bson:"postId"`
	UserId   primitive.ObjectID `bson:"userId"`
	Reaction string             `bson:"reaction"`
}

type Comment struct {
	Id     primitive.ObjectID `bson:"_id"`
	PostId primitive.ObjectID `bson:"postId"`
	UserId primitive.ObjectID `bson:"userId"`
	Text   string             `bson:"text"`
}
