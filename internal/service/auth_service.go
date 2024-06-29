package service

import (
	"database/sql"
	"edusync/internal/config"
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/helper"
	"edusync/pkg/logger"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo   repository.UserRepo
	logger logger.Logger
}

func NewAuthService(repo repository.UserRepo, logger logger.Logger) *AuthService {
	return &AuthService{
		repo:   repo,
		logger: logger,
	}
}

type jwtCustomClaim struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}

func (s *AuthService) GenerateToken(input model.UserLoginReq) (token string, err error) {
	input.Password, err = helper.GenerateHash(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.UserReader.GetByUsernameAndPassword(input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("wrong username or password")
		}
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtCustomClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.GetConfig().JWTExpired)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.RoleId,
	})

	return jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
}

func (s *AuthService) ParseToken(accessToken string) (*jwtCustomClaim, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, errors.New("token claims are not of type *jwtCustomClaim")
	}

	return claims, nil
}
