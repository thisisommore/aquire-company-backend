// Package pasetoclaims provides claim declaration for token generation and verification
package pasetoclaims

import (
	"template-app/models/user"
	"template-app/pkg/store"
	"time"

	"github.com/vk-rv/pvx"
)

// CustomClaims defines claims for paseto containing Email id, signed by and RegisteredClaims
type CustomClaims struct {
	EmailId  string `json:"emailId"`
	SignedBy string `json:"signedBy"`
	pvx.RegisteredClaims
}

// Valid checks if the claims are valid agaist RegisteredClaims and checks if Email id
// exist in database
func (c CustomClaims) Valid() error {
	db := store.DB
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	err := db.Model(&user.User{}).Where("email_id = ?", c.EmailId).First(&user.User{}).Error
	return err
}

// New returns CustomClaims with Email id, signed by and expiration
func New(emailId string, expiration time.Duration, signedBy string) CustomClaims {
	expirationTime := time.Now().Add(expiration)
	return CustomClaims{
		emailId,
		signedBy,
		pvx.RegisteredClaims{
			Expiration: &expirationTime,
		},
	}
}
