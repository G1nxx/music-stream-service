package response

import e "music-stream-service/domain/entities"

type UserModel struct {
	Id      int64  `json:"id" binding:"required"`
	Login   string `json:"login" binding:"required"`
	PicPath string `json:"artwork"`
}

func NewUserModel(user *e.User, options ...func(*UserModel) (*UserModel, error)) (*UserModel, error) {
	usr := &UserModel{}

	usr.Id = user.ID
	usr.Login = user.Login
	usr.PicPath = user.PicPath

	for _, opt := range options {
		opt(usr)
	}

	return usr, nil
}

type ArtistModel struct {
	Id      int64  `json:"id" binding:"required"`
	Login   string `json:"title" binding:"required"`
	PicPath string `json:"artwork"`
	Type    string `json:"type" json-default:"Artist"`
}

func NewArtistInListModel(user *e.User, options ...func(*ArtistModel) (*ArtistModel, error)) (*ArtistModel, error) {
	usr := &ArtistModel{}

	usr.Id = user.ID
	usr.Login = user.Login
	usr.PicPath = user.PicPath
	usr.Type = "Artist"

	for _, opt := range options {
		opt(usr)
	}

	return usr, nil
}

type UserProfileModel struct {
	UserModel
	Email string `json:"email" binding:"required"`
}

func NewUserProgileModel(user *e.User, options ...func(*UserProfileModel) (*UserProfileModel, error)) (*UserProfileModel, error) {
	usr := &UserProfileModel{}

	usr.Id = user.ID
	usr.Login = user.Login
	usr.PicPath = user.PicPath
	usr.Email = user.Email

	for _, opt := range options {
		opt(usr)
	}

	return usr, nil
}
