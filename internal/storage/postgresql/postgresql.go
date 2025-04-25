package postgresql

import (
	"database/sql"
	"fmt"
	e "music-stream-service/domain/entities/users"

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

	err = newTableUsers(db)
	if err != nil {
		return nil, err
	}

	err = newTablePlaylists(db)
	if err != nil {
		return nil, err
	}

	err = newTableAuthors(db)
	if err != nil {
		return nil, err
	}

	err = newTableAlbums(db)
	if err != nil {
		return nil, err
	}

	err = newTableUsersToAlbums(db)
	if err != nil {
		return nil, err
	}

	err = newTableUsersToAuthors(db)
	if err != nil {
		return nil, err
	}

	err = newTableTraks(db)
	if err != nil {
		return nil, err
	}

	err = newTableTracksToPlaylists(db)
	if err != nil {
		return nil, err
	}

	err = newTableUsersToPlaylists(db)
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}

func newTableTracksToPlaylists(db *sql.DB) error {
	const op = "storage.postgresql.newtabletrackstoplaylists"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS tracks_playlists (
		playlist_id BIGINT NOT NULL,
		track_id BIGINT NOT NULL,
		PRIMARY KEY (playlist_id, track_id),
		FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
		FOREIGN KEY (track_id)	  REFERENCES tracks(id)	   ON DELETE CASCADE
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableUsersToPlaylists(db *sql.DB) error {
	const op = "storage.postgresql.newtableuserstoplaylists"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users_playlists (
		playlist_id BIGINT NOT NULL,
		user_id BIGINT NOT NULL,
		PRIMARY KEY (playlist_id, user_id),
		FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id)	  REFERENCES users(id)	   ON DELETE CASCADE
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableUsersToAlbums(db *sql.DB) error {
	const op = "storage.postgresql.newtableuserstoalbums"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users_albums (
		album_id BIGINT NOT NULL,
		user_id BIGINT NOT NULL,
		PRIMARY KEY (album_id, user_id),
		FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id)  REFERENCES users(id)  ON DELETE CASCADE
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableUsersToAuthors(db *sql.DB) error {
	const op = "storage.postgresql.newtableuserstoauthors"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users_authors (
		author_id BIGINT NOT NULL,
		user_id BIGINT NOT NULL,
		PRIMARY KEY (author_id, user_id),
		FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id)   REFERENCES users(id)   ON DELETE CASCADE
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableAuthors(db *sql.DB) error {
	const op = "storage.postgresql.newtableauthors"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "authors" (
		id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		user_id		   INTEGER REFERENCES users(id) ON DELETE CASCADE,
		verified	   BOOLEAN
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableAlbums(db *sql.DB) error {
	const op = "storage.postgresql.newtablealbums"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "albums" (
		id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name           TEXT    NOT NULL,
		author_id 	   INTEGER REFERENCES  authors(id) ON DELETE CASCADE,
		genre          TEXT,
		length_sec     INTEGER,
		recording_date TIMESTAMP,
		cover_path     TEXT    NOT NULL,
		type           TEXT
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableTraks(db *sql.DB) error {
	const op = "storage.postgresql.newtabletraks"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "tracks" (
		id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name           TEXT    NOT NULL,
		album_id	   INTEGER REFERENCES albums(id)  ON DELETE CASCADE,
		author_id 	   INTEGER REFERENCES  authors(id) ON DELETE CASCADE,
		genre          TEXT,
		length_sec     INTEGER,
		recording_date TIMESTAMP,
		size_in_bytes  INTEGER,
		number         SMALLINT,
		format         TEXT,
		path           TEXT    NOT NULL,
		cover_path     TEXT    NOT NULL
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTableUsers(db *sql.DB) error {
	const op = "storage.postgresql.newtableusers"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "users" (
		id        INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		login 	  TEXT    NOT NULL,
		email 	  TEXT    NOT NULL UNIQUE,
		pswd_hash TEXT	  NOT NULL
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTablePlaylists(db *sql.DB) error {
	const op = "storage.postgresql.newtableusers"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "playlists" (
		id        		INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name 	  		TEXT    NOT NULL,
		discription 	TEXT    NOT NULL,
		creator_id 		INTEGER REFERENCES users(id) ON DELETE CASCADE,
		creation_time 	TIMESTAMP
	);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) CreateUser(user e.User) error {
	const op = "storage.postgresql.createuser"

	stmt, err := s.DB.Prepare(`
    INSERT INTO users(id, login, email, paswdhash) VALUES($1, $2, $3, $4);
	`);	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(user.Id, user.Login, user.Email, user.PswdHash)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
