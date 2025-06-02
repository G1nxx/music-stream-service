package entities

const (
	RoleAdmin  = 0
	RoleUser   = 1
	RoleArtist = 2
)

type UserRole = int8

var rolesList []UserRole = []UserRole{
	RoleAdmin,
	RoleUser,
	RoleArtist,
}

type User struct {
	ID       int64
	Login    string
	Email    string
	PswdHash string
	Role     UserRole
	PicPath  string
}

func NewUser(login string, email string, pswdHash string, options ...func(*User)) (*User, error) {
	usr := &User{}

	usr.Login = login
	usr.Email = email
	usr.PswdHash = pswdHash

	for _, opt := range options {
		opt(usr)
	}

	return usr, nil
}
