package blog

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// List of dynamic labels
var labelNames = []string{"label_1", "label_2"}

// Create a counter collector
var exampleCollector = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name:        "example_metric_total",
		Help:        "Number of example metrics used",
		ConstLabels: map[string]string{"key": "value"},
	},
	labelNames,
)

// When using you set the values of the dynamic labels and then increment the counter
func useCollector(){
	exampleCollector.WithLabelValues("value1","value2").Inc()
}