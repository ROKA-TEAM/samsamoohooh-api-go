package port

type AuthService interface {
	GenerateSecureToken(length int) (string, error)
}
