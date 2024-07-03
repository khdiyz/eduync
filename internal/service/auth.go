package service

import (
	"database/sql"
	"edusync/internal/config"
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/helper"
	"edusync/pkg/logger"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
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
	UserId int64  `json:"user_id"`
	RoleId int64  `json:"role_id"`
	Type   string `json:"type"`
}

func (s *AuthService) CreateToken(user model.User, tokenType string, expiresAt time.Time) (*model.Token, error) {
	claims := &jwtCustomClaim{
		UserId: user.Id,
		RoleId: user.RoleId,
		Type:   tokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return &model.Token{
		User:      user,
		Token:     token,
		Type:      tokenType,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GenerateTokens(user model.User) (*model.Token, *model.Token, error) {
	accessExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTAccessExpirationHours) * time.Hour)
	refreshExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTRefreshExpirationDays) * time.Hour * 24)

	accessToken, err := s.CreateToken(user, constants.TokenTypeAccess, accessExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.CreateToken(user, constants.TokenTypeRefresh, refreshExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ParseToken(token string) (*jwtCustomClaim, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, errors.New("token claims are not of type *jwtCustomClaim")
	}

	return claims, nil
}

func (s *AuthService) Login(input model.LoginRequest) (*model.Token, *model.Token, error) {
	user, err := s.repo.GetByUsername(input.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, serviceError(errors.New("wrong username or password"), codes.Unauthenticated)
		}
		return nil, nil, serviceError(err, codes.Internal)
	}

	hashPassword, err := helper.GenerateHash(input.Password)
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	if user.Password != hashPassword {
		return nil, nil, serviceError(errors.New("wrong username or password"), codes.Unauthenticated)
	}

	return s.GenerateTokens(user)
}
