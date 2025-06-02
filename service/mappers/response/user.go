package persistanceMappers

import (
	e "music-stream-service/domain/entities"
	persistance "music-stream-service/service/dtos/response"
)

func ToUserModel(usr e.User) (*persistance.UserModel, error) {
	user, err := persistance.NewUserModel(&usr)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ToArtistModel(usr e.User) (*persistance.ArtistModel, error) {
	author, err := persistance.NewArtistInListModel(&usr)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func ToUserProfileModel(usr e.User) (*persistance.UserModel, error) {
	user, err := persistance.NewUserModel(&usr)
	if err != nil {
		return nil, err
	}
	return user, nil
}
