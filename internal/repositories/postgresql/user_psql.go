package postgresql

import (
	"database/sql"
	_ "errors"
	"fmt"

	// persistance "main/repositories/postgres/entities_models"
	// persistanceMappers "main/repositories/postgres/mappers"
	"music-stream-service/service/repository"
)

type UserPostgres struct {
	repository.UserRepository
	DB *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{DB: db}
}

func (userRepo *UserPostgres) SubscribeToAlbum(userId, albumId int64) error {
	const op = "repositories.user_psql.subscribetoalbum"

	stmt, err := userRepo.DB.Prepare(`
    INSERT INTO users_albums(user_id, album_id) VALUES($1, $2);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(userId, albumId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (userRepo *UserPostgres) UnsubscribeFromAlbum(userId, albumId int64) error {
	const op = "repositories.user_psql.unsubscribefromalbum"

	stmt, err := userRepo.DB.Prepare(`
        DELETE FROM users_albums 
        WHERE user_id = $1 AND album_id = $2;
    `)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := stmt.Exec(userId, albumId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no subscription found for user %d to album %d", op, userId, albumId)
	}

	return nil
}

func (userRepo *UserPostgres) SubscribeToUser(first_userId, second_userId int64) error {
	const op = "repositories.user_psql.subscribetoalbum"

	stmt, err := userRepo.DB.Prepare(`
    INSERT INTO users_users(first_user_id, second_user_id) VALUES($1, $2);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(first_userId, second_userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (userRepo *UserPostgres) UnsubscribeFromUser(first_userId, second_userId int64) error {
	const op = "repositories.user_psql.unsubscribefromuser"

	stmt, err := userRepo.DB.Prepare(`
        DELETE FROM users_users 
        WHERE first_user_id = $1 AND second_user_id = $2;
    `)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := stmt.Exec(first_userId, second_userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no subscription found for user %d to user %d", op, first_userId, second_userId)
	}

	return nil
}

func (userRepo *UserPostgres) SubscribeToPlaylist(userId, playlistId int64) error {
	const op = "repositories.user_psql.subscribetoplaylist"

	stmt, err := userRepo.DB.Prepare(`
    INSERT INTO users_playlists(user_id, playlist_id) VALUES($1, $2);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(userId, playlistId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (userRepo *UserPostgres) UnsubscribeFromPlaylist(userId, playlistId int64) error {
	const op = "repositories.user_psql.unsubscribefromplaylist"

	stmt, err := userRepo.DB.Prepare(`
        DELETE FROM users_playlists 
        WHERE user_id = $1 AND playlist_id = $2;
    `)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := stmt.Exec(userId, playlistId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no subscription found for user %d to playlist %d", op, userId, playlistId)
	}

	return nil
}

func (userRepo *UserPostgres) AddTrackToPlaylist(trackId, playlistId int64) error {
	const op = "repositories.user_psql.addtracktoplaylist"

	stmt, err := userRepo.DB.Prepare(`
    INSERT INTO tracks_playlists(user_id, playlist_id) VALUES($1, $2);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(trackId, playlistId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (userRepo *UserPostgres) RemoveTrackFromPlaylist(trackId, playlistId int64) error {
	const op = "repositories.user_psql.removetrackfromplaylist"

	stmt, err := userRepo.DB.Prepare(`
        DELETE FROM tracks_playlists 
        WHERE track_id = $1 AND playlist_id = $2;
    `)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := stmt.Exec(trackId, playlistId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no subscription found for track %d to playlist %d", op, trackId, playlistId)
	}

	return nil
}
