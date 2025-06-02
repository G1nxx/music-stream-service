package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"music-stream-service/service/dtos/request"
	"music-stream-service/service/interfaces"
	e "music-stream-service/domain/entities"
	request_mappers "music-stream-service/service/mappers/request"
	"music-stream-service/service/repository"
	psh "github.com/vzglad-smerti/password_hash"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"user_id"`
}

type AuthService struct {
	interfaces.Authorization
	interfaces.TokenAuth
	repo          repository.AuthorizationRepository
	Log           *slog.Logger
	signingKey    []byte
	refreshSecret []byte
}

func (serv *AuthService) GetUserRole(userId int64) (e.UserRole, error) {
	user, err := serv.repo.GetUserById(userId)
	if err != nil {
		return 4, err
	}
	return user.Role, nil
}

func NewAuthService(repo repository.AuthorizationRepository, sl *slog.Logger) (*AuthService, error) {
	signingKey, err := generateRandomKey(32)
	if err != nil {
		return nil, err
	}

	refreshSecret, err := generateRandomKey(32)
	if err != nil {
		return nil, err
	}

	serv := &AuthService{
		repo:          repo,
		Log:           sl,
		signingKey:    signingKey,
		refreshSecret: refreshSecret,
	}
	sl.Debug("auth service successfully initiated")
	return serv, nil
}

func generateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func (serv *AuthService) AddUser(user request.UserAuthModel) error {
	userEntity, err := request_mappers.ToUserEntity(user)
	if err != nil {
		return err
	}
	return serv.repo.AddUser(*userEntity)
}

func (serv *AuthService) GenerateTokens(email, password string) (*interfaces.TokenPair, error) {
	pswd_hash, err := psh.Hash(password);
	if err != nil {
		return nil, err
	}
	user, err := serv.repo.GetUser(email, pswd_hash)
	if err != nil {
		return nil, err
	}

	return serv.GenerateTokensForUser(user.ID)
}

func (serv *AuthService) generateAccessToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userId,
	})
	return token.SignedString(serv.signingKey)
}

func (serv *AuthService) generateRefreshToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(userId))),
	})
	return token.SignedString(serv.refreshSecret)
}

func (serv *AuthService) ParseAccessToken(accessToken string) (int64, error) {
	claims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return serv.signingKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	return claims.UserId, nil
}

func (serv *AuthService) RefreshTokens(refreshToken string) (*interfaces.TokenPair, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return serv.refreshSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userIdBytes, err := base64.StdEncoding.DecodeString(claims.Subject)
	if err != nil {
		return nil, err
	}

	userId := int64(0)
	if _, err := fmt.Sscanf(string(userIdBytes), "%d", &userId); err != nil {
		return nil, err
	}

	if _, err := serv.repo.GetUserById(userId); err != nil {
		return nil, err
	}

	return serv.GenerateTokensForUser(userId)
}

func (serv *AuthService) GenerateTokensForUser(userId int64) (*interfaces.TokenPair, error) {
	accessToken, err := serv.generateAccessToken(userId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := serv.generateRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	return &interfaces.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().Add(accessTokenTTL).Unix(),
	}, nil
}
