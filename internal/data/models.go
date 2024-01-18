package data

import (
	"database/sql"
	"errors"

	"github.com/suensky/greenlight/internal/jsonlog"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
	Tokens TokenModel
}

func NewModels(db *sql.DB, logger *jsonlog.Logger) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
		Users: UserModel{
			DB:     db,
			Logger: logger,
		},
		Tokens: TokenModel{
			DB:     db,
			Logger: logger,
		},
	}
}
