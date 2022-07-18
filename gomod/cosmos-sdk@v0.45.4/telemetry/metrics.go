package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	metrics "github.com/armon/go-metrics"
	metricsprom "github.com/armon/go-metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)



var globalLabels = []metrics.Label{}


const (
	FormatDefault    = ""
	FormatPrometheus = "prometheus"
	FormatText       = "text"
)


type Config struct {

	ServiceName string `mapstructure:"service-name"`




	Enabled bool `mapstructure:"enabled"`


	EnableHostname bool `mapstructure:"enable-hostname"`


	EnableHostnameLabel bool `mapstructure:"enable-hostname-label"`


	EnableServiceLabel bool `mapstructure:"enable-service-label"`



	PrometheusRetentionTime int64 `mapstructure:"prometheus-retention-time"`



	//


	GlobalLabels [][]string `mapstructure:"global-labels"`
}






type Metrics struct {
	memSink           *metrics.InmemSink
	prometheusEnabled bool
}


type GatherResponse struct {
	Metrics     []byte
	ContentType string
}


func New(cfg Config) (*Metrics, error) {
	if !cfg.Enabled {
		return nil, nil
	}

	if numGlobalLables := len(cfg.GlobalLabels); numGlobalLables > 0 {
		parsedGlobalLabels := make([]metrics.Label, numGlobalLables)
		for i, gl := range cfg.GlobalLabels {
			parsedGlobalLabels[i] = NewLabel(gl[0], gl[1])
		}

		globalLabels = parsedGlobalLabels
	}

	metricsConf := metrics.DefaultConfig(cfg.ServiceName)
	metricsConf.EnableHostname = cfg.EnableHostname
	metricsConf.EnableHostnameLabel = cfg.EnableHostnameLabel

	memSink := metrics.NewInmemSink(10*time.Second, time.Minute)
	metrics.DefaultInmemSignal(memSink)

	m := &Metrics{memSink: memSink}
	fanout := metrics.FanoutSink{memSink}

	if cfg.PrometheusRetentionTime > 0 {
		m.prometheusEnabled = true
		prometheusOpts := metricsprom.PrometheusOpts{
			Expiration: time.Duration(cfg.PrometheusRetentionTime) * time.Second,
		}

		promSink, err := metricsprom.NewPrometheusSinkFrom(prometheusOpts)
		if err != nil {
			return nil, err
		}

		fanout = append(fanout, promSink)
	}

	if _, err := metrics.NewGlobal(metricsConf, fanout); err != nil {
		return nil, err
	}

	return m, nil
}




func (m *Metrics) Gather(format string) (GatherResponse, error) {
	switch format {
	case FormatPrometheus:
		return m.gatherPrometheus()

	case FormatText:
		return m.gatherGeneric()

	case FormatDefault:
		return m.gatherGeneric()

	default:
		return GatherResponse{}, fmt.Errorf("unsupported metrics format: %s", format)
	}
}

func (m *Metrics) gatherPrometheus() (GatherResponse, error) {
	if !m.prometheusEnabled {
		return GatherResponse{}, fmt.Errorf("prometheus metrics are not enabled")
	}

	metricsFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return GatherResponse{}, fmt.Errorf("failed to gather prometheus metrics: %w", err)
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	e := expfmt.NewEncoder(buf, expfmt.FmtText)
	for _, mf := range metricsFamilies {
		if err := e.Encode(mf); err != nil {
			return GatherResponse{}, fmt.Errorf("failed to encode prometheus metrics: %w", err)
		}
	}

	return GatherResponse{ContentType: string(expfmt.FmtText), Metrics: buf.Bytes()}, nil
}

func (m *Metrics) gatherGeneric() (GatherResponse, error) {
	summary, err := m.memSink.DisplayMetrics(nil, nil)
	if err != nil {
		return GatherResponse{}, fmt.Errorf("failed to gather in-memory metrics: %w", err)
	}

	content, err := json.Marshal(summary)
	if err != nil {
		return GatherResponse{}, fmt.Errorf("failed to encode in-memory metrics: %w", err)
	}

	return GatherResponse{ContentType: "application/json", Metrics: content}, nil
}
