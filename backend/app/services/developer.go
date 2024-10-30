package services

import (
	"backend/app/api"
	"backend/app/awa"
	"backend/app/db/postgres"
)

func GetDeveloperServices(githubLogin string, token *string) (*api.DeveloperApi, error) {
	status, ok := postgres.CacheDevelopersSet.Load(githubLogin)
	if !ok {
		if *token != "" {
			go awa.FetchDeveloper(githubLogin, token)
		}
		return &api.DeveloperApi{}, ErrorDataNeedFetch
	}
	if status == postgres.DataProcessing {
		return &api.DeveloperApi{}, ErrorDataProcessing
	}
	dev := postgres.CacheDevelopers[githubLogin]

	return dev, nil
}
