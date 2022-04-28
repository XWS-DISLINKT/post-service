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
