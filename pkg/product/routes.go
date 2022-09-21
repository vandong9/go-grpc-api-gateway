package product

import (
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
)

func RegisterRouter(g *gin.Engine, c *config.Config) {
	svc := &ServiceClient{Client: InitServiceClient(c)}

	routes := g.Group("/product")
	routes.POST("/create", svc.CreateProduct)
}

func (s *ServiceClient) CreateProduct(ctx *gin.Context) {
	CreateProduct(ctx, s.Client)
}
