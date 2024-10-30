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
	stored := postgres.CacheDevelopers[githubLogin]
	dev := &api.DeveloperApi{
		Id:         stored.GithubId,
		Login:      stored.Login,
		Type:       stored.Type,
		Name:       stored.Name,
		Company:    "",
		Blog:       "",
		Location:   "",
		Email:      "",
		CreatedAt:  stored.CreatedAt,
		Languages:  nil,
		TalentRank: stored.TalentRank,
	}
	if stored.Company != nil {
		dev.Company = *stored.Company
	}
	if stored.Blog != nil {
		dev.Blog = *stored.Blog
	}
	if stored.Location != nil {
		dev.Location = *stored.Location
	}
	if stored.Email != nil {
		dev.Email = *stored.Email
	}
	cons, err := postgres.GetContributionsByDeveloper(githubLogin, 0, 10)
	if err != nil {
		return nil, err
	}
	dev.Contributions = cons
	lang, err := postgres.GetLanguages("users", dev.Id)
	if err != nil {
		return nil, err
	}
	dev.Languages = lang
	return dev, nil
}
