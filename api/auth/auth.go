// Package auth provides Api to authenticate user by validating his supabase jwt and giving him paseto
package auth

import (
	"errors"
	"net/http"
	"strings"
	"template-app/models/company/companymethods"
	"template-app/pkg/paseto"
	"template-app/pkg/supabase"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.POST("", supabaseAuth)
	}
}

func supabaseAuth(c *gin.Context) {
	var body AuthRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}
	// If unexpected error
	sbUser, err := supabase.GetSBUser(body.SupabaseToken)
	if err != nil {
		logo.Errorf("failed to get supabase user: %s", err)
		errMsg := "failed to verify and get paseto"
		if strings.Contains(err.Error(), "invalid JWT: unable to parse or verify signature, token is expired by ") {
			errMsg += ": supabase JWT token expired"
		}
		httpo.NewErrorResponse(500, errMsg).Send(c, 500)
		return
	}

	_, err = companymethods.Get(sbUser.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = companymethods.Add(sbUser.Email)
			if err != nil {
				logo.Errorf("failed to add user into database: %s", err)
				httpo.NewErrorResponse(500, "failed to verify and get paseto").Send(c, 500)
				return
			}
		} else {
			logo.Errorf("failed to check if user exist in database: %s", err)
			httpo.NewErrorResponse(500, "failed to verify and get paseto").Send(c, 500)
			return
		}
	}

	pasetoToken, err := paseto.GetPasetoForUser(sbUser.Email)
	if err != nil {
		logo.Errorf("failed to get paseto: %s", err)

		httpo.NewErrorResponse(500, "failed to verify and get paseto").Send(c, 500)
		return
	}

	payload := AuthPayload{
		Token: pasetoToken,
	}
	httpo.NewSuccessResponseP(http.StatusOK, "token generated successfully", payload).
		SendD(c)
}
