package managers

import (
	"HarborMaster/models"

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

func (mgr *RegistryManager) GetRegistryInfo() (info *models.RegistryInfo, err error) {
	repos, err := mgr.hub.Repositories()
	if err != nil {
		log.WithError(err).Error("Failed to get repositories")
		return
	}

	info = &models.RegistryInfo{
		URL:             mgr.hub.URL,
		RepositoryInfos: make([]*models.RepositoryInfo, len(repos)),
	}

	for it, repo := range repos {
		tags, err := mgr.hub.Tags(repo)
		if err != nil {
			log.WithField("repository", repo).WithError(err).Error("Failed to get tags for repo")
			continue
		}

		info.RepositoryInfos[it] = &models.RepositoryInfo{
			Name: repo,
			Tags: tags,
		}
	}

	return
}

func (mgr *RegistryManager) getRepositoryTags(repository string) (tags []string, err error) {

	return
}
