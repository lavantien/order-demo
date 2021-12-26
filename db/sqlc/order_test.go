package db

import (
	"context"
	"order-demo/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)
	product := createRandomProduct(t)
	quantity := util.RandomQuantity()
	price := product.Cost * quantity
	arg := CreateOrderParams{
		Owner:     user.Username,
		ProductID: product.ID,
		Quantity:  quantity,
		Price:     price,
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)
	require.Equal(t, arg.Owner, order.Owner)
	require.Equal(t, arg.ProductID, order.ProductID)
	require.Equal(t, arg.Quantity, order.Quantity)
	require.Equal(t, arg.Price, order.Price)
	require.NotZero(t, order.ID)
	require.NotZero(t, order.CreatedAt)
	return order
}

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)
	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.Owner, order2.Owner)
	require.Equal(t, order1.ProductID, order2.ProductID)
	require.Equal(t, order1.Quantity, order2.Quantity)
	require.Equal(t, order1.Price, order2.Price)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}
	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}
	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)
	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}

func TestListOrdersByOwner(t *testing.T) {
	var lastOrder Order
	for i := 0; i < 10; i++ {
		lastOrder = createRandomOrder(t)
	}
	arg := ListOrdersByOwnerParams{
		Limit:  5,
		Offset: 0,
		Owner:  lastOrder.Owner,
	}
	orders, err := testQueries.ListOrdersByOwner(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, lastOrder.Owner, order.Owner)
	}
}
