package manager

import (
	"github.com/heroku/docker-registry-client/registry"
	log "github.com/sirupsen/logrus"
)

type RegistryManager struct {
	hub  *registry.Registry
	done chan struct{}
}

var registryManager *RegistryManager

func CreateRegistryManager(url, username, password string) (*RegistryManager, error) {
	if registryManager != nil {
		return registryManager, nil
	}

	createLog := log.WithFields(log.Fields{"url": url, "username": username})
	createLog.Info("Connect with registry")

	hub, err := registry.New(url, username, password)
	if err != nil {
		createLog.WithError(err).Error("Failed to connect to registry")
		return nil, err
	}

	registryManager = &RegistryManager{
		hub:  hub,
		done: make(chan struct{}),
	}
	return registryManager, nil
}

func GetRegistryManager() *RegistryManager {
	return registryManager
}

func (mgr *RegistryManager) GetRepositories() ([]string, error) {
	return mgr.hub.Repositories()
}
