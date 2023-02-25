package site_model_methods

import (
	"template-app/models/sitemodel"
	"template-app/pkg/store"
)

func Add(site *sitemodel.Site) error {
	db := store.DB
	return db.Create(site).Error
}
