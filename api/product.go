package api

import (
	"net/http"
	db "order-demo/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type listProductsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	products, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, products)
}

type addToCartRequest struct {
	ProductID int64 `json:"product_id" biding:"required"`
	Quantity  int64 `json:"quantity" biding:"required,min=1"`
}

func (server *Server) addToCart(ctx *gin.Context) {
	var req addToCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.AddToCartTxParams{
		Product:  product,
		Quantity: req.Quantity,
	}
	tempProduct, err := server.store.AddToCartTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tempProduct)
}

type removeFromCartRequest struct {
	ProductID int64 `json:"product_id" biding:"required"`
	Quantity  int64 `json:"quantity" biding:"required,min=1"`
}

func (server *Server) removeFromCart(ctx *gin.Context) {
	var req removeFromCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.RemoveFromCartTxParams{
		Product:  product,
		Quantity: req.Quantity,
	}
	tempProduct, err := server.store.RemoveFromCartTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tempProduct)
}

type createProductParams struct {
	Name     string `json:"name" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=0"`
	Cost     int64  `json:"cost" binding:"required,min=0"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateProductParams{
		Name:     req.Name,
		Cost:     req.Cost,
		Quantity: req.Quantity,
	}
	product, err := server.store.CreateProduct(ctx, arg)
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
	ctx.JSON(http.StatusOK, product)
}
