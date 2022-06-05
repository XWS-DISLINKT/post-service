package api

import (
	"post-service/domain"

	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapJobToDomain(reaction *pb.Job) *domain.Job {
	return &domain.Job{
		Id:          primitive.NewObjectID(),
		Position:    reaction.Position,
		CompanyName: reaction.CompanyName,
		Location:    reaction.Location,
		Description: reaction.Description,
		ClosingDate: reaction.ClosingDate,
	}
}

func mapJob(post *domain.Job) *pb.Job {
	postPb := &pb.Job{
		Id:          post.Id.Hex(),
		Position:    post.Position,
		CompanyName: post.CompanyName,
		Seniority:   0,
		Location:    post.Location,
		Description: post.Description,
		ClosingDate: nil,
	}
	return postPb
}

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

func mapReaction(reaction *domain.PostReaction) *pb.PostReaction {
	reactionPb := &pb.PostReaction{
		Id:       reaction.Id.Hex(),
		UserId:   reaction.UserId.Hex(),
		PostId:   reaction.PostId.Hex(),
		Reaction: reaction.Reaction,
		Username: reaction.Username,
	}
	return reactionPb
}

func mapReactionToDomain(reaction *pb.PostReaction) *domain.PostReaction {
	id, err := primitive.ObjectIDFromHex(reaction.Id)
	if err != nil {
		return nil
	}
	userId, err := primitive.ObjectIDFromHex(reaction.UserId)
	if err != nil {
		return nil
	}
	postId, err := primitive.ObjectIDFromHex(reaction.PostId)
	if err != nil {
		return nil
	}
	return &domain.PostReaction{
		Id:       id,
		UserId:   userId,
		PostId:   postId,
		Reaction: reaction.Reaction,
		Username: reaction.Username,
	}
}

func mapComment(comment *domain.Comment) *pb.Comment {
	commentPb := &pb.Comment{
		Id:       comment.Id.Hex(),
		UserId:   comment.UserId.Hex(),
		PostId:   comment.PostId.Hex(),
		Text:     comment.Text,
		Username: comment.Username,
	}
	return commentPb
}

func mapCommentToDomain(comment *pb.Comment) *domain.Comment {
	id, err := primitive.ObjectIDFromHex(comment.Id)
	if err != nil {
		return nil
	}
	userId, err := primitive.ObjectIDFromHex(comment.UserId)
	if err != nil {
		return nil
	}
	postId, err := primitive.ObjectIDFromHex(comment.PostId)
	if err != nil {
		return nil
	}
	return &domain.Comment{
		Id:       id,
		UserId:   userId,
		PostId:   postId,
		Text:     comment.Text,
		Username: comment.Username,
	}
}
