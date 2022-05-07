package persistence

import (
	"context"
	"post-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE             = "post"
	COLLECTION           = "post"
	REACTIONS_COLLECTION = "reactions"
)

type PostMongoDb struct {
	posts     *mongo.Collection
	reactions *mongo.Collection
}

func NewPostMongoDb(client *mongo.Client) domain.IPostService {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	reactions := client.Database(DATABASE).Collection(REACTIONS_COLLECTION)
	return &PostMongoDb{
		posts:     posts,
		reactions: reactions,
	}
}

func (collection *PostMongoDb) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return collection.filterPostsOne(filter)
}

func (collection *PostMongoDb) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return collection.filterPosts(filter)
}

func (collection *PostMongoDb) GetByUser(id primitive.ObjectID) ([]*domain.Post, error) {
	filter := bson.M{"userId": id}
	return collection.filterPosts(filter)
}

func (collection *PostMongoDb) Insert(post *domain.Post) error {
	result, err := collection.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (collection *PostMongoDb) DeleteAll() {
	collection.posts.DeleteMany(context.TODO(), bson.D{{}})
	collection.reactions.DeleteMany(context.TODO(), bson.D{{}})
}

func (collection *PostMongoDb) DeleteReaction(postId primitive.ObjectID, userId primitive.ObjectID) {
	filter := bson.M{"userId": userId, "postId": postId}
	collection.reactions.DeleteOne(context.TODO(), filter)
}

func (collection *PostMongoDb) InsertReaction(reaction *domain.PostReaction) error {
	result, err := collection.reactions.InsertOne(context.TODO(), reaction)
	if err != nil {
		return err
	}
	reaction.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (collection *PostMongoDb) filterPosts(filter interface{}) ([]*domain.Post, error) {
	cursor, err := collection.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (collection *PostMongoDb) filterPostsOne(filter interface{}) (post *domain.Post, err error) {
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
