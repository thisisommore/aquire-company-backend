// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"template-app/models/sitemodel"
	"template-app/pkg/store"

	"github.com/TheLazarusNetwork/go-helpers/logo"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(
		&sitemodel.Site{},
	)
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
