package postgresql

import (
	"database/sql"
	_ "errors"
	"fmt"
	e "music-stream-service/domain/entities"
	// persistance "main/repositories/postgres/entities_models"
	// persistanceMappers "main/repositories/postgres/mappers"
	"music-stream-service/service/repository"
)

type AuthPostgres struct {
	repository.AuthorizationRepository
	DB *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{DB: db}
}

func (authRepo *AuthPostgres) AddUser(user e.User) error {
	const op = "repositories.auth_psql.adduser"

	stmt, err := authRepo.DB.Prepare(`
    INSERT INTO users(login, email, pswd_hash, role) VALUES($1, $2, $3, 1);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(user.Login, user.Email, user.PswdHash)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// func (authRepo *AuthPostgres) GetUser(username, password string) (*users.User, error) {
// 	const op = "repositories.postgresql.getuser"
// }