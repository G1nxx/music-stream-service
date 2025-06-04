package controllers

import (
	"encoding/json"
	"fmt"
	"music-stream-service/controllers/middleware"
	controllerResponse "music-stream-service/controllers/response"
	"music-stream-service/service/dtos/request"
	"music-stream-service/service/interfaces"
	"net/http"
	"strconv"
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

func (controller *UserController) SubscribeToContent(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)

	var input request.UserSubsToContentModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	
	uId, err := strconv.ParseInt(input.UserID, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Wrong UserId type, must be number"}`, http.StatusBadRequest)
		return
	}
	cId, err := strconv.ParseInt(input.ContentID, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Wrong ContentId type, must be number"}`, http.StatusBadRequest)
		return
	}

	switch input.Type {
	case "Album":
		err = controller.service.SubscribeToAlbum(uId, cId)
	case "Artist":
		err = controller.service.SubscribeToUser(uId, cId)
	case "Playlist":
		err = controller.service.SubscribeToPlaylist(uId, cId)
	default:
		err = fmt.Errorf("%s", "Wrong subscription content type")
	}
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	controllerResponse.OkResponse(writer)
}

func (controller *UserController) UnsubscribeFromContent(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)

	var input request.UserSubsToContentModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	uId, err := strconv.ParseInt(input.UserID, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Wrong UserId type, must be number"}`, http.StatusBadRequest)
		return
	}
	cId, err := strconv.ParseInt(input.ContentID, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Wrong ContentId type, must be number"}`, http.StatusBadRequest)
		return
	}

	switch input.Type {
	case "Album":
		err = controller.service.UnsubscribeFromAlbum(uId, cId)
	case "Artist":
		err = controller.service.UnsubscribeFromUser(uId, cId)
	case "Playlist":
		err = controller.service.UnsubscribeFromPlaylist(uId, cId)
	default:
		err = fmt.Errorf("%s", "Wrong subscription content type")
	}
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
