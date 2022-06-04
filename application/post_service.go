package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"post-service/domain"
	"strconv"
)

type PostService struct {
	iPostService domain.IPostService
}

func NewPostService(iPostService domain.IPostService) *PostService {
	return &PostService{
		iPostService: iPostService,
	}
}
func (service *PostService) SearchJobsByPosition(search string) ([]*domain.Job, error) {
	return service.iPostService.SearchJobsByPosition(search)
}

func (service *PostService) RegisterApiKey(key *domain.UserApiKey) error {
	key.ApiKey = strconv.Itoa(rand.Int())
	return service.iPostService.RegisterApiKey(key)
}

func (service *PostService) GetAllJobs() ([]*domain.Job, error) {
	return service.iPostService.GetAllJobs()
}

func (service *PostService) InsertJob(job *domain.Job) error {
	return service.iPostService.InsertJob(job)
}

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.iPostService.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.iPostService.GetAll()
}

func (service *PostService) GetByUser(id primitive.ObjectID) ([]*domain.Post, error) {
	return service.iPostService.GetByUser(id)
}

func (service *PostService) Create(postRequest *domain.Post) error {
	return service.iPostService.Insert(postRequest)
}

func (service *PostService) InsertReaction(reaction *domain.PostReaction) error {
	return service.iPostService.InsertReaction(reaction)
}

func (service *PostService) InsertComment(comment *domain.Comment) error {
	return service.iPostService.InsertComment(comment)
}

func (service *PostService) DeleteReaction(postId primitive.ObjectID, userId primitive.ObjectID) {
	service.iPostService.DeleteReaction(postId, userId)
}

func (service *PostService) GetAllReactionsByPost(id primitive.ObjectID) ([]*domain.PostReaction, error) {
	return service.iPostService.GetAllReactionsByPost(id)
}

func (service *PostService) GetAllCommentsByPost(id primitive.ObjectID) ([]*domain.Comment, error) {
	return service.iPostService.GetAllCommentsByPost(id)
}

func (service *PostService) GetUserApiKey(id primitive.ObjectID) (*domain.UserApiKey, error) {
	return service.iPostService.GetUserApiKey(id)
}
