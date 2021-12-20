package db

import (
	"context"
	"order-demo/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:     util.RandomName(),
		Cost:     util.RandomCost(),
		Quantity: util.RandomQuantity(),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Cost, product.Cost)
	require.Equal(t, arg.Quantity, product.Quantity)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)
	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Cost, product2.Cost)
	require.Equal(t, product1.Quantity, product2.Quantity)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}
	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}
	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)
	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
