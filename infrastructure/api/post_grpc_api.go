package api

import (
	"context"
	"fmt"
	"post-service/application"
	"post-service/domain"
	"post-service/infrastructure/services"
	"post-service/startup/config"

	connection "github.com/XWS-DISLINKT/dislinkt/common/proto/connection-service"
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	profile "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
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

func (handler *PostHandler) GetSuggestJobsFor(ctx context.Context, request *pb.GetSuggestJobsForRequest) (*pb.GetSuggestJobsForResponse, error) {

	cfg := config.NewConfig()
	profileAddress := fmt.Sprintf(cfg.ProfileServiceHost + ":" + cfg.ProfileServicePort)
	profileResponse, _ := services.ProfilesClient(profileAddress).Get(context.TODO(), &profile.GetRequest{Id: request.Id})

	jobs, err := handler.service.SuggestJobs(profileResponse.Education[0].Degree, profileResponse.Experience[0].JobTitle)
	if err != nil {
		return nil, err
	}
	response := &pb.GetSuggestJobsForResponse{}
	for _, job := range jobs {
		response.JobPositions = append(response.JobPositions, &pb.JobPosition{
			JobId:    job.JobId.Hex(),
			Position: job.Position,
		})
	}
	return response, nil
}

func (handler *PostHandler) SuggestJob(ctx context.Context, request *pb.SuggestJobRequest) (*pb.SuggestJobResponse, error) {
	jobs, err := handler.service.SuggestJobs(request.Skill, request.Experience)
	if err != nil {
		return nil, err
	}
	response := &pb.SuggestJobResponse{}
	for _, job := range jobs {
		response.JobPositions = append(response.JobPositions, &pb.JobPosition{
			JobId:    job.JobId.Hex(),
			Position: job.Position,
		})
	}
	return response, nil
}

func (handler *PostHandler) PostJobDislinkt(ctx context.Context, request *pb.PostJobDislinktRequest) (*pb.Job, error) {
	response := request.Job
	job := mapJobToDomain(response)
	err := handler.service.InsertJob(job)
	if err != nil {
		return nil, err
	}
	response.Id = job.Id.Hex()
	return response, nil
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
	userId, err := handler.getUserID(request.ApiKey)
	if err != nil {
		return nil, err
	}
	response := request.Job
	job := mapJobToDomain(response)
	job.UserId = userId
	err = handler.service.InsertJob(job)
	if err != nil {
		return nil, err
	}
	response.Id = job.Id.Hex()
	response.UserId = userId
	return response, nil
}

func (handler *PostHandler) getUserID(apiKey string) (string, error) {
	userApiKey, err := handler.service.GetUserApiKey(apiKey)
	if err != nil {
		return "", err
	}
	return userApiKey.UserId.Hex(), nil
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
		Posts: []*pb.PostM{},
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
		Posts: []*pb.PostM{},
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
	cfg := config.NewConfig()
	connectionAddress := fmt.Sprintf(cfg.ConnectionServiceHost + ":" + cfg.ConnectionServicePort)
	connectionResponse, _ := services.ConnectionsClient(connectionAddress).GetConnectionsUsernamesFor(context.TODO(),
		&connection.GetConnectionsUsernamesRequest{Id: id})

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
		Posts: []*pb.PostM{},
	}
	for _, post := range feed {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Post(ctx context.Context, request *pb.PostM) (*pb.PostM, error) {
	post := mapPostToDomain(request)
	err := handler.service.Create(post)
	if err != nil {
		return nil, err
	}

	//dobavljanje idjeva konekcija
	connectionIdsStr := make([]string, 0)
	cfg := config.NewConfig()
	connectionAddress := fmt.Sprintf(cfg.ConnectionServiceHost + ":" + cfg.ConnectionServicePort)
	connectionResponse, _ := services.ConnectionsClient(connectionAddress).GetConnectionsUsernamesFor(context.TODO(),
		&connection.GetConnectionsUsernamesRequest{Id: post.UserId.Hex()})

	if connectionResponse.Usernames != nil {
		connectionIdsStr = connectionResponse.Usernames //[]string{"623b0cc3a34d25d8567f9f85"} //
	} else {
		connectionIdsStr = []string{} //"623b0cc3a34d25d8567f9f86"}
	}
	fmt.Printf("connection ids: {%s}", connectionIdsStr)
	//kreiranje notifikacijeza svakog od njih
	profileAddress := fmt.Sprintf(cfg.ProfileServiceHost + ":" + cfg.ProfileServicePort)
	for _, cis := range connectionIdsStr {
		profileResponse, _ := services.ProfilesClient(profileAddress).SendNotification(context.TODO(),
			&profile.NewNotificationRequest{SenderId: post.UserId.Hex(), ReceiverId: cis, NotificationType: "post"})

		fmt.Printf("\ncreated notification {%s}", profileResponse.Id)
	}

	return mapPost(post), nil
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
