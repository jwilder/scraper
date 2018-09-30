package main

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	sd_config "github.com/prometheus/prometheus/discovery/config"
	"github.com/prometheus/prometheus/discovery/kubernetes"
	"github.com/prometheus/prometheus/pkg/labels"

	"os"

	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
)

func main() {
	discoveryManager := discovery.NewManager(context.Background(), log.NewLogfmtLogger(os.Stdout))
	go discoveryManager.Run()

	// These scrape endpoints are hard-coded, but should be discovered as well
	cfg := &config.Config{
		ScrapeConfigs: []*config.ScrapeConfig{{
			JobName:     "k8s",
			HonorLabels: true,
			MetricsPath: "/metrics",
			Scheme:      "https",
		}},
	}

	c := make(map[string]sd_config.ServiceDiscoveryConfig)
	sdConfig := sd_config.ServiceDiscoveryConfig{}

	// Discover everything...
	for _, v := range []string{"node", "endpoints", "service", "pod", "ingress"} {
		k8sCfg := &kubernetes.SDConfig{
			Role: kubernetes.Role(v),
		}
		sdConfig.KubernetesSDConfigs = append(sdConfig.KubernetesSDConfigs, k8sCfg)
	}
	c["k8s"] = sdConfig

	discoveryManager.ApplyConfig(c)

	// Set some default scrape intervals and timeouts
	cfg.ScrapeConfigs[0].ScrapeInterval.Set("10s")
	cfg.ScrapeConfigs[0].ScrapeTimeout.Set("10s")

	// Configure our scrape manager with our storage
	scrapeManager := scrape.NewManager(log.NewLogfmtLogger(os.Stdout), &store{})

	scrapeManager.ApplyConfig(cfg)

	// Start everything up...
	scrapeManager.Run(discoveryManager.SyncCh())

}

// This is a storage.Appender that accepts scraped metrics data.
type store struct{}

func (s *store) Add(l labels.Labels, t int64, v float64) (uint64, error) {
	fmt.Printf("store.Add %v %d %v\n", l.String(), t, v)
	return 1, nil
}

func (s *store) AddFast(l labels.Labels, ref uint64, t int64, v float64) error {
	fmt.Printf("store.AddFast %v %d %d %v\n", l.String(), ref, t, v)
	return nil
}

func (s *store) Commit() error {
	fmt.Println("store.Commit")
	return nil
}

func (s *store) Rollback() error {
	fmt.Println("store.Rollback")
	return nil
}

func (s *store) Appender() (storage.Appender, error) {
	return s, nil
}
