package response

type AllSubsResponse struct {
	Playlists []interface{} `json:"playlists,omitempty"`
	Artists   []interface{} `json:"artists,omitempty"`
	Albums    []interface{} `json:"albums,omitempty"`
}

type ArtistResponse struct {
	Artist     ArtistModel        `json:"artist"`
	Releases   []AlbumInListModel `json:"releases"`
	Attachment PlaylistViewModel  `json:"attached"`
	Tracks     []TrackInSubsModel `json:"tracks"`
}

type PlaylistResponse struct {
	Playlist PlaylistViewModel  `json:"playlist"`
	Tracks   []TrackInSubsModel `json:"tracks"`
}
type AlbumResponse struct {
	Album  AlbumViewModel     `json:"album"`
	Tracks []TrackInSubsModel `json:"tracks"`
}