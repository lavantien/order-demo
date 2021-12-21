package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error)
	AddToCartTx(ctx context.Context, arg AddToCartTxParams) (AddToCartTxResult, error)
	RemoveFromCartTx(ctx context.Context, arg RemoveFromCartTxParams) (RemoveFromCartTxResult, error)
}

type DBStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &DBStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *DBStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) // Default Isolation Level of postgres is Read Committed
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type OrderTxParams struct {
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int64   `json:"quantity"`
}

type OrderTxResult struct {
	User    User    `json:"user"`
	Product Product `json:"product"`
	Order   Order   `json:"order"`
}

// var txKey = struct{}{}

// OrderTx performs a payment from the user to the warehouse
// It creates a order record, add the appropriate user and product, calculate the price, and update product's quantity, within a single database transaction
func (store *DBStore) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// txName := ctx.Value(txKey)
		// Check if buyer's quantity exceed quantity in warehouse
		if arg.Quantity > arg.Product.Quantity {
			result = OrderTxResult{
				User:    arg.User,
				Product: arg.Product,
				Order:   Order{},
			}
			return errors.New("not enough quantity in warehouse")
		}
		// Calculate price = quantity * product's cost
		price := arg.Quantity * arg.Product.Cost
		// Create order with calculated information
		// log.Println(txName, "create order")
		order, err := q.CreateOrder(ctx, CreateOrderParams{
			UserID:    arg.User.ID,
			ProductID: arg.Product.ID,
			Quantity:  arg.Quantity,
			Price:     price,
		})
		if err != nil {
			return err
		}
		// Adjust product quantity and form the result
		product, err := updateProductQuantity(ctx, q, arg.Product, arg.Quantity)
		result = OrderTxResult{
			User:    arg.User,
			Product: product,
			Order:   order,
		}
		// Need deadlock avoidance?
		return err
	})
	return result, err
}

func updateProductQuantity(ctx context.Context, q *Queries, product Product, quantity int64) (result Product, err error) {
	result, err = q.UpdateProduct(ctx, UpdateProductParams{
		ID:       product.ID,
		Quantity: product.Quantity - quantity,
	})
	return
}

type AddToCartTxParams struct {
	Quantity int64   `json:"quantity"`
	Product  Product `json:"product"`
}

type AddToCartTxResult struct {
	Product Product `json:"product"`
}

// If this operation fails, prevent user from adding
func (store *DBStore) AddToCartTx(ctx context.Context, arg AddToCartTxParams) (AddToCartTxResult, error) {
	var result AddToCartTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		if arg.Quantity > arg.Product.Quantity {
			result = AddToCartTxResult{
				Product: arg.Product,
			}
			return errors.New("not enough quantity in warehouse")
		}
		arg.Product.Quantity -= arg.Quantity
		result = AddToCartTxResult{
			Product: arg.Product,
		}
		return nil
	})
	return result, err
}

type RemoveFromCartTxParams struct {
	Quantity int64   `json:"quantity"`
	Product  Product `json:"product"`
}

type RemoveFromCartTxResult struct {
	Product Product `json:"product"`
}

func (store *DBStore) RemoveFromCartTx(ctx context.Context, arg RemoveFromCartTxParams) (RemoveFromCartTxResult, error) {
	var result RemoveFromCartTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		arg.Product.Quantity += arg.Quantity
		result = RemoveFromCartTxResult{
			Product: arg.Product,
		}
		return nil
	})
	return result, err
}
