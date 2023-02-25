package supabase

import (
	"errors"
	"template-app/models/supabaseuser"
	"template-app/pkg/envconfig"

	"github.com/go-resty/resty/v2"
)

func GetSBUser(token string) (*supabaseuser.SBUser, error) {
	var headers = map[string]string{
		"accept": "application/json",
		"apiKey": envconfig.EnvVars.SUPABASE_SECRET_KEY,
	}

	var SBClient = resty.New().SetBaseURL(envconfig.EnvVars.SUPABASE_BASE_URL).SetHeaders(headers)
	r, e := SBClient.R().SetAuthToken(token).SetError(supabaseuser.SBError{}).SetResult(supabaseuser.SBUser{}).Get("auth/v1/user")
	if e != nil {
		return nil, e
	}

	if r.StatusCode() != 200 {
		err := r.Error().(*supabaseuser.SBError)
		return nil, errors.New(err.Msg)
	}

	sbu := r.Result().(*supabaseuser.SBUser)

	return sbu, nil
}
