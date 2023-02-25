package sitemodel

type Site struct {
	Name string `json:"name" gorm:"primaryKey"`
}
