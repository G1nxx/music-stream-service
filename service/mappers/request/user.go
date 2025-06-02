package persistanceMappers

import (
	e "music-stream-service/domain/entities"
	persistance "music-stream-service/service/dtos/request"
)

func ToUserEntity(usr persistance.UserAuthModel) (*e.User, error) {
	return e.NewUser(usr.Login, usr.Email, usr.Password)
}