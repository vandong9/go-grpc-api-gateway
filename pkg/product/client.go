package product

import (
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	pc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil
	}

	return pb.NewProductServiceClient(pc)
}
