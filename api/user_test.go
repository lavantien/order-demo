package api

import (
	db "order-demo/db/sqlc"
	"order-demo/util"
)

func randomUser() db.User {
	return db.User{
		Username: util.RandomName(),
		HashedPassword: util.RandomPassword(),
		FullName: util.RandomName(),
		Email:          util.RandomEmail(),
	}
}
