package order

import (
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitOrderServiceClient(config *config.Config) pb.OrderServiceClient {

	con, err := grpc.Dial(config.OrderSvcUrl)

	if err != nil {
		return nil
	}
	return pb.NewOrderServiceClient(con)

}
