package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post-service/domain"
)

var posts = []*domain.Post{
	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f82"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		Text:    "Promovisan sam",
		Picture: nil,
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt", "https://github.com"},
	},
	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f83"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f83"),
		Text:    "Dao sam otkaz",
		Picture: nil,
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt"},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
