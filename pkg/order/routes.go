package order

import (
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/auth"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/order/routes"
)

func RegisterRoutes(e gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	middle := auth.InitMiddlewareConfig(authSvc)

	svc := &ServiceClient{Client: InitOrderServiceClient(c)}
	routes := e.Group("/order")
	routes.Use(middle.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
