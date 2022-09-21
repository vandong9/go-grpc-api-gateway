package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vandong9/go-grpc-api-gateway.git/pkg/order/pb"
	"net/http"
)

type CreateOrderRequestBody struct {
	productId int64 `json:"productId"`
	quantity  int64 `json:"quantity"`
	userId    int64 `json:"userId"`
}

func CreateOrder(ctx *gin.Context, client pb.OrderServiceClient) {
	b := CreateOrderRequestBody{}

	err := ctx.BindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		ProductId: b.productId,
		Quantity:  b.quantity,
		UserId:    b.userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
