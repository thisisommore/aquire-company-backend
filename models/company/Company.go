// Package company provides model and methods for storing and retriving company identified by Email id
package company

// Company company model with Email id
type Company struct {
	EmailId      string `json:"emailId" gorm:"primaryKey;not null"`
	Name         string `json:"name" gorm:"unique"`
	Price int `json:"price"`
	OpenToAquire *bool  `json:"openToHire" `
}
