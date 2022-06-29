package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	Id      primitive.ObjectID `bson:"_id"`
	UserId  primitive.ObjectID `bson:"userId"`
	Text    string             `bson:"text"`
	Picture string             `bson:"picture"`
	Links   []string           `bson:"links"`
}

type PostReaction struct {
	Id       primitive.ObjectID `bson:"_id"`
	PostId   primitive.ObjectID `bson:"postId"`
	UserId   primitive.ObjectID `bson:"userId"`
	Reaction string             `bson:"reaction"`
	Username string             `bson:"username"`
}

type Comment struct {
	Id       primitive.ObjectID `bson:"_id"`
	PostId   primitive.ObjectID `bson:"postId"`
	UserId   primitive.ObjectID `bson:"userId"`
	Text     string             `bson:"text"`
	Username string             `bson:"username"`
}

type Job struct {
	Id          primitive.ObjectID `bson:"_id"`
	Position    string             `bson:"position"`
	CompanyName string             `bson:"companyName"`
	//Seniority   Job_Seniority          `bson:"seniority"`
	Location    string                 `bson:"location"`
	Description string                 `bson:"description"`
	ClosingDate *timestamppb.Timestamp `bson:"closingDate"`
	UserId      string                 `bson:"userId"`
}

type UserApiKey struct {
	UserId primitive.ObjectID `bson:"_id"`
	ApiKey string             `bson:"apiKey"`
}

type JobPosition struct {
	JobId    primitive.ObjectID `bson:"_id"`
	Position string             `bson:"position"`
}
