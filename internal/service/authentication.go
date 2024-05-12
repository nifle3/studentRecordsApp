package service

import (
	"context"
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
		Auth(ctx context.Context, role, email string) (uuid.UUID, string, error)
	}

	StudentAuthDB interface {
		Auth(ctx context.Context, email string) (uuid.UUID, string, error)
	}
)

type Auth struct {
	userAuth    UserAuthDB
	studentAuth StudentAuthDB
	secretKey   []byte
}

func NewAuth(secretKey []byte, studentDB StudentAuthDB, userDB UserAuthDB) Auth {
	return Auth{
		secretKey:   secretKey,
		studentAuth: studentDB,
		userAuth:    userDB,
	}
}

func (a Auth) Auth(ctx context.Context, role, email, pass string) (string, *customError.Http) {
	var hashPassword string
	var id uuid.UUID
	var err error

	if role == entities.UserWorker || role == entities.UserAdmin {
		id, hashPassword, err = a.userAuth.Auth(ctx, role, email)
	} else if role == entities.UserStudent {
		id, hashPassword, err = a.studentAuth.Auth(ctx, email)
	} else {
		return "", customError.New(http.StatusBadRequest, "Invalid role")
	}

	if err != nil {
		return "", customError.New(http.StatusUnauthorized, "Invalid login")
	}

	err = password.CheckHash(pass, []byte(hashPassword))
	if err != nil {
		return "", customError.New(http.StatusUnauthorized, "Invalid password")
	}

	result, cErr := a.generateToken(ctx, id, role)
	if cErr != nil {
		return "", customError.New(http.StatusInternalServerError, cErr.Error())
	}

	return result, nil
}

func (a Auth) generateToken(_ context.Context, id uuid.UUID, role string) (string, error) {
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

func (a Auth) ValidateRequireRole(ctx context.Context, token, role string) (uuid.UUID, *customError.Http) {
	userId, userRole, err := a.ValidateToken(ctx, token)
	if err != nil {
		return uuid.Nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	if userRole != role {
		return uuid.Nil, customError.New(http.StatusUnauthorized, fmt.Sprintf("Your role is not %s", role))
	}

	return userId, nil
}

func (a Auth) ValidateToken(_ context.Context, token string) (uuid.UUID, string, *customError.Http) {
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
