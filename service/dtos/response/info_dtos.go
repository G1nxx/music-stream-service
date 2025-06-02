package response

type AllSubsResponse struct {
	Playlists []interface{} `json:"playlists,omitempty"`
	Artists   []interface{} `json:"artists,omitempty"`
	Albums    []interface{} `json:"albums,omitempty"`
}