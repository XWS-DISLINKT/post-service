package api

import (
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	"post-service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:     post.Id.Hex(),
		UserId: post.UserId.Hex(),
		Text:   post.Text,
	}
	return postPb
}
