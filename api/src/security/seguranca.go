package security

import "golang.org/x/crypto/bcrypt"

// Faz o hash de uma string
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// verifica se a senha informada é compatível
func VerificaSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
