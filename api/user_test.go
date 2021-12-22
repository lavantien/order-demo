package api

import (
	db "order-demo/db/sqlc"
	"order-demo/util"
)

func randomUser() db.User {
	return db.User{
		ID:             util.RandomInt(1, 1000),
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomPassword(),
	}
}
