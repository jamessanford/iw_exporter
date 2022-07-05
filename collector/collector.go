package collector

import (
	"log"

	"github.com/jamessanford/iw_exporter/iw"
	"github.com/prometheus/client_golang/prometheus"
)

// Instrumentation for the collector itself
var (
	iwCollectionsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "iw_collector_requests_total",
			Help: "number of requests to collect wireless data",
		})
	iwErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "iw_collector_errors_total",
			Help: "number of errors while collecting wireless data",
		})
)

func init() {
	prometheus.MustRegister(iwCollectionsTotal)
	prometheus.MustRegister(iwErrorsTotal)
}

// iwCollector implements the prometheus.Collector interface
type iwCollector struct {
}

// NewIWCollector returns a prometheus.Collector that exports wifi station data.
func NewIWCollector() prometheus.Collector {
	return &iwCollector{}
}

func (i *iwCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, l := range iw.StationMetrics {
		ch <- l.Desc
	}
}

func (i *iwCollector) Collect(ch chan<- prometheus.Metric) {
	iwCollectionsTotal.Inc()

	// TODO: export device metadata (channel, txpower)
	devs, err := iw.DeviceCmd()
	if err != nil {
		iwErrorsTotal.Inc()
		log.Println(err)
		return
	}

	// TODO: export total station count across all devices
	for _, d := range devs {
		if err = iw.StationDumpCmd(d.Iface, ch); err != nil {
			iwErrorsTotal.Inc()
			log.Println(err)
			// keep going
		}
	}
}
