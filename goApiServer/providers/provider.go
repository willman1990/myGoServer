package providers

import "database/sql"

type Providers struct {
	LoginProvider *LoginProvider
}

var loginProvider *LoginProvider


func NewProviders(db *sql.DB) *Providers {
	loginProvider = NewLoginProvider(db)
	return &Providers{LoginProvider: loginProvider}
}