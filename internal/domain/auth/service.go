package auth

type Service struct {
	tokenRepo RefreshTokenRepository
}

func NewService(tokenRepo RefreshTokenRepository) *Service {
	return &Service{tokenRepo: tokenRepo}
}

// Add methods for login, refresh, password reset, etc. in implementation
