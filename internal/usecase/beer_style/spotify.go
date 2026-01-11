package beerstyle

import (
	"context"
	"karhub/internal/entity"
)

type Playlist interface {
	Get(ctx context.Context, id string) (playlist entity.SpotifyPlaylistResponse, err error)
}

type Search interface {
	Playlist(ctx context.Context, style string) (searchResp entity.SpotifySearchResponse, err error)
}
