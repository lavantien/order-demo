package api

import (
	"net/http"
	db "order-demo/db/sqlc"
	"order-demo/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type listOrdersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) listOrders(ctx *gin.Context) {
	var req listOrdersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Enforce authorization Rule0 and Rule6
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var orders []db.Order
	var err error
	if authPayload.Username == "admin" {
		arg := db.ListOrdersParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
		}
		orders, err = server.store.ListOrders(ctx, arg)
	} else {
		arg := db.ListOrdersByOwnerParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
			Owner:  authPayload.Username,
		}
		orders, err = server.store.ListOrdersByOwner(ctx, arg)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

type createOrderRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required,min=1"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Enforce authorization Rule5
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.OrderTxParams{
		User:     user,
		Product:  product,
		Quantity: req.Quantity,
	}
	orderResult, err := server.store.OrderTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderResult)
}
