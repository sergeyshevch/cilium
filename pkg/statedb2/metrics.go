// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package statedb2

import (
	"github.com/cilium/cilium/pkg/metrics"
	"github.com/cilium/cilium/pkg/metrics/metric"
)

type Metrics struct {
	// How long a read transaction was held.
	WriteTxnDuration metric.Vec[metric.Observer]
	// How long it took to acquire a write transaction for all tables.
	WriteTxnAcquisition metric.Vec[metric.Observer]
	// How long writers were blocked while waiting to acquire a write transaction for a specific table.
	TableContention metric.Vec[metric.Gauge]
	// The amount of objects in a given table.
	TableObjectCount metric.Vec[metric.Gauge]
	// The current revision of a given table.
	TableRevision metric.Vec[metric.Gauge]
	// The amount of delete trackers for a given table.
	TableDeleteTrackerCount metric.Vec[metric.Gauge]
	// The amount of objects in the graveyard for a given table.
	TableGraveyardObjects metric.Vec[metric.Gauge]
	// The lowest revision of a given table that has been processed by the graveyard garbage collector.
	TableGraveyardLowWatermark metric.Vec[metric.Gauge]
	// The time it took to clean the graveyard for a given table.
	TableGraveyardCleaningDuration metric.Vec[metric.Observer]
}

func NewMetrics() Metrics {
	return Metrics{
		WriteTxnDuration: metric.NewHistogramVec(metric.HistogramOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "write_txn_duration_seconds",
			Help:      "How long a write transaction was held.",
		}, []string{"tables", "package"}),
		WriteTxnAcquisition: metric.NewHistogramVec(metric.HistogramOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "write_txn_acquisition_seconds",
			Help:      "How long it took to acquire a write transaction for all tables.",
		}, []string{"tables", "package"}),
		TableContention: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_contention_seconds",
			Help:      "How long writers were blocked while waiting to acquire a write transaction for a specific table.",
		}, []string{"table"}),
		TableObjectCount: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_objects_total",
			Help:      "The amount of objects in a given table.",
		}, []string{"table"}),
		TableRevision: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_revision",
			Help:      "The current revision of a given table.",
		}, []string{"table"}),
		TableDeleteTrackerCount: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_delete_trackers_total",
			Help:      "The amount of delete trackers for a given table.",
		}, []string{"table"}),
		TableGraveyardObjects: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_graveyard_objects_total",
			Help:      "The amount of objects in the graveyard for a given table.",
		}, []string{"table"}),
		TableGraveyardLowWatermark: metric.NewGaugeVec(metric.GaugeOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_graveyard_low_watermark",
			Help:      "The lowest revision of a given table that has been processed by the graveyard garbage collector.",
		}, []string{"table"}),
		TableGraveyardCleaningDuration: metric.NewHistogramVec(metric.HistogramOpts{
			Namespace: metrics.CiliumAgentNamespace,
			Subsystem: "statedb",
			Name:      "table_graveyard_cleaning_duration_seconds",
			Help:      "The time it took to clean the graveyard for a given table.",
		}, []string{"table"}),
	}
}
