package postgresql

import (
	"database/sql"
	"fmt"
	e "music-stream-service/domain/entities"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(psqlInfo string) (*Storage, error) {
	const op = "storage.postgresql.new"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "users" (
		id        SERIAL PRIMARY KEY,
		login 	  TEXT   NOT NULL,
		email 	  TEXT   NOT NULL UNIQUE,
		paswdhash TEXT	 NOT NULL
	);
	`)
	// stmt, err := db.Prepare(`
	// DROP TABLE IF EXISTS "users"
	// `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{DB: db}, nil
}

// func (s *Storage) SaveUser(user e.User) (int64, error) {
// 	const op = "storage.postgresql.saveuser"

// 	stmt, err := s.DB.Prepare(`
//     INSERT INTO users(id, login, email, paswdhash) VALUES($1, $2, $3, $4);
// 	`)
// 	if err != nil {
// 		return 0, fmt.Errorf("%s: %w", op, err)
// 	}

// 	res, err := stmt.Exec(user.Id, user.Login, user.Email, user.Paswdhash)
// 	if err != nil {
// 		return 0, fmt.Errorf("%s: %w", op, err)
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return 0, fmt.Errorf("%s: %w", op, err)
// 	}

// 	return id, nil
// }

func (s *Storage) CreateUser(user e.User) (error) {
	const op = "storage.postgresql.saveuser"

	stmt, err := s.DB.Prepare(`
    INSERT INTO users(id, login, email, paswdhash) VALUES($1, $2, $3, $4);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(user.Id, user.Login, user.Email, user.Paswdhash)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
