package services

import (
	connection "github.com/XWS-DISLINKT/dislinkt/common/proto/connection-service"
	profile "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ConnectionsClient(address string) connection.ConnectionServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to connection service: %v", err)
	}
	return connection.NewConnectionServiceClient(conn)
}

func ProfilesClient(address string) profile.ProfileServiceClient {
	prof, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to profile service: %v", err)
	}
	return profile.NewProfileServiceClient(prof)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
