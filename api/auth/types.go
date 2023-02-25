package auth

type AuthRequest struct {
	SupabaseToken string `json:"supabaseToken" binding:"required"`
}

type AuthPayload struct {
	Token string `json:"token"`
}
