package iw

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Station holds information about the current wifi station
type Station struct {
	iface string
	mac   string
}

func (s *Station) line(line string, ch chan<- prometheus.Metric) {
	// "  signal:    -78 dBm"
	kv := strings.SplitN(line, ":", 2)
	k := strings.TrimSpace(kv[0])
	metric, ok := StationMetrics[k]
	if !ok {
		return // unsupported field
	}

	f := strings.Fields(kv[1])
	if len(f) < 1 {
		return // no value field
	}

	v := f[0]
	if v == "yes" {
		v = "1"
	} else if v == "no" {
		v = "0"
	}
	// TODO: look for "ms", "seconds", "MBit", and do conversions.

	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("%q: %q: %v\n", line, v, err)
		return
	}
	ch <- prometheus.MustNewConstMetric(
		metric.Desc, metric.ValueType, value, s.iface, s.mac)
}

// ReadStation reads "station dump" text and outputs prometheus metrics
func ReadStation(iface string, r io.Reader, ch chan<- prometheus.Metric) error {
	var station *Station

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		if strings.HasPrefix(t, "Station ") {
			// "Station f8:f1:b6:00:00:04 (on wlan0)"
			f := strings.Fields(t)
			if len(f) >= 2 {
				station = &Station{iface: iface, mac: f[1]}
			} else {
				station = &Station{iface: iface, mac: "unknown"}
			}
		} else if station != nil {
			station.line(t, ch)
		}
	}

	return s.Err()
}
