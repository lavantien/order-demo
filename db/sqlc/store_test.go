package db

import (
	"context"
	"log"
	"order-demo/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderTx(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)
	product := createRandomProduct(t)
	quantity := product.Quantity - util.RandomQuantity()
	log.Println(">> before:", product.Quantity)
	// Run n concurrent orders
	n := 1
	errs := make(chan error)
	results := make(chan OrderTxResult)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// ctx := context.WithValue(context.Background(), txKey, txName)
			ctx := context.Background()
			result, err := store.OrderTx(ctx, OrderTxParams{
				User:     user,
				Product:  product,
				Quantity: quantity,
			})
			errs <- err
			results <- result
		}()
	}
	// Check result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		// Check user
		resultUser := result.User
		require.NotEmpty(t, user)
		require.Equal(t, user.ID, resultUser.ID)
		// Check product
		resultProduct := result.Product
		require.NotEmpty(t, resultProduct)
		require.Equal(t, product.ID, resultProduct.ID)
		// Check order
		order := result.Order
		require.NotEmpty(t, order)
		require.Equal(t, user.ID, order.UserID)
		require.Equal(t, product.ID, order.ProductID)
		require.Equal(t, quantity, order.Quantity)
		require.NotZero(t, order.ID)
		require.NotZero(t, order.CreatedAt)
		_, err = store.GetOrder(context.Background(), order.ID)
		require.NoError(t, err)
		// Check price
		expectedPrice := quantity * product.Cost
		require.Equal(t, expectedPrice, order.Price)
	}
	// Check the final updated product
	updatedProduct, err := testQueries.GetProduct(context.Background(), product.ID)
	require.NoError(t, err)
	log.Println(">> after:", updatedProduct.Quantity)
	require.Equal(t, product.Quantity-quantity, updatedProduct.Quantity)
}

func createProductAndQuantity(t *testing.T) (Product, int64) {
	product := createRandomProduct(t)
	quantity := product.Quantity - util.RandomQuantity()
	return product, quantity
}

func TestAddToCartTx(t *testing.T) {
	store := NewStore(testDB)
	product, quantity := createProductAndQuantity(t)
	log.Println(">> before:", product.Quantity)
	// Run n concurrent orders
	n := 1
	errs := make(chan error)
	results := make(chan AddToCartTxResult)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// ctx := context.WithValue(context.Background(), txKey, txName)
			ctx := context.Background()
			result, err := store.AddToCartTx(ctx, AddToCartTxParams{
				Product:  product,
				Quantity: quantity,
			})
			errs <- err
			results <- result
		}()
	}
	// Check result
	var tempProduct Product
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		// Check product
		tempProduct = result.Product
		require.NotEmpty(t, tempProduct)
		require.Equal(t, product.ID, tempProduct.ID)
	}
	// Check the final updated product
	log.Println(">> after:", tempProduct.Quantity)
	require.Equal(t, product.Quantity-quantity, tempProduct.Quantity)
}

func TestRemoveFromCartTx(t *testing.T) {
	store := NewStore(testDB)
	product, quantity := createProductAndQuantity(t)
	log.Println(">> before:", product.Quantity)
	// Run n concurrent orders
	n := 1
	errs := make(chan error)
	results := make(chan RemoveFromCartTxResult)
	for i := 0; i < n; i++ {
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// ctx := context.WithValue(context.Background(), txKey, txName)
			ctx := context.Background()
			result, err := store.RemoveFromCartTx(ctx, RemoveFromCartTxParams{
				Product:  product,
				Quantity: quantity,
			})
			errs <- err
			results <- result
		}()
	}
	// Check result
	var tempProduct Product
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		// Check product
		tempProduct = result.Product
		require.NotEmpty(t, tempProduct)
		require.Equal(t, product.ID, tempProduct.ID)
	}
	// Check the final updated product
	log.Println(">> after:", tempProduct.Quantity)
	require.Equal(t, product.Quantity+quantity, tempProduct.Quantity)
}
