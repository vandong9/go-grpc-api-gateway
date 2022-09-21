package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/product/pb"
	"net/http"
)

type FindOneRequestBody struct {
	Id int64 `json:"id"`
}

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	b := FindOneRequestBody{}

	err := ctx.BindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: b.Id,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
