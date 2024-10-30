package postgres

import (
	"backend/app/models"
	"sync"
)

var (
	CacheDevelopersSet sync.Map
	CacheDevelopers    map[string]*models.DeveloperStored
	CacheReposSet      sync.Map
	CacheRepos         map[string]*models.RepoStored
)

const (
	DataNeverCalled = 0 + iota
	DataProcessing
	DataStored
)

func CacheInit() error {
	devData := make([]models.DeveloperStored, 0)
	res := pdb.Find(&devData)
	if res.Error != nil {
		return res.Error
	}
	for _, dev := range devData {
		CacheDevelopersSet.Store(dev.Login, DataStored)
		CacheDevelopers[dev.Login] = &dev
	}
	repoData := make([]models.RepoStored, 0)
	res = pdb.Find(&repoData)
	if res.Error != nil {
		return res.Error
	}
	for _, repo := range repoData {
		CacheReposSet.Store(repo.FullName, DataStored)
		CacheRepos[repo.FullName] = &repo
	}
	return nil
}
