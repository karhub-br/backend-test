package spotify

import (
	"context"
	"encoding/json"
	"io"
	"karhub/internal/entity"
	"net/http"
)

type playlist struct {
	url   string
	token string
}

func (p *playlist) Get(ctx context.Context, id string) (playlist entity.SpotifyPlaylistResponse, err error) {
	endpoint := "/playlists/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.url+endpoint, nil)
	if err != nil {
		return entity.SpotifyPlaylistResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+p.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.SpotifyPlaylistResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.SpotifyPlaylistResponse{}, err
	}

	err = json.Unmarshal(body, &playlist)

	return
}

func NewPlaylist(url string, token string) *playlist {
	return &playlist{url: url, token: token}
}
