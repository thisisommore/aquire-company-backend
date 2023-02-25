// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"template-app/models/company"
	"template-app/pkg/store"

	"github.com/TheLazarusNetwork/go-helpers/logo"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(
		&company.Company{},
	)
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
