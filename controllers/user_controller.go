package controllers

import (
	"encoding/json"
	"music-stream-service/controllers/middleware"
	controllerResponse "music-stream-service/controllers/response"
	"music-stream-service/service/dtos/request"
	"music-stream-service/service/interfaces"
	"net/http"
)

type UserController struct {
	service interfaces.UserActivity
	//tokenAuth  interfaces.TokenAuth
	middleware middleware.Middleware
}

func NewUserController(userService interfaces.UserActivity, mware middleware.Middleware) *UserController {
	return &UserController{
		service:    userService,
		middleware: mware,
	}
}

func (controller *UserController) SubscribeToUser(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToUserModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.SubscribeToUser(input.FirstUserID, input.SecondUserID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) UnsubscribeFromUser(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToUserModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.UnsubscribeFromUser(input.FirstUserID, input.SecondUserID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) SubscribeToAlbum(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToAlbumModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.SubscribeToAlbum(input.UserID, input.AlbumID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) UnsubscribeFromAlbum(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToAlbumModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.UnsubscribeFromAlbum(input.UserID, input.AlbumID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) SubscribeToPlaylist(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToPlaylistModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.SubscribeToPlaylist(input.UserID, input.PlaylistID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) UnsubscribeFromPlaylist(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.UserSubsToPlaylistModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.UnsubscribeFromPlaylist(input.UserID, input.PlaylistID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) AddTrackToPlaylist(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.TrackAddsToPlaylistModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.AddTrackToPlaylist(input.TrackID, input.PlaylistID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *UserController) RemoveTrackFromPlaylist(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.TrackAddsToPlaylistModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.RemoveTrackfromPlaylist(input.TrackID, input.PlaylistID)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}
