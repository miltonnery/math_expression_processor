package main

import (
	"flag"
	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"go-kit-base/infrastructure/configuration/environment"
	"go-kit-base/service/compute"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {


	//Loading environment
	env, err := environment.New()
	var (
		//TODO Set port from environment
		port     = env.GetString("server.port")
		httpAddr = flag.String("http.addr", ":"+port, "HTTP Listen Address")
	)
	if err != nil {
		logrus.Errorf("Error during properties loading!", err)
	}

	//Setting logs
	logger := kitLog.NewLogfmtLogger(os.Stderr)
	//validator := validation.New()
	var ts compute.Service
	ts = compute.NewService()
	ts = compute.NewLoggingMiddleware(logger, ts)
	ts = compute.NewInstrumentingMiddleware(configureInstrumentingMWForExpressionProcessing(ts))
	//ts = compute.NewValidationMiddleware(validator, ts)

	mux := http.NewServeMux()
	mux.Handle("/evaluate", compute.MakeHandler(ts))
	http.Handle("/", accessControl(mux))

	srv := http.Server{
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
		Addr:         *httpAddr,
	}

	log.Printf("Server is up and running on port: %s", port)
	_ = srv.ListenAndServe()
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func configureInstrumentingMWForExpressionProcessing(ts compute.Service) (
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
	service compute.Service) {
	service = ts
	fieldKeys := []string{"Method", "Input_expression"}
	requestCount = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "compute",
		Subsystem: "compute_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "compute",
		Subsystem: "compute_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "compute",
		Subsystem: "compute_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here
	return
}
