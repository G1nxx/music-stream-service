package controllers

import (
	"encoding/json"
	"fmt"
	"music-stream-service/controllers/middleware"
	controllerResponse "music-stream-service/controllers/response"
	"music-stream-service/service/dtos/request"
	"music-stream-service/service/dtos/response"
	"music-stream-service/service/interfaces"
	"net/http"
	"strconv"
)

type MusicInfoController struct {
	service interfaces.MusicInfoActivity
	//tokenAuth  interfaces.TokenAuth
	middleware middleware.Middleware
}

func NewMusicInfoController(musicInfoService interfaces.MusicInfoActivity, mware middleware.Middleware) *MusicInfoController {
	return &MusicInfoController{
		service:    musicInfoService,
		middleware: mware,
	}
}

func (controller *MusicInfoController) GetAllSubscriptionsInListByFilter(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.InfoModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	switch fl := input.Filter; fl {
	case "Albums":
		albs, err := controller.service.GetAllAlbumsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		jsonData, err := json.Marshal(albs)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal albums data")
			return
		}

		controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)

	case "Playlists":
		plsts, err := controller.service.GetAllPlaylistsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		jsonData, err := json.Marshal(plsts)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal playlists data")
			return
		}

		controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)

	case "Artists":
		artsts, err := controller.service.GetAllArtistsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		jsonData, err := json.Marshal(artsts)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal artists data")
			return
		}

		controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)

	case "All":
		albs, err := controller.service.GetAllAlbumsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		plsts, err := controller.service.GetAllPlaylistsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		artsts, err := controller.service.GetAllArtistsInList(input.UserID)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}

		var combined []any
		for _, pl := range plsts {
			combined = append(combined, pl)
		}
		for _, ar := range artsts {
			combined = append(combined, ar)
		}
		for _, al := range albs {
			combined = append(combined, al)
		}

		jsonData, err := json.Marshal(combined)
		if err != nil {
			controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal albums data")
			return
		}

		controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)

	default:
		err = fmt.Errorf("wrong subscriptions filter")
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
}

func (controller *MusicInfoController) GetAlbumInfoByAlbumId(writer http.ResponseWriter, req *http.Request) {
	var r request.TracksFromAlbumModel

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(writer, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if r.AlbumId == "" {
		http.Error(writer, `{"error": "AlbumId is required"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(r.AlbumId, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Vrong AlbumId type, must be number"}`, http.StatusBadRequest)
		return
	}

	tracks, err := controller.service.GetTracksFromAlbum(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	albumInfo, err := controller.service.GetAlbumInfo((id))
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	type AlbumResponse struct {
		Album  response.AlbumViewModel     `json:"album"`
		Tracks []response.TrackInSubsModel `json:"tracks"`
	}

	resp := AlbumResponse{
		Album:  *albumInfo,
		Tracks: tracks,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal tracks_in_album data")
		return
	}

	controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)
}

func (controller *MusicInfoController) GetPlaylistInfoByPlaylistId(writer http.ResponseWriter, req *http.Request) {
	var r request.TracksFromPlaylistModel

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(writer, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if r.PlaylistId == "" {
		http.Error(writer, `{"error": "PlaylistId is required"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(r.PlaylistId, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Vrong PlaylistId type, must be number"}`, http.StatusBadRequest)
		return
	}

	tracks, err := controller.service.GetTracksFromPlaylist(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	playlistInfo, err := controller.service.GetPlaylistInfo(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	playlistInfo.Saves, err = controller.service.GetPlaylistSaves(id)
	if err != nil {
		// http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		// return
	}

	type PlaylistResponse struct {
		Playlist response.PlaylistViewModel  `json:"playlist"`
		Tracks   []response.TrackInSubsModel `json:"tracks"`
	}

	resp := PlaylistResponse{
		Playlist: *playlistInfo,
		Tracks:   tracks,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal tracks_in_playlist data")
		return
	}

	controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)
}

func (controller *MusicInfoController) GetArtistInfoByArtistId(writer http.ResponseWriter, req *http.Request) {
	var r request.ReleasesFromArtistModel

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(writer, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if r.ArtistId == "" {
		http.Error(writer, `{"error": "ArtistId is required"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(r.ArtistId, 10, 64)
	if err != nil {
		http.Error(writer, `{"error": "Vrong ArtistId type, must be number"}`, http.StatusBadRequest)
		return
	}

	releases, err := controller.service.GetReleasesFromArtist(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	artistInfo, err := controller.service.GetArtistInfo(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	attachedPlaylist, err := controller.service.GetArtistAttachment(id)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	tracks, err := controller.service.GetTracksFromPlaylist(attachedPlaylist.ID)
	if err != nil {
		http.Error(writer, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	type ArtistResponse struct {
		Artist     response.ArtistModel        `json:"artist"`
		Releases   []response.AlbumInListModel `json:"releases"`
		Attachment response.PlaylistViewModel  `json:"attached"`
		Tracks     []response.TrackInSubsModel `json:"tracks"`
	}

	resp := ArtistResponse{
		Artist:     *artistInfo,
		Releases:   releases,
		Attachment: *attachedPlaylist,
		Tracks:     tracks,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, "failed to marshal tracks_in_artist data")
		return
	}

	controllerResponse.InfoResponse(writer, http.StatusOK, jsonData)
}
