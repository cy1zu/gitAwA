package postgres

import (
	"backend/app/models"
	"errors"
	"sync"
)

var (
	CacheDevelopersSet sync.Map
	CacheDevelopers    map[string]models.DeveloperApi
	CacheReposSet      sync.Map
	CacheRepos         map[string]models.RepoStored
)

const (
	DataNeverCalled = 0 + iota
	DataProcessing
	DataStored
	DataUpdating
)

func CacheInit() error {
	CacheDevelopers = make(map[string]models.DeveloperApi)
	CacheRepos = make(map[string]models.RepoStored)

	var devData []models.DeveloperStored
	res := pdb.Find(&devData)
	if res.Error != nil {
		return res.Error
	}
	var err error
	for _, dev := range devData {
		detail := models.DeveloperApi{
			Login:         dev.Login,
			Type:          dev.Type,
			Location:      dev.Location,
			Languages:     nil,
			TopLanguages:  dev.TopLanguages,
			Contributions: nil,
			TalentRank:    dev.TalentRank,
			Nation:        dev.Nation,
		}
		detail.Contributions, err = GetContributionsByDeveloper(dev.Login)
		if err != nil {
			return err
		}
		detail.Languages, err = GetLanguages("users", dev.GithubId)
		if err != nil {
			return err
		}
		CacheDevelopers[detail.Login] = detail
		CacheDevelopersSet.Store(detail.Login, DataStored)
	}
	return nil
}

func CacheInsertRepo(*models.Repos) {

}

func CacheInsertDeveloper(dev *models.Developer) error {
	status, ok := CacheDevelopersSet.Load(dev.Login)
	if !ok || status != DataProcessing {
		return errors.New("developer not fetched")
	}
	detail := models.DeveloperApi{
		Login:         dev.Login,
		Type:          dev.Type,
		Location:      dev.Location,
		Languages:     dev.Languages,
		Contributions: nil,
		TalentRank:    dev.TalentRank,
		Nation:        dev.Nation,
	}
	topLang, topSize := "", int64(0)
	for lang, size := range dev.Languages {
		if size > topSize {
			topLang = lang
			topSize = size
		}
	}
	detail.TopLanguages = topLang
	cons, err := GetContributionsByDeveloper(dev.Login)
	if err != nil {
		CacheDevelopersSet.Delete(dev.Login)
		delete(CacheDevelopers, dev.Login)
		return err
	}
	detail.Contributions = cons
	CacheDevelopers[dev.Login] = detail
	CacheDevelopersSet.Store(dev.Login, DataStored)
	return nil
}
