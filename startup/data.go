package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post-service/domain"
)

var posts = []*domain.Post{
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f81"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		Text:    "Promovisan sam",
		Picture: "https://picsum.photos/id/1/200/300",
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt", "https://github.com"},
	},
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f83"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f84"),
		Text:    "Dao sam otkaz",
		Picture: "https://picsum.photos/id/338/200/300",
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt"},
	},
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f84"),
		Text:    "Ponovo sam se zaposlio",
		Picture: "",
		Links:   []string{},
	},
}

var reactions = []*domain.PostReaction{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f71"),
		PostId:   getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f83"),
		Reaction: "like",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f72"),
		PostId:   getObjectId("523b0cc3a34d25d8567f9f83"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f83"),
		Reaction: "dislike",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f73"),
		PostId:   getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f82"),
		Reaction: "like",
	},
}

var comments = []*domain.Comment{
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f71"),
		PostId: getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f82"),
		Text:   "Ut sagittis augue nulla, non suscipit leo malesuada vitae. In ac pretium lorem, at pretium sapien.",
	},
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f72"),
		PostId: getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f83"),
		Text:   "Vestibulum bibendum efficitur felis sit amet volutpat. Nulla ipsum elit, auctor ut tortor quis, tincidunt pretium risus. Phasellus in odio lacus.",
	},
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f73"),
		PostId: getObjectId("523b0cc3a34d25d8567f9f81"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f84"),
		Text:   "Sed finibus eleifend neque.",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
