package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post-service/domain"
)

type PostService struct {
	iPostService domain.IPostService
}

func NewPostService(iPostService domain.IPostService) *PostService {
	return &PostService{
		iPostService: iPostService,
	}
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
