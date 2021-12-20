// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetOrder(ctx context.Context, id int64) (Order, error)
	GetOrderForUpdate(ctx context.Context, id int64) (Order, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (Product, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserForUpdate(ctx context.Context, id int64) (User, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
}

var _ Querier = (*Queries)(nil)