package interfaces

import (
	e "music-stream-service/domain/entities"
	"music-stream-service/service/dtos/request"
)


type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type RoleAccess interface {
	GetUserRole(userId int64) (e.UserRole, error)
}

type TokenAuth interface {
	GenerateTokens(email, password string) (*TokenPair, error)
	RefreshTokens(refreshToken string) (*TokenPair, error)
	ParseAccessToken(accessToken string) (int64, error)
	GenerateTokensForUser(userId int64) (*TokenPair, error)
	RoleAccess
}

type Authorization interface {
	AddUser(user request.UserAuthModel) error
	//GetUser(username, password string) (*response.UserAuthModel, error)
}
