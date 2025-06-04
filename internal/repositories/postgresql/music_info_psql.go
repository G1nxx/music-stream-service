package postgresql

import (
	"database/sql"
	_ "errors"
	"fmt"
	e "music-stream-service/domain/entities"
	"sort"
	"time"

	// persistance "main/repositories/postgres/entities_models"
	// persistanceMappers "main/repositories/postgres/mappers"
	"music-stream-service/service/repository"

	"github.com/lib/pq"
)

type MusicInfoPostgres struct {
	repository.MusicInfoRepository
	DB *sql.DB
}

func NewMusicInfoPostgres(db *sql.DB) *MusicInfoPostgres {
	return &MusicInfoPostgres{DB: db}
}

func (miRepo *MusicInfoPostgres) GetAllAlbums(userId int64) ([]e.Album, error) {
	const op = "repositories.musicinfo_psql.GetAllAlbums"

	albumIdsQuery := `
        SELECT album_id 
        FROM users_albums 
        WHERE user_id = $1
    `

	rows, err := miRepo.DB.Query(albumIdsQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get user albums: %w", op, err)
	}
	defer rows.Close()

	var albumIDs []int64
	for rows.Next() {
		var albumID int64
		if err := rows.Scan(&albumID); err != nil {
			return nil, fmt.Errorf("%s: failed to scan album id: %w", op, err)
		}
		albumIDs = append(albumIDs, albumID)
	}

	if len(albumIDs) == 0 {
		return []e.Album{}, nil
	}

	albumsQuery := `
        SELECT id, name, author_id, genre, length_sec, recording_date, cover_path, type
        FROM albums 
        WHERE id = ANY($1::BIGINT[])
    `

	albumRows, err := miRepo.DB.Query(albumsQuery, pq.Array(albumIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get albums info: %w", op, err)
	}
	defer albumRows.Close()

	var albums []e.Album = make([]e.Album, 0, len(albumIDs))
	for albumRows.Next() {
		var album e.Album
		err := albumRows.Scan(
			&album.ID,
			&album.Name,
			&album.AuthorID,
			&album.Genre,
			&album.Length,
			&album.RecordingDate,
			&album.CoverPath,
			&album.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan album: %w", op, err)
		}

		authorQuery := `
			SELECT login 
			FROM users 
			WHERE id = $1
    	`
		var authorName string
		err = miRepo.DB.QueryRow(authorQuery, album.AuthorID).Scan(&authorName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
		}
		album.AuthorName = authorName
		albums = append(albums, album)
	}
	return albums, nil
}

func (miRepo *MusicInfoPostgres) GetAllArtists(userId int64) ([]e.User, error) {
	const op = "repositories.musicinfo_psql.GetAllArtists"

	artistIdsQuery := `
        SELECT second_user_id 
        FROM users_users 
        WHERE first_user_id = $1
    `

	rows, err := miRepo.DB.Query(artistIdsQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get user artists: %w", op, err)
	}
	defer rows.Close()

	var artistsIDs []int64
	for rows.Next() {
		var artistID int64
		if err := rows.Scan(&artistID); err != nil {
			return nil, fmt.Errorf("%s: failed to scan artist id: %w", op, err)
		}
		artistsIDs = append(artistsIDs, artistID)
	}

	if len(artistsIDs) == 0 {
		return []e.User{}, nil
	}

	artistsQuery := `
        SELECT id, login, email, pswd_hash, role, pic_path
        FROM users 
        WHERE id = ANY($1::BIGINT[])
    `

	artistsRows, err := miRepo.DB.Query(artistsQuery, pq.Array(artistsIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get artists info: %w", op, err)
	}
	defer artistsRows.Close()

	var artists []e.User = make([]e.User, 0, len(artistsIDs))
	for artistsRows.Next() {
		var artist e.User
		err := artistsRows.Scan(
			&artist.ID,
			&artist.Login,
			&artist.Email,
			&artist.PswdHash,
			&artist.Role,
			&artist.PicPath,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan artist: %w", op, err)
		}

		if artist.Role == e.RoleArtist {
			artists = append(artists, artist)
		}
	}
	return artists, nil
}

func (miRepo *MusicInfoPostgres) GetAllPlaylists(userId int64) ([]e.Playlist, error) {
	const op = "repositories.musicinfo_psql.GetAllPlaylists"

	playistIdsQuery := `
        SELECT playlist_id 
        FROM users_playlists 
        WHERE user_id = $1
    `

	rows, err := miRepo.DB.Query(playistIdsQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get user playlists: %w", op, err)
	}
	defer rows.Close()

	var playlistIDs []int64
	for rows.Next() {
		var playlistID int64
		if err := rows.Scan(&playlistID); err != nil {
			return nil, fmt.Errorf("%s: failed to scan playlist id: %w", op, err)
		}
		playlistIDs = append(playlistIDs, playlistID)
	}

	if len(playlistIDs) == 0 {
		return []e.Playlist{}, nil
	}

	playlistsQuery := `
        SELECT id, name, creator_id, length_sec, creation_time, cover_path
        FROM playlists 
        WHERE id = ANY($1::BIGINT[])
    `

	playlistRows, err := miRepo.DB.Query(playlistsQuery, pq.Array(playlistIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get playlists info: %w", op, err)
	}
	defer playlistRows.Close()

	var playlists []e.Playlist = make([]e.Playlist, 0, len(playlistIDs))
	for playlistRows.Next() {
		var playlist e.Playlist
		err := playlistRows.Scan(
			&playlist.ID,
			&playlist.Name,
			&playlist.CreatorID,
			&playlist.Length,
			&playlist.CreationTime,
			&playlist.CoverPath,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan playlist: %w", op, err)
		}

		if playlist.Name == "Liked Songs" {
			var count int64
			err := miRepo.DB.QueryRow(`
				SELECT COUNT(track_id) FROM tracks_playlists WHERE playlist_id = $1
			`, playlist.ID).Scan(&count)
			if err != nil {
				playlist.CreatorName = "♡"
			} else {
				playlist.CreatorName = fmt.Sprintf("%d song%s", count, map[bool]string{true: "", false: "s"}[count == 1])
			}
		} else {
			authorQuery := `
				SELECT login 
				FROM users 
				WHERE id = $1
			`

			var creatorName string
			err = miRepo.DB.QueryRow(authorQuery, playlist.CreatorID).Scan(&creatorName)
			if err != nil {
				return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
			}
			playlist.CreatorName = creatorName
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (miRepo *MusicInfoPostgres) GetTracksFromAlbum(albumId int64) ([]e.Track, error) {
	const op = "repositories.musicinfo_psql.GetTracksFromAlbum"

	trackIdsQuery := `
        SELECT id 
        FROM tracks 
        WHERE album_id = $1
    `

	rows, err := miRepo.DB.Query(trackIdsQuery, albumId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get tracks from album: %w", op, err)
	}
	defer rows.Close()

	var trackIDs []int64
	for rows.Next() {
		var trackID int64
		if err := rows.Scan(&trackID); err != nil {
			return nil, fmt.Errorf("%s: failed to scan tracks id: %w", op, err)
		}
		trackIDs = append(trackIDs, trackID)
	}

	if len(trackIDs) == 0 {
		return []e.Track{}, nil
	}

	tracksQuery := `
        SELECT id, name, album_id, author_id, genre, length_sec, recording_date, size_in_bytes, number, format, path, cover_path
        FROM tracks 
        WHERE id = ANY($1::BIGINT[])
    `

	trackRows, err := miRepo.DB.Query(tracksQuery, pq.Array(trackIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get tracks from album info: %w", op, err)
	}
	defer trackRows.Close()

	var tracks []e.Track = make([]e.Track, 0, len(trackIDs))
	for trackRows.Next() {
		var track e.Track
		err := trackRows.Scan(
			&track.ID,
			&track.Name,
			&track.AlbumID,
			&track.AuthorID,
			&track.Genre,
			&track.Length,
			&track.RecordingDate,
			&track.SizeInBytes,
			&track.Number,
			&track.Format,
			&track.Path,
			&track.CoverPath,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan album: %w", op, err)
		}

		authorQuery := `
			SELECT login 
			FROM users 
			WHERE id = $1
    	`
		var artistName string
		err = miRepo.DB.QueryRow(authorQuery, track.AuthorID).Scan(&artistName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
		}
		track.AuthorName = artistName

		albumQuery := `
			SELECT name
			FROM albums 
			WHERE id = $1
    	`
		var albumName string
		err = miRepo.DB.QueryRow(albumQuery, track.AlbumID).Scan(&albumName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get album name: %w", op, err)
		}
		track.AlbumName = albumName

		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (miRepo *MusicInfoPostgres) GetAlbum(albumId int64) (*e.Album, error) {
	const op = "repositories.musicinfo_psql.GetAlbum"

	albumQuery := `
        SELECT id, name, author_id, genre, length_sec, recording_date, cover_path, type
        FROM albums 
        WHERE id = $1
    `

	albumRows, err := miRepo.DB.Query(albumQuery, albumId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get album info: %w", op, err)
	}
	defer albumRows.Close()

	var album e.Album
	albumRows.Next()
	err = albumRows.Scan(
		&album.ID,
		&album.Name,
		&album.AuthorID,
		&album.Genre,
		&album.Length,
		&album.RecordingDate,
		&album.CoverPath,
		&album.Type,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to scan album: %w", op, err)
	}

	authorQuery := `
		SELECT login 
		FROM users 
		WHERE id = $1
	`
	var authorName string
	err = miRepo.DB.QueryRow(authorQuery, album.AuthorID).Scan(&authorName)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
	}
	album.AuthorName = authorName
	return &album, nil
}

func (miRepo *MusicInfoPostgres) GetTracksFromPlaylist(playlistId int64) ([]e.Track, error) {
	const op = "repositories.musicinfo_psql.GetTracksFromPlaylist"

	trackIdsQuery := `
        SELECT track_id, added_at
        FROM tracks_playlists
        WHERE playlist_id = $1
    `

	rows, err := miRepo.DB.Query(trackIdsQuery, playlistId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get tracks from playlist: %w", op, err)
	}
	defer rows.Close()

	var trackIDs []int64
	addedAts := make(map[int64]time.Time)
	for rows.Next() {
		var trackID int64
		var addedAt time.Time
		if err := rows.Scan(&trackID, &addedAt); err != nil {
			return nil, fmt.Errorf("%s: failed to scan tracks id: %w", op, err)
		}
		trackIDs = append(trackIDs, trackID)
		addedAts[trackID] = addedAt
	}

	if len(trackIDs) == 0 {
		return []e.Track{}, nil
	}

	tracksQuery := `
        SELECT id, name, album_id, author_id, genre, length_sec, recording_date, size_in_bytes, number, format, path, cover_path
        FROM tracks 
        WHERE id = ANY($1::BIGINT[])
    `

	trackRows, err := miRepo.DB.Query(tracksQuery, pq.Array(trackIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get tracks from playlist info: %w", op, err)
	}
	defer trackRows.Close()

	var tracks []e.Track = make([]e.Track, 0, len(trackIDs))
	for trackRows.Next() {
		var track e.Track
		err := trackRows.Scan(
			&track.ID,
			&track.Name,
			&track.AlbumID,
			&track.AuthorID,
			&track.Genre,
			&track.Length,
			&track.RecordingDate,
			&track.SizeInBytes,
			&track.Number,
			&track.Format,
			&track.Path,
			&track.CoverPath,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan playlist: %w", op, err)
		}

		authorQuery := `
            SELECT login 
            FROM users 
            WHERE id = $1
        `
		var artistName string
		err = miRepo.DB.QueryRow(authorQuery, track.AuthorID).Scan(&artistName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
		}
		track.AuthorName = artistName

		albumQuery := `
            SELECT name
            FROM albums
            WHERE id = $1
        `
		var albumName string
		err = miRepo.DB.QueryRow(albumQuery, track.AlbumID).Scan(&albumName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get album name: %w", op, err)
		}
		track.AlbumName = albumName

		tracks = append(tracks, track)
	}

	sort.Slice(tracks, func(i, j int) bool {
		return addedAts[tracks[i].ID].After(addedAts[tracks[j].ID])
	})

	for i := range tracks {
		tracks[i].Number = int16(i + 1)
	}

	return tracks, nil
}

func (miRepo *MusicInfoPostgres) GetPlaylist(playlistId int64) (*e.Playlist, error) {
	const op = "repositories.musicinfo_psql.GetPlaylist"

	playlistQuery := `
        SELECT id, name, creator_id, length_sec, creation_time, cover_path, attached_to, description
        FROM playlists 
        WHERE id = $1
    `

	playlistRows, err := miRepo.DB.Query(playlistQuery, playlistId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get playlist info: %w", op, err)
	}
	defer playlistRows.Close()

	var playlist e.Playlist
	playlistRows.Next()
	err = playlistRows.Scan(
		&playlist.ID,
		&playlist.Name,
		&playlist.CreatorID,
		&playlist.Length,
		&playlist.CreationTime,
		&playlist.CoverPath,
		&playlist.AttachedTo,
		&playlist.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to scan playlist: %w", op, err)
	}

	creatorQuery := `
		SELECT login 
		FROM users 
		WHERE id = $1
	`
	var creatorName string
	err = miRepo.DB.QueryRow(creatorQuery, playlist.CreatorID).Scan(&creatorName)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
	}
	playlist.CreatorName = creatorName
	return &playlist, nil
}

func (miRepo *MusicInfoPostgres) GetPlaylistSaves(playlistId int64) (int64, error) {
	const op = "repositories.musicinfo_psql.GetPlaylistSaves"

	playlistQuery := `
        SELECT user_id
        FROM users_playlists 
        WHERE playlist_id = $1
    `

	playlistRows, err := miRepo.DB.Query(playlistQuery, playlistId)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get playlist info: %w", op, err)
	}
	defer playlistRows.Close()

	var count int64 = 0
	for playlistRows.Next() {
		var userId int64
		if err := playlistRows.Scan(&userId); err != nil {
			return 0, fmt.Errorf("%s: failed to scan user id: %w", op, err)
		}
		count++
	}

	if err := playlistRows.Err(); err != nil {
		return 0, fmt.Errorf("%s: error after scanning rows: %w", op, err)
	}

	return count, nil
}

func (miRepo *MusicInfoPostgres) GetReleasesFromArtist(artistId int64) ([]e.Album, error) {
	const op = "repositories.musicinfo_psql.GetReleasesFromArtist"

	albumIdsQuery := `
        SELECT id 
        FROM albums 
        WHERE author_id = $1
    `

	rows, err := miRepo.DB.Query(albumIdsQuery, artistId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get artist releases: %w", op, err)
	}
	defer rows.Close()

	var albumIDs []int64
	for rows.Next() {
		var albumID int64
		if err := rows.Scan(&albumID); err != nil {
			return nil, fmt.Errorf("%s: failed to scan release id: %w", op, err)
		}
		albumIDs = append(albumIDs, albumID)
	}

	if len(albumIDs) == 0 {
		return []e.Album{}, nil
	}

	albumsQuery := `
        SELECT id, name, author_id, genre, length_sec, recording_date, cover_path, type
        FROM albums 
        WHERE id = ANY($1::BIGINT[])
    `

	albumRows, err := miRepo.DB.Query(albumsQuery, pq.Array(albumIDs))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get releases info: %w", op, err)
	}
	defer albumRows.Close()

	var releases []e.Album = make([]e.Album, 0, len(albumIDs))
	for albumRows.Next() {
		var album e.Album
		err := albumRows.Scan(
			&album.ID,
			&album.Name,
			&album.AuthorID,
			&album.Genre,
			&album.Length,
			&album.RecordingDate,
			&album.CoverPath,
			&album.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan release: %w", op, err)
		}

		authorQuery := `
			SELECT login 
			FROM users 
			WHERE id = $1
    	`
		var authorName string
		err = miRepo.DB.QueryRow(authorQuery, album.AuthorID).Scan(&authorName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get author name: %w", op, err)
		}
		album.AuthorName = authorName
		releases = append(releases, album)
	}
	return releases, nil
}

func (miRepo *MusicInfoPostgres) GetArtist(artistId int64) (*e.User, error) {
	const op = "repositories.musicinfo_psql.GetArtist"

	artistQuery := `
        SELECT id, login, pic_path
        FROM users 
        WHERE id = $1
    `

	artistRows, err := miRepo.DB.Query(artistQuery, artistId)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get artist info: %w", op, err)
	}
	defer artistRows.Close()

	var artist e.User
	artistRows.Next()
	err = artistRows.Scan(
		&artist.ID,
		&artist.Login,
		&artist.PicPath,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to scan artist: %w", op, err)
	}

	return &artist, nil
}

func (miRepo *MusicInfoPostgres) GetArtistAttachmentId(artistId int64) (int64, error) {
	const op = "repositories.musicinfo_psql.GetArtistAttachmentId"

	playlistQuery := `
        SELECT id
        FROM playlists 
        WHERE attached_to = $1 AND NOT name = $2
    `

	playlistRows, err := miRepo.DB.Query(playlistQuery, artistId, "Liked Songs")
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get playlist info: %w", op, err)
	}
	defer playlistRows.Close()

	playlistRows.Next()
	var playlistId int64
	if err := playlistRows.Scan(&playlistId); err != nil {
		return 0, fmt.Errorf("%s: failed to scan user id: %w", op, err)
	}

	return playlistId, nil
}

func (miRepo *MusicInfoPostgres) GetlikedSongsId(userId int64) (int64, error) {
	const op = "repositories.musicinfo_psql.GetArtistAttachmentId"

	playlistQuery := `
        SELECT id
        FROM playlists 
        WHERE attached_to = $1 AND name = $2
    `

	playlistRows, err := miRepo.DB.Query(playlistQuery, userId, "Liked Songs")
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get playlist info: %w", op, err)
	}
	defer playlistRows.Close()

	playlistRows.Next()
	var playlistId int64
	if err := playlistRows.Scan(&playlistId); err != nil {
		return 0, fmt.Errorf("%s: failed to scan user id: %w", op, err)
	}

	return playlistId, nil
}

func (miRepo *MusicInfoPostgres) GetIsFollowedArtist(uId, cId int64) (bool, error) {
	const op = "repositories.musicinfo_psql.GetIsFollowedArtist"

	 query := `
        SELECT EXISTS(
            SELECT 1 
            FROM users_users
            WHERE first_user_id = $1 AND second_user_id = $2
        )`
    
    var isFollowed bool
    err := miRepo.DB.QueryRow(query, uId, cId).Scan(&isFollowed)
    if err != nil {
        return false, fmt.Errorf("%s: failed to check follow status: %w", op, err)
    }

    return isFollowed, nil
}

func (miRepo *MusicInfoPostgres) GetIsFollowedAlbum(uId, cId int64) (bool, error) {
	const op = "repositories.musicinfo_psql.GetIsFollowedAlbum"

	 query := `
        SELECT EXISTS(
            SELECT 1 
            FROM users_albums
            WHERE user_id = $1 AND album_id = $2
        )`
    
    var isFollowed bool
    err := miRepo.DB.QueryRow(query, uId, cId).Scan(&isFollowed)
    if err != nil {
        return false, fmt.Errorf("%s: failed to check follow status: %w", op, err)
    }

    return isFollowed, nil
}
func (miRepo *MusicInfoPostgres) GetIsFollowedPlaylist(uId, cId int64) (bool, error) {
	const op = "repositories.musicinfo_psql.GetIsFollowedPlaylist"

	 query := `
        SELECT EXISTS(
            SELECT 1 
            FROM users_playlists
            WHERE user_id = $1 AND playlist_id = $2
        )`
    
    var isFollowed bool
    err := miRepo.DB.QueryRow(query, uId, cId).Scan(&isFollowed)
    if err != nil {
        return false, fmt.Errorf("%s: failed to check follow status: %w", op, err)
    }

    return isFollowed, nil
}
