package startup

import (
	"fmt"
	post "github.com/XWS-DISLINKT/dislinkt/common/proto/post-service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"post-service/application"
	"post-service/domain"
	"post-service/infrastructure/api"
	"post-service/infrastructure/persistence"
	"post-service/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	iPostService := server.initIPostService(mongoClient)

	postService := server.initPostService(iPostService)

	postHandler := server.initPostHandler(postService)

	server.startGrpcServer(postHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initIPostService(client *mongo.Client) domain.IPostService {
	collection := persistence.NewPostMongoDb(client)
	collection.DeleteAll()
	for _, post := range posts {
		err := collection.Insert(post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return collection
}

func (server *Server) initPostService(iPostService domain.IPostService) *application.PostService {
	return application.NewPostService(iPostService)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
