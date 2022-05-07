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
	POST_COLLECTION      = "post"
	REACTIONS_COLLECTION = "reactions"
	COMMENT_COLLECTION   = "comments"
)

type PostMongoDb struct {
	posts     *mongo.Collection
	reactions *mongo.Collection
	comments  *mongo.Collection
}

func NewPostMongoDb(client *mongo.Client) domain.IPostService {
	posts := client.Database(DATABASE).Collection(POST_COLLECTION)
	reactions := client.Database(DATABASE).Collection(REACTIONS_COLLECTION)
	comments := client.Database(DATABASE).Collection(COMMENT_COLLECTION)
	return &PostMongoDb{
		posts:     posts,
		reactions: reactions,
		comments:  comments,
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

func (collection *PostMongoDb) GetAllReactionsByPost(id primitive.ObjectID) ([]*domain.PostReaction, error) {
	filter := bson.M{"postId": id}
	return collection.filterReactions(filter)
}

func (collection *PostMongoDb) GetAllCommentsByPost(id primitive.ObjectID) ([]*domain.Comment, error) {
	filter := bson.M{"postId": id}
	return collection.filterComments(filter)
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
	collection.comments.DeleteMany(context.TODO(), bson.D{{}})
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

func (collection *PostMongoDb) InsertComment(comment *domain.Comment) error {
	result, err := collection.comments.InsertOne(context.TODO(), comment)
	if err != nil {
		return err
	}
	comment.Id = result.InsertedID.(primitive.ObjectID)
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
func (collection *PostMongoDb) filterReactions(filter interface{}) ([]*domain.PostReaction, error) {
	cursor, err := collection.reactions.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeReaction(cursor)
}

func (collection *PostMongoDb) filterComments(filter interface{}) ([]*domain.Comment, error) {
	cursor, err := collection.comments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeComment(cursor)
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

func decodeReaction(cursor *mongo.Cursor) (reactions []*domain.PostReaction, err error) {
	for cursor.Next(context.TODO()) {
		var reaction domain.PostReaction
		err = cursor.Decode(&reaction)
		if err != nil {
			return
		}
		reactions = append(reactions, &reaction)
	}
	err = cursor.Err()
	return
}
func decodeComment(cursor *mongo.Cursor) (comments []*domain.Comment, err error) {
	for cursor.Next(context.TODO()) {
		var comment domain.Comment
		err = cursor.Decode(&comment)
		if err != nil {
			return
		}
		comments = append(comments, &comment)
	}
	err = cursor.Err()
	return
}
