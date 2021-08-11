package entprom

import (
	"context"
	"time"

	"entgo.io/ent"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//Ent dynamic dimensions
const (
	mutationType = "mutation_type"
	mutationOp   = "mutation_op"
)

var entLabels = []string{mutationType, mutationOp}

// Create collector for total operations counter
func initOpsProcessedTotal(constLabels prometheus.Labels) *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "ent_operation_total",
			Help:        "Number of ent mutation operations",
			ConstLabels: constLabels,
		},
		entLabels,
	)
}

// Create collector for error counter
func initOpsProcessedError(constLabels prometheus.Labels) *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "ent_operation_error",
			Help:        "Number of failed ent mutation operations",
			ConstLabels: constLabels,
		},
		entLabels,
	)
}

// Create collector for duration histogram collector
func initOpsDuration(constLabels prometheus.Labels) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "ent_operation_duration_seconds",
			Help:        "Time in seconds per operation",
			ConstLabels: constLabels,
		},
		entLabels,
	)
}

// Hook init collectors, count total at beginning error on mutation error and duration also after.
func Hook(constLabels prometheus.Labels) ent.Hook {
	opsProcessedTotal := initOpsProcessedTotal(constLabels)
	opsProcessedError := initOpsProcessedError(constLabels)
	opsDuration := initOpsDuration(constLabels)
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			labels := prometheus.Labels{mutationType: m.Type(), mutationOp: m.Op().String()}
			opsProcessedTotal.With(labels).Inc()
			v, err := next.Mutate(ctx, m)
			if err != nil {
				opsProcessedError.With(labels).Inc()
			}
			duration := time.Since(start)
			opsDuration.With(labels).Observe(duration.Seconds())
			return v, err
		})
	}
}
