// Package pasetomiddleware defines middleware to verify PASETO token with required claims
package pasetomiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"template-app/pkg/paseto"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/vk-rv/pvx"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var EmailIdInContext = "EmailId"
var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		err = fmt.Errorf("failed to bind header, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(headers.Authorization, ErrAuthHeaderMissing)
		httpo.NewErrorResponse(httpo.AuthHeaderMissing, ErrAuthHeaderMissing.Error()).Send(c, http.StatusBadRequest)
		c.Abort()
		return
	}
	emailId, err := paseto.VerifyPaseto(headers.Authorization)
	if err != nil {
		var validationErr *pvx.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.HasExpiredErr() {
				err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
				logValidationFailed(headers.Authorization, err)
				httpo.NewErrorResponse(httpo.TokenExpired, "token expired").Send(c, http.StatusUnauthorized)
				c.Abort()
				return
			}
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set(EmailIdInContext, emailId)
}

func logValidationFailed(token string, err error) {
	logo.Warnf("validation failed with token %v and error: %v", token, err)
}
