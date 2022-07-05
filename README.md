**Connected wireless station exporter for Prometheus.**

Intended to run directly on wireless access points (for example WRT1200AC running OpenWRT)

This shells out to `iw dev` and `iw dev <dev> station dump`.  It does not use netlink directly.

#### Install, run, view

    go install github.com/jamessanford/iw_exporter@latest
    iw_exporter
    curl http://localhost:6798/metrics

#### Metrics exported

```
iw_collector_errors_total 0
iw_collector_requests_total 13737
iw_connected_seconds{interface="wlan0",station="ac:22:ff:ff:ff:03"} 554
iw_inactive_time_ms{interface="wlan0",station="ac:22:ff:ff:ff:03"} 3870
iw_is_associated{interface="wlan0",station="ac:22:ff:ff:ff:03"} 1
iw_is_authenticated{interface="wlan0",station="ac:22:ff:ff:ff:03"} 1
iw_is_authorized{interface="wlan0",station="ac:22:ff:ff:ff:03"} 1
iw_receive_bitrate_mbps{interface="wlan0",station="ac:22:ff:ff:ff:03"} 48
iw_receive_bytes_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 156412
iw_receive_packets_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 789
iw_signal_strength_avg_dbm{interface="wlan0",station="ac:22:ff:ff:ff:03"} -74
iw_signal_strength_dbm{interface="wlan0",station="ac:22:ff:ff:ff:03"} -74
iw_transmit_bitrate_mbps{interface="wlan0",station="ac:22:ff:ff:ff:03"} 135
iw_transmit_bytes_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 660556
iw_transmit_failures_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 0
iw_transmit_packets_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 727
iw_transmit_retries_total{interface="wlan0",station="ac:22:ff:ff:ff:03"} 0
```
