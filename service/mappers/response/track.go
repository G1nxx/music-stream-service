package persistanceMappers

import (
	e "music-stream-service/domain/entities"
	persistance "music-stream-service/service/dtos/response"
)

func ToTrackModel(trk e.Track) (*persistance.TrackModel, error) {
	track, err := persistance.NewTrackModel(&trk)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func ToTrackInSubsModel(trk e.Track) (*persistance.TrackInSubsModel, error) {
	track, err := persistance.NewTrackInAlbumModel(&trk)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func ToTrackPlayingModel(trk e.Track) (*persistance.TrackPlayingModel, error) {
	track, err := persistance.NewTrackPlayingModel(&trk)
	if err != nil {
		return nil, err
	}
	return track, nil
}
