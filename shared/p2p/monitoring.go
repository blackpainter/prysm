package p2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-host"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	peerCountMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "p2p_peer_count",
		Help: "The number of currently connected peers",
	})
)

func init() {
	prometheus.MustRegister(peerCountMetric)
}

func startPeerWatcher(ctx context.Context, h host.Host) {

	go (func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Update peer count (minus ourselves)
				peerCountMetric.Set(float64(len(h.Peerstore().Peers()) - 1))

				// Wait 1 second to update again
				time.Sleep(1 * time.Second)
			}
		}
	})()

}
