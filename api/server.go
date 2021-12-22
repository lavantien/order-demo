package api

import (
	db "order-demo/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.GET("/products", server.listProducts)
	router.POST("/products", server.createProduct)
	router.POST("/products/cart/add", server.addToCart)
	router.POST("/products/cart/remove", server.removeFromCart)
	router.GET("/orders", server.listOrders)
	router.POST("/orders", server.createOrder)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
