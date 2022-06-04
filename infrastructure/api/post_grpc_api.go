package api

import (
	"context"
	"os"
	"post-service/application"
	"post-service/domain"
	"post-service/infrastructure/services"

	connection "github.com/XWS-DISLINKT/dislinkt/common/proto/connection-service"
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

func (handler *PostHandler) RegisterApiKey(ctx context.Context, request *pb.GetApiKeyRequest) (*pb.GetApiKeyResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	userApiKey := domain.UserApiKey{
		UserId: userId,
	}
	err = handler.service.RegisterApiKey(&userApiKey)
	if err != nil {
		return nil, err
	}
	return &pb.GetApiKeyResponse{
		UserId: userApiKey.UserId.Hex(),
		ApiKey: userApiKey.ApiKey,
	}, nil
}

func (handler *PostHandler) PostJob(ctx context.Context, request *pb.PostJobRequest) (*pb.Job, error) {
	if !handler.apiKeyValid(request.UserId, request.ApiKey) {
		return nil, nil
	}
	response := request.Job
	job := mapJobToDomain(response)
	err := handler.service.InsertJob(job)
	if err != nil {
		return nil, err
	}
	response.Id = job.Id.Hex()
	return response, nil
}

func (handler *PostHandler) apiKeyValid(userId string, apiKey string) bool {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return false
	}
	userApiKey, err := handler.service.GetUserApiKey(objectId)
	if err != nil {
		return false
	}
	return userApiKey.ApiKey == apiKey
}

func (handler *PostHandler) SearchJobsByPosition(ctx context.Context, request *pb.SearchJobsByPositionRequest) (*pb.SearchJobsByPositionResponse, error) {
	jobs, err := handler.service.SearchJobsByPosition(request.Search)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchJobsByPositionResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *PostHandler) GetAllJobs(ctx context.Context, request *pb.GetAllJobsRequest) (*pb.GetAllJobsResponse, error) {
	jobs, err := handler.service.GetAllJobs()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllJobsResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
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

func (handler *PostHandler) GetFeed(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	id := request.Id
	connectionIdsStr := make([]string, 0)
	connectionResponse, _ := services.ConnectionsClient("localhost:8004").GetConnectionsUsernamesFor(context.TODO(),
		&connection.GetConnectionsUsernamesRequest{Id: id})
	if _, err := os.Stat("/.dockerenv"); err == nil {
		connectionResponse, _ = services.ConnectionsClient(os.Getenv("CONNECTION_SERVICE_HOST")+":"+os.Getenv("CONNECTION_SERVICE_PORT")).GetConnectionsUsernamesFor(context.TODO(),
			&connection.GetConnectionsUsernamesRequest{Id: id})
	}

	if connectionResponse.Usernames != nil {
		connectionIdsStr = connectionResponse.Usernames //[]string{"623b0cc3a34d25d8567f9f85"} //
	} else {
		connectionIdsStr = []string{} //"623b0cc3a34d25d8567f9f86"}
	}
	//connectionIdsStr := []string{"623b0cc3a34d25d8567f9f86"}
	var feed []*domain.Post
	for _, cIdStr := range connectionIdsStr {
		cId, err := primitive.ObjectIDFromHex(cIdStr)
		if err != nil {
			return nil, err
		}
		posts, err := handler.service.GetByUser(cId)
		feed = append(feed, posts...)
	}

	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range feed {
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
