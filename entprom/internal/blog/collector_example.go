package blog

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var labelNames = []string{"label_1", "label_2"}

var exampleCollector = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name:        "example_metric_total",
		Help:        "Number of example metrics used",
		ConstLabels: map[string]string{"key": "value"},
	},
	labelNames,
)

func useCollector(){
	exampleCollector.WithLabelValues("value1","value2").Inc()
}