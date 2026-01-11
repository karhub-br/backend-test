package main

import (
	handlers "karhub/internal/handler"
	beerStyleRepo "karhub/internal/repository/beer_style"
	"karhub/internal/server"
	"karhub/internal/spotify"
	beerstyle "karhub/internal/usecase/beer_style"
	"karhub/pkg/config"
	"karhub/pkg/postgres"
	"karhub/pkg/rabbit"
	"karhub/pkg/telemetry"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	telemetry.InitMetrics()

	http.Handle("/metrics", promhttp.Handler())

	go http.ListenAndServe(":9091", nil)

	server.NewServer(injectDependency()).Start()

}

func injectDependency() (server.Handler, string) {
	cfg := config.LoadConfig()
	db := postgres.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	_, channel := rabbit.Connect(cfg.RabbitHost, cfg.RabbitUser, cfg.RabbitPass, cfg.RabbitPort)

	querier := postgres.NewQuerier(db)
	repoSql := beerStyleRepo.NewRepo(querier)
	searchSpotify := spotify.NewSearch(cfg.SpotifyURL, cfg.SpotifyToken)
	playlist := spotify.NewPlaylist(cfg.SpotifyURL, cfg.SpotifyToken)
	rb := rabbit.NewRabbit(channel)
	usecase := beerstyle.NewBeerStyle(repoSql, searchSpotify, playlist, rb)
	usecaseWithDecorator := beerstyle.NewBeerStyleDecorator(usecase)
	return handlers.NewBeersHandler(usecaseWithDecorator), cfg.ServerPort

}
