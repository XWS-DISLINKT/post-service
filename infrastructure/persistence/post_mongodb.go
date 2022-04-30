package persistence

import (
	"context"
	"post-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "post"
)

type PostMongoDb struct {
	posts *mongo.Collection
}

func NewPostMongoDb(client *mongo.Client) domain.IPostService {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDb{
		posts: posts,
	}
}

func (collection *PostMongoDb) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return collection.filterOne(filter)
}

func (collection *PostMongoDb) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return collection.filter(filter)
}

func (collection *PostMongoDb) Insert(post *domain.Post) (*domain.Post, error) {
	result, err := collection.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return post, err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return post, nil
}

func (collection *PostMongoDb) DeleteAll() {
	collection.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func (collection *PostMongoDb) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := collection.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (collection *PostMongoDb) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := collection.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
