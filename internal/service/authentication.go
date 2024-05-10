package service

import (
	"fmt"
	"net/http"
	"studentRecordsApp/pkg/customError"
	"studentRecordsApp/pkg/password"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
)

type claims struct {
	Id   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	jwt.RegisteredClaims
}

type (
	UserAuthDB interface {
		Auth(role, email string) (uuid.UUID, string, error)
	}

	StudentAuthDB interface {
		Auth(email string) (uuid.UUID, string, error)
	}
)

type Auth struct {
	userAuth    UserAuthDB
	studentAuth StudentAuthDB
	secretKey   []byte
}

func NewAuth(secretKey []byte, studentDB StudentAuthDB, userDB UserAuthDB) *Auth {
	return &Auth{
		secretKey:   secretKey,
		studentAuth: studentDB,
		userAuth:    userDB,
	}
}

func (a Auth) Auth(role, email, pass string) (string, error) {
	var hashPassword string
	var id uuid.UUID
	var err error

	if role == entities.UserWorker || role == entities.UserAdmin {
		id, hashPassword, err = a.userAuth.Auth(role, email)
	} else {
		id, hashPassword, err = a.studentAuth.Auth(email)
	}

	if err != nil {
		return "", customError.New(http.StatusUnauthorized, "Invalid login")
	}

	err = password.CheckHash(pass, []byte(hashPassword))
	if err != nil {
		return "", customError.New(http.StatusUnauthorized, "Invalid password")
	}

	return a.generateToken(id, role)
}

func (a Auth) generateToken(id uuid.UUID, role string) (string, error) {
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

func (a Auth) ValidateRequireRole(token, role string) (uuid.UUID, error) {
	userId, userRole, err := a.ValidateToken(token)
	if err != nil {
		return uuid.Nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	if userRole != role {
		return uuid.Nil, customError.New(http.StatusUnauthorized, fmt.Sprintf("Your role is not %s", role))
	}

	return userId, nil
}

func (a Auth) ValidateToken(token string) (uuid.UUID, string, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if err != nil {
		return uuid.Nil, "", customError.New(http.StatusInternalServerError, err.Error())
	}

	claim, ok := jwtToken.Claims.(*claims)
	if !ok {
		return uuid.Nil, "", customError.New(http.StatusUnauthorized, "Uknown claims")
	}

	return claim.Id, claim.Role, nil
}
