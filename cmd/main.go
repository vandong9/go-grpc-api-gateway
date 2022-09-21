package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/auth"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/config"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRouter(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)

}
