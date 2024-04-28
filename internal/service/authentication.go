package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type claims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type Auth struct {
	secretKey []byte
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) GenerateToken(id string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(12 * time.Hour),
			},
		},
	})

	return token.SignedString(a.secretKey)
}

func (a Auth) ValidateTokenWithRequireRole(token, role string) (string, error) {
	userId, userRole, err := a.ValidateToken(token)
	if err != nil {
		return "", err
	}

	if userRole != role {
		return "", fmt.Errorf("401 user role is not %s", role)
	}

	return userId, nil
}

func (a Auth) ValidateToken(token string) (string, string, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if err != nil {
		return "", "", fmt.Errorf("400 %s", err)
	}

	claim, ok := jwtToken.Claims.(*claims)
	if !ok {
		return "", "", fmt.Errorf("400 uknown claims")
	}

	return claim.Id, claim.Role, nil
}
