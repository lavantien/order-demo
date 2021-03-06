// Code generated by sqlc. DO NOT EDIT.
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO
    orders (owner, product_id, quantity, price)
VALUES ($1, $2, $3, $4) RETURNING id, owner, product_id, quantity, price, created_at
`

type CreateOrderParams struct {
	Owner     string `json:"owner"`
	ProductID int64  `json:"product_id"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.Owner,
		arg.ProductID,
		arg.Quantity,
		arg.Price,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT
    id, owner, product_id, quantity, price, created_at
FROM
    orders
WHERE
    id = $1
LIMIT
    1
`

func (q *Queries) GetOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT
    id, owner, product_id, quantity, price, created_at
FROM
    orders
ORDER BY
    id
LIMIT
    $1
OFFSET
    $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.ProductID,
			&i.Quantity,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrdersByOwner = `-- name: ListOrdersByOwner :many
SELECT
    id, owner, product_id, quantity, price, created_at
FROM
    orders
WHERE
    owner = $3
ORDER BY
    id
LIMIT
    $1
OFFSET
    $2
`

type ListOrdersByOwnerParams struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Owner  string `json:"owner"`
}

func (q *Queries) ListOrdersByOwner(ctx context.Context, arg ListOrdersByOwnerParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrdersByOwner, arg.Limit, arg.Offset, arg.Owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.ProductID,
			&i.Quantity,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
