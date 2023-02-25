package httphelper

import (
	"fmt"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

func HandleDBError(err error, action string, c *gin.Context) {
	c.Abort()
	apiRes := GetDBErrorAPIRes(err, action)
	apiRes.Send(c, 500)
}

func GetDBErrorAPIRes(err error, action string) (apiRes *httpo.ApiErrorResponse[any]) {
	// Check for gorm errors
	if err == gorm.ErrRecordNotFound {
		return httpo.NewErrorResponse(404, "some of the fields not found")
	}

	// Check for postgres errors
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case pgerrcode.ForeignKeyViolation:
			err = fmt.Errorf("some of fields doesn't exist: %s", pgErr.ConstraintName)
			return httpo.NewErrorResponse(httpo.ItemDoesNotExist, err.Error())
		case pgerrcode.UniqueViolation:
			err = fmt.Errorf("some of fields already taken: %s", pgErr.ConstraintName)
			return httpo.NewErrorResponse(httpo.ItemAlreadyExist, err.Error())
		}
	}

	logo.Errorf("failed to %s: %s", action, err)
	return httpo.NewErrorResponse(500, "internal server error")
}
