package middleware

import (
	"errors"
	"music-stream-service/domain/entities"
	serviceInterfaces "music-stream-service/service/interfaces"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

type Middleware struct {
	authMiddleware serviceInterfaces.TokenAuth
}

func NewMiddleware(authMiddleware serviceInterfaces.TokenAuth) *Middleware {
	return &Middleware{
		authMiddleware: authMiddleware,
	}
}

func (middleware *Middleware) UserIdentity(req *http.Request) (int64, error) {
	header := req.Header.Get(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid header length")
	}
	userId, err := middleware.authMiddleware.ParseAccessToken(headerParts[1])
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (middleware *Middleware) GetUserRole(req *http.Request) (entities.UserRole, error) {
	userId, err := middleware.UserIdentity(req)
	if err != nil {
		return 4,err
	}
	role, err :=  middleware.authMiddleware.GetUserRole(userId)
	if err != nil {
		return 4, err
	}
	return role, nil
}

func (middleware *Middleware) EnableCors(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3243")
    writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
    writer.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}