package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port uint16 = 8080

	requests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_requests_total",
		Help: "Counter of received requests",
	})
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	envPort, ok := os.LookupEnv("PORT")
	if !ok {
		slog.Warn("$PORT not set")
		return
	}

	newPort, err := strconv.ParseUint(envPort, 10, 16)
	if err != nil {
		slog.Error("failed to parse $PORT", "PORT", envPort, "err", err)
		return
	}

	port = uint16(newPort)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests.Inc()
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte("Hello World! I'm a teapot.\n"))
	})
	mux.Handle("/metrics", promhttp.Handler())

	slog.Info("starting server", "port", port)
	slog.Error("server exited", "error", http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
