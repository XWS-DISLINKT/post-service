package api

import (
	"context"
	"post-service/application"

	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnsafePostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetByUser(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	posts, err := handler.service.GetByUser(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Post(ctx context.Context, request *pb.PostRequest) (*pb.PostResponse, error) {
	post := mapPostToDomain((*request).Post)
	err := handler.service.Create(post)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: mapPost(post)}, nil
}

func (handler *PostHandler) LikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	reaction := mapReactionToDomain((*request).Reaction)
	handler.service.DeleteReaction(reaction.PostId, reaction.UserId)
	reaction.Reaction = "like"
	err := handler.service.InsertReaction(reaction)
	if err != nil {
		return nil, err
	}
	return &pb.ReactionResponse{PostReaction: mapReaction(reaction)}, nil
}

func (handler *PostHandler) DislikePost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	reaction := mapReactionToDomain((*request).Reaction)
	handler.service.DeleteReaction(reaction.PostId, reaction.UserId)
	reaction.Reaction = "dislike"
	err := handler.service.InsertReaction(reaction)
	if err != nil {
		return nil, err
	}
	return &pb.ReactionResponse{PostReaction: mapReaction(reaction)}, nil
}

func (handler *PostHandler) CommentPost(ctx context.Context, request *pb.CommentRequest) (*pb.CommentResponse, error) {
	comment := mapCommentToDomain((*request).Comment)
	err := handler.service.InsertComment(comment)
	if err != nil {
		return nil, err
	}
	return &pb.CommentResponse{Comment: mapComment(comment)}, nil
}

func (handler *PostHandler) GetAllCommentsByPost(ctx context.Context, request *pb.GetRequest) (*pb.AllCommentsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	comments, err := handler.service.GetAllCommentsByPost(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.AllCommentsResponse{
		Comments: []*pb.Comment{},
	}
	for _, comment := range comments {
		current := mapComment(comment)
		response.Comments = append(response.Comments, current)
	}
	return response, nil
}

func (handler *PostHandler) GetAllReactionsByPost(ctx context.Context, request *pb.GetRequest) (*pb.AllReactionsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	comments, err := handler.service.GetAllReactionsByPost(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.AllReactionsResponse{
		Reactions: []*pb.PostReaction{},
	}
	for _, reaction := range comments {
		current := mapReaction(reaction)
		response.Reactions = append(response.Reactions, current)
	}
	return response, nil
}
