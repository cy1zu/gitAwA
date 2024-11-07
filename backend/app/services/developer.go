package services

import (
	"backend/app/awa"
	"backend/app/db/postgres"
	"backend/app/models"
)

func GetDeveloperServices(githubLogin string, token *string) (models.DeveloperApi, error) {
	status, ok := postgres.CacheDevelopersSet.Load(githubLogin)
	println(status, ok)
	if !ok {
		if *token != "" {
			go awa.FetchDeveloper(githubLogin, token)
		}
		return models.DeveloperApi{}, ErrorDataNeedFetch
	}
	if status == postgres.DataProcessing {
		return models.DeveloperApi{}, ErrorDataProcessing
	}
	dev := postgres.CacheDevelopers[githubLogin]
	return dev, nil
}

func GetDeveloperListServices(login string, lang string, nation string) (*[]models.DevCardApi, error) {
	storeds, err := postgres.GetDevelopersList(login, lang, nation)
	if err != nil {
		return nil, err
	}
	if storeds == nil {
		return nil, nil
	}
	data := make([]models.DevCardApi, 0, len(*storeds))
	for _, stored := range *storeds {
		dev := models.DevCardApi{
			Login:        stored.Login,
			TalentRank:   stored.TalentRank,
			TopLanguages: stored.TopLanguages,
			Nation:       stored.Nation,
		}
		data = append(data, dev)
	}
	return &data, nil
}
