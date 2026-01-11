package main

import (
	"karhub/internal/consumer"
	beerstyle "karhub/internal/repository/beer_style"
	"karhub/internal/usecase/reprocess"
	"karhub/pkg/config"
	"karhub/pkg/postgres"
	"karhub/pkg/rabbit"
	"karhub/pkg/telemetry"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	telemetry.InitMetrics()

	http.Handle("/metrics", promhttp.Handler())

	go http.ListenAndServe(":9091", nil)
	
	uc, rb := injectDependency()

	log.Fatal(consumer.NewReprocess(rb, uc).Consume())

}

func injectDependency() (consumer.Reprocess, consumer.Rabbit) {
	cfg := config.LoadConfig()
	db := postgres.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	_, channel := rabbit.Connect(cfg.RabbitHost, cfg.RabbitUser, cfg.RabbitPass, cfg.RabbitPort)

	querier := postgres.NewQuerier(db)
	repoSql := beerstyle.NewRepo(querier)
	usecase := reprocess.NewReprocess(repoSql)
	usecaseWithDecorator := reprocess.NewReprocessDecorator(usecase)
	rb := rabbit.NewRabbit(channel)
	return usecaseWithDecorator, rb

}
