package application

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"post-service/domain"
	"post-service/infrastructure/persistence"
	"strconv"
)

type PostService struct {
	iPostService domain.IPostService
	neo4jDriver  neo4j.Driver
}

func NewPostService(iPostService domain.IPostService) *PostService {
	neo4jDriver, err := persistence.GetDriver()
	if err != nil {
		panic(err)
	}
	err = neo4jDriver.VerifyConnectivity()
	if err != nil {
		panic(err)
	}
	return &PostService{
		iPostService: iPostService,
		neo4jDriver:  neo4jDriver,
	}
}
func (service *PostService) SearchJobsByPosition(search string) ([]*domain.Job, error) {
	return service.iPostService.SearchJobsByPosition(search)
}

func (service *PostService) RegisterApiKey(key *domain.UserApiKey) error {
	userApiKey, err := service.iPostService.GetUserApiKeyById(key.UserId)
	if err == nil {
		*key = *userApiKey
		return nil
	}
	key.ApiKey = strconv.Itoa(rand.Int())
	return service.iPostService.RegisterApiKey(key)
}

func (service *PostService) GetAllJobs() ([]*domain.Job, error) {
	return service.iPostService.GetAllJobs()
}

func (service *PostService) InsertJob(job *domain.Job) error {
	session := service.neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"create (job:Job {id: $id, position: $position}) return job is not null",
			map[string]interface{}{"id": job.Id.Hex(), "userId": job.UserId, "description": job.Description, "location": job.Location, "companyName": job.CompanyName, "position": job.Position, "closingDate": job.ClosingDate.String()})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	return service.iPostService.InsertJob(job)
}

type JobPositions struct {
	jobPositions []*domain.JobPosition
}

func (service *PostService) SuggestJobs(skill string, experience string) ([]*domain.JobPosition, error) {
	session := service.neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	jobs, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		records, err := transaction.Run(
			"match (job:Job) where job.position = $skill or job.position = $experience return job.position, job.id",
			map[string]interface{}{"skill": skill, "experience": experience})

		positions := JobPositions{jobPositions: []*domain.JobPosition{}}

		if records == nil {
			return positions, nil
		}

		for records.Next() {
			record := records.Record()
			position, _ := record.Get("job.position")
			jobId, _ := record.Get("job.id")
			hexJobId, _ := primitive.ObjectIDFromHex(jobId.(string))
			positions.jobPositions = append(positions.jobPositions, &domain.JobPosition{
				JobId:    hexJobId,
				Position: position.(string),
			})
		}
		return positions, err
	})
	return jobs.(JobPositions).jobPositions, err
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

func (service *PostService) GetUserApiKey(apiKey string) (*domain.UserApiKey, error) {
	return service.iPostService.GetUserApiKey(apiKey)
}
