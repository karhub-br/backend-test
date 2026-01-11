package spotify

import (
	"context"
	"encoding/json"
	"io"
	"karhub/internal/entity"
	"net/http"
)

type search struct {
	url   string
	token string
}

func (s *search) Playlist(ctx context.Context, style string) (searchResp entity.SpotifySearchResponse, err error) {
	var query = "/search?q=" + style + "&type=playlist&limit=1"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.url+query, nil)
	if err != nil {
		return entity.SpotifySearchResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+s.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.SpotifySearchResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.SpotifySearchResponse{}, err
	}

	err = json.Unmarshal(body, &searchResp)

	return
}

func NewSearch(url string, token string) *search {
	return &search{url: url, token: token}
}
