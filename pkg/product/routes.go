package product

import (
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/auth"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/product/routes"
)

func RegisterRoutes(g *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitMiddlewareConfig(authSvc)

	svc := &ServiceClient{Client: InitServiceClient(c)}

	routes := g.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (s *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, s.Client)
}

func (s *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, s.Client)
}
