package cmd

import "github.com/gin-gonic/gin"
import (
	log
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway/pkg/config"
)

func main() {
	c, err := config.LoadConfig();
	if err != nil {
		log.Fatalln("Failed at config", err);
	}

	r := gin.Default();
}