package api

import (
	"post-service/domain"

	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:      post.Id.Hex(),
		UserId:  post.UserId.Hex(),
		Text:    post.Text,
		Picture: post.Picture,
		Links:   post.Links,
	}
	return postPb
}

func mapPostToDomain(post *pb.Post) *domain.Post {
	id, err := primitive.ObjectIDFromHex(post.Id)
	if err != nil {
		return nil
	}
	userId, err := primitive.ObjectIDFromHex(post.UserId)
	if err != nil {
		return nil
	}
	return &domain.Post{
		Id:      id,
		UserId:  userId,
		Text:    post.Text,
		Picture: post.Picture,
		Links:   post.Links,
	}
}
