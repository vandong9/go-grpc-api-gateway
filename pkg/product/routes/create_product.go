package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/product/pb"
	"net/http"
)

type CreateProductData struct {
	Name  string `json:"name"`
	Sku   string `json:"sku"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, client pb.ProductServiceClient) {
	b := CreateProductData{}

	err := ctx.BindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  b.Name,
		Sku:   b.Sku,
		Stock: b.Stock,
		Price: b.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
