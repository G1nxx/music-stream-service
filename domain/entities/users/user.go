package users

type User struct {
	Id       int64
	Login    string
	Email    string
	PswdHash string
}
