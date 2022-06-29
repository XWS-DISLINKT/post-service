package persistence

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"post-service/startup/config"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func GetDriver() (neo4j.Driver, error) {
	cfg := config.NewConfig()
	uri := fmt.Sprintf("%s://%s:%s/", cfg.Neo4jProtocol, cfg.Neo4jHost, cfg.Neo4jPort)
	return neo4j.NewDriver(uri, neo4j.BasicAuth(cfg.Neo4jUsername, cfg.Neo4jPassword, ""))
}
