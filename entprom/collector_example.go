package entprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var labelNames []string = []string{"label_1", "label_2"}

var exampleCollector = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name:        "ent_operation_total",
		Help:        "Number of ent mutation operations",
		ConstLabels: map[string]string{"key": "value"},
	},
	entLabels,
)
