package managers

import (
	"HarborMaster/models"
	"errors"
	"sort"

	"github.com/mheidinger/docker-registry-client/registry"
	log "github.com/sirupsen/logrus"
)

const LatestTag = "latest"

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
		tags, err := mgr.GetTagsForRepo(repo)
		if err != nil {
			continue
		}

		info.RepositoryInfos[it] = &models.RepositoryInfo{
			Name: repo,
			Tags: tags,
		}
	}

	return
}

type TagSort []string

func (a TagSort) Len() int {
	return len(a)
}

func (a TagSort) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		return a[i] < a[j]
	}
	return len(a[i]) < len(a[j])
}

func (a TagSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (mgr *RegistryManager) GetTagsForRepo(repository string) (tags []string, err error) {
	tags, err = mgr.hub.Tags(repository)
	if err != nil {
		log.WithField("repository", repository).WithError(err).Error("Failed to get tags for repo")
		return
	}

	sort.Sort(TagSort(tags))

	latestIt := -1
	for it, tag := range tags {
		if tag == LatestTag {
			latestIt = it
			break
		}
	}
	if latestIt != -1 {
		newTags := []string{LatestTag}
		newTags = append(newTags, tags[:latestIt]...)
		if latestIt != len(tags)-1 {
			newTags = append(newTags, tags[latestIt+1:]...)
		}
		tags = newTags
	}

	return
}

func (mgr *RegistryManager) DeleteTag(repository, tag string) (err error) {
	deleteLog := log.WithFields(log.Fields{"repository": repository, "tag": tag})

	if tag == LatestTag {
		deleteLog.Info("Can't delete 'latest' tag")
		return errors.New("Can't delete 'latest' tag")
	}

	digest, err := mgr.hub.ManifestDigestV2(repository, tag)
	if err != nil {
		deleteLog.WithError(err).Error("Failed to get digest to delete tag")
		return
	}

	err = mgr.hub.DeleteManifest(repository, digest)
	if err != nil {
		deleteLog.WithError(err).Error("Failed to delete tag with digest")
		return
	}
	return
}
