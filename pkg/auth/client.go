package auth

import (
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/auth/pb"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return
	}

	return pb.NewAuthServiceClient(cc)
}
