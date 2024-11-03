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
