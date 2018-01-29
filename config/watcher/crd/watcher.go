package crd

import (
	"fmt"
	"time"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/solo-io/glue/config/watcher"
)

func NewCrdWatcher(masterUrl, kubeconfigPath string, resyncDuration time.Duration) (watcher.Watcher, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build rest config: %v", err)
	}
	err = registerCrds(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to register crds: %v", err)
	}

	ctl, err := newCrdController(cfg, resyncDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize crd controller: %v", err)
	}

	return ctl, nil
}
