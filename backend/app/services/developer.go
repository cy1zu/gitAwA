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

func GetLanguageListServices(lang string, page int) (*[]models.DeveloperApi, error) {
	start := (page - 1) * 10
	end := start + 10
	data, err := postgres.GetDevelopersListByLanguages(lang, start, end)
	if err != nil {
		return nil, err
	}
	return data, nil
}
