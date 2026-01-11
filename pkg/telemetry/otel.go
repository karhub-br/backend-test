package telemetry

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var meter = otel.Meter("karhub-metrics")

func InitMetrics() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	otel.SetMeterProvider(provider)
}

func TrackDuration(ctx context.Context, action string, start time.Time) {
	elapsed := time.Since(start).Seconds()

	histogram, _ := meter.Float64Histogram(
		"usecase_duration_seconds",
		metric.WithDescription("Duração da execução do usecase"),
	)

	histogram.Record(ctx, elapsed, metric.WithAttributes(
		attribute.String("action", action),
	))
}
