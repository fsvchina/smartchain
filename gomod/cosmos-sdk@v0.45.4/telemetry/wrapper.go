package telemetry

import (
	"time"

	metrics "github.com/armon/go-metrics"
)


const (
	MetricKeyBeginBlocker = "begin_blocker"
	MetricKeyEndBlocker   = "end_blocker"
	MetricLabelNameModule = "module"
)


func NewLabel(name, value string) metrics.Label {
	return metrics.Label{Name: name, Value: value}
}




func ModuleMeasureSince(module string, start time.Time, keys ...string) {
	metrics.MeasureSinceWithLabels(
		keys,
		start.UTC(),
		append([]metrics.Label{NewLabel(MetricLabelNameModule, module)}, globalLabels...),
	)
}




func ModuleSetGauge(module string, val float32, keys ...string) {
	metrics.SetGaugeWithLabels(
		keys,
		val,
		append([]metrics.Label{NewLabel(MetricLabelNameModule, module)}, globalLabels...),
	)
}



func IncrCounter(val float32, keys ...string) {
	metrics.IncrCounterWithLabels(keys, val, globalLabels)
}



func IncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.IncrCounterWithLabels(keys, val, append(labels, globalLabels...))
}



func SetGauge(val float32, keys ...string) {
	metrics.SetGaugeWithLabels(keys, val, globalLabels)
}



func SetGaugeWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.SetGaugeWithLabels(keys, val, append(labels, globalLabels...))
}



func MeasureSince(start time.Time, keys ...string) {
	metrics.MeasureSinceWithLabels(keys, start.UTC(), globalLabels)
}
