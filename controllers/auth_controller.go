package controllers

import (
	"encoding/json"
	"music-stream-service/controllers/middleware"
	controllerResponse "music-stream-service/controllers/response"
	"music-stream-service/service/dtos/request"
	"music-stream-service/service/interfaces"
	"net/http"

	psh "github.com/vzglad-smerti/password_hash"
)

type AuthController struct {
	service    interfaces.Authorization
	//tokenAuth  interfaces.TokenAuth
	middleware middleware.Middleware
}

func NewAuthController(authService interfaces.Authorization, mware middleware.Middleware) *AuthController {
	return &AuthController{
		service: authService,
		middleware: mware,
	}
}

func (controller *AuthController) signUp(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserAuthModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	input.Password, err = psh.Hash(input.Password)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = controller.service.AddUser(input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}
