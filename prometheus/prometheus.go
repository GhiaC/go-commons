package prometheus

import (
	"fmt"
	"github.com/ghiac/go-commons/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Prometheus struct{}

func NewPrometheus(port int) *Prometheus {
	go runPrometheusHttpServer(port)
	return &Prometheus{}
}

func runPrometheusHttpServer(port int) {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Logger.Error("Failed to listen prometheus server: ", err)
	}
}
