package auth

type AuthRequest struct {
	SupabaseToken string `json:"supabaseToken"`
}

type AuthPayload struct {
	Token string `json:"token"`
}
