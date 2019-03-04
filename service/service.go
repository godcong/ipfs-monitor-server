package service

import "github.com/godcong/ipfs-monitor-server/config"

// Service ...
type Service struct {
	config *config.Configure
	rest   *RestServer
}

var globalService *Service

// New ...
func New(cfg *config.Configure) *Service {
	return &Service{
		config: cfg,
		rest:   NewRestServer(cfg),
	}
}

// Start ...
func Start(cfg *config.Configure) {
	globalService = New(cfg)
	globalService.rest.Start()
}

// Stop ...
func Stop() {
	globalService.rest.Stop()
}
