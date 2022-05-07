package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post-service/domain"
)

var posts = []*domain.Post{
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f81"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f85"),
		Text:    "Promovisan sam",
		Picture: nil,
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt", "https://github.com"},
	},
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f83"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f86"),
		Text:    "Dao sam otkaz",
		Picture: nil,
		Links:   []string{"https://github.com/XWS-DISLINKT/dislinkt"},
	},
	{
		Id:      getObjectId("523b0cc3a34d25d8567f9f84"),
		UserId:  getObjectId("623b0cc3a34d25d8567f9f86"),
		Text:    "Ponovo sam se zaposlio",
		Picture: nil,
		Links:   []string{},
	},
}

var reactions = []*domain.PostReaction{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f71"),
		PostId:   getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f91"),
		Reaction: "like",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f72"),
		PostId:   getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f92"),
		Reaction: "dislike",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f73"),
		PostId:   getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId:   getObjectId("623b0cc3a34d25d8567f9f93"),
		Reaction: "like",
	},
}

var comments = []*domain.Comment{
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f71"),
		PostId: getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f91"),
		Text:   "Ut sagittis augue nulla, non suscipit leo malesuada vitae. In ac pretium lorem, at pretium sapien.",
	},
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f72"),
		PostId: getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f92"),
		Text:   "Vestibulum bibendum efficitur felis sit amet volutpat. Nulla ipsum elit, auctor ut tortor quis, tincidunt pretium risus. Phasellus in odio lacus.",
	},
	{
		Id:     getObjectId("223b0cc3a34d25d8567f9f73"),
		PostId: getObjectId("623b0cc3a34d25d8567f9f91"),
		UserId: getObjectId("623b0cc3a34d25d8567f9f93"),
		Text:   "Sed finibus eleifend neque.",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
