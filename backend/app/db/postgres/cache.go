package postgres

import (
	"backend/app/api"
	"backend/app/models"
	"sync"
)

var (
	CacheDevelopersSet sync.Map
	CacheDevelopers    map[string]*api.DeveloperApi
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
	var err error
	for _, dev := range devData {
		detail := &api.DeveloperApi{
			Id:            dev.GithubId,
			Login:         dev.Login,
			Type:          dev.Type,
			Name:          dev.Name,
			Company:       dev.Company,
			Blog:          dev.Blog,
			Location:      dev.Location,
			Email:         dev.Email,
			CreatedAt:     dev.CreatedAt,
			Languages:     nil,
			Contributions: nil,
			TalentRank:    dev.TalentRank,
		}
		detail.Contributions, err = GetContributionsByDeveloper(dev.Login)
		if err != nil {
			return err
		}

		detail.Languages, err = GetLanguages("users", dev.GithubId)
		if err == nil {
			return err
		}
	}
	return nil
}

func CacheInsertRepo(*models.Repos) {

}

func CacheInsertDeveloper(dev *models.Developer) {

}
