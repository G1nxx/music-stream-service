package postgresql

import (
	"database/sql"
	"fmt"
	"music-stream-service/internal/config"

	//rep "music-stream-service/service/repository"

	_ "github.com/lib/pq"
)


func New(dbServer config.DBServer) (*sql.DB, error) {
	const op = "storage.postgresql.new"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbServer.Host, dbServer.Port, dbServer.Username, 
		dbServer.Password, dbServer.DBName, dbServer.SSLMode)

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

	// err = newTableAuthors(db)
	// if err != nil {
	// 	return nil, err
	// }

	err = newTableAlbums(db)
	if err != nil {
		return nil, err
	}

	err = newTableUsersToAlbums(db)
	if err != nil {
		return nil, err
	}

	err = newTableUsersToUsers(db)
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

	return db, nil
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
	`)
	if err != nil {
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
	`)
	if err != nil {
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
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// func newTableUsersToAuthors(db *sql.DB) error {
// 	const op = "storage.postgresql.newtableuserstoauthors"

// 	stmt, err := db.Prepare(`
// 	CREATE TABLE IF NOT EXISTS users_authors (
// 		author_id BIGINT NOT NULL,
// 		user_id BIGINT NOT NULL,
// 		PRIMARY KEY (author_id, user_id),
// 		FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE,
// 		FOREIGN KEY (user_id)   REFERENCES users(id)   ON DELETE CASCADE
// 	);
// 	`)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	_, err = stmt.Exec()
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	return nil
// }

func newTableUsersToUsers(db *sql.DB) error {
	const op = "storage.postgresql.newtableuserstousers"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users_users (
		first_user_id  BIGINT NOT NULL,
		second_user_id BIGINT NOT NULL,
		added_at       TIMESTAMP,
		PRIMARY KEY (first_user_id, second_user_id),
		FOREIGN KEY (first_user_id)  REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (second_user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// func newTableAuthors(db *sql.DB) error {
// 	const op = "storage.postgresql.newtableauthors"

// 	stmt, err := db.Prepare(`
// 	CREATE TABLE IF NOT EXISTS "authors" (
// 		id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
// 		user_id		   INTEGER REFERENCES users(id) ON DELETE CASCADE,
// 		verified	   BOOLEAN
// 	);
// 	`)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	_, err = stmt.Exec()
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	return nil
// }

func newTableAlbums(db *sql.DB) error {
	const op = "storage.postgresql.newtablealbums"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "albums" (
		id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name           TEXT    NOT NULL CHECK (name <> ''),
		author_id 	   INTEGER REFERENCES users(id) ON DELETE CASCADE,
		genre          TEXT,
		length_sec     INTEGER,
		recording_date TIMESTAMP,
		cover_path     TEXT    NOT NULL,
		type           TEXT
	);
	`)
	if err != nil {
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
		name           TEXT    NOT NULL CHECK (name <> ''),
		album_id	   INTEGER REFERENCES albums(id) ON DELETE CASCADE,
		author_id 	   INTEGER REFERENCES users(id)  ON DELETE CASCADE,
		genre          TEXT,
		length_sec     INTEGER,
		recording_date TIMESTAMP,
		size_in_bytes  INTEGER,
		number         SMALLINT,
		format         TEXT,
		path           TEXT    NOT NULL CHECK (path <> ''),
		cover_path     TEXT
	);
	`)
	if err != nil {
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
		id        INTEGER 	PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		login 	  TEXT    	NOT NULL CHECK (login <> ''),
		email 	  TEXT    	NOT NULL UNIQUE CHECK (email <> ''),	
		pswd_hash TEXT	  	NOT NULL CHECK (pswd_hash <> ''),
		role	  SMALLINT,
		pic_path  TEXT
	);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func newTablePlaylists(db *sql.DB) error {
	const op = "storage.postgresql.newtableplaylists"

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "playlists" (
		id        		INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name 	  		TEXT    NOT NULL CHECK (name <> ''),
		creator_id 		INTEGER REFERENCES users(id) ON DELETE CASCADE,
		length_sec      INTEGER,
		creation_time 	TIMESTAMP,
		cover_path		TEXT,
		attached_to     INTEGER REFERENCES users(id) ON DELETE CASCADE,
		description 	TEXT 	NOT NULL
	);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
