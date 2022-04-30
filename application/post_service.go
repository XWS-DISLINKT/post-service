package application

import (
	"post-service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *PostService) Insert(postRequest *domain.Post) (post *domain.Post, err error) {
	return service.iPostService.Insert(postRequest)
}
