package supabaseuser

type SBUser struct {
	ID       string       `json:"sub"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Role     string       `json:"role"`
	UserData UserMetaData `json:"user_metadata"`
}

type SBError struct {
	Code int
	Msg  string
}

type UserMetaData struct {
	AvatarURL     string `json:"avatar_url"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FullName      string `json:"full_name"`
	ProviderID    string `json:"provider_id"`
	Picture       string `json:"picture"`
}
