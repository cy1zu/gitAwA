package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertUser(dev models.Developer) error {
	developer := models.DeveloperStored{
		GithubId:   dev.Id,
		Login:      dev.Login,
		Name:       dev.Name,
		Type:       dev.Type,
		Company:    &dev.Company,
		Blog:       &dev.Blog,
		Location:   &dev.Location,
		Email:      &dev.Email,
		CreatedAt:  dev.CreatedAt,
		TalentRank: dev.TalentRank,
	}
	res := pdb.Create(&developer)
	if res.Error != nil {
		zap.L().Error("insert user failed", zap.Error(res.Error))
		return res.Error
	}
	for _, repo := range *dev.ContributedRepos {
		err := InsertRepo(repo)
		if err != nil {
			continue
		}
		// insert contributions
		err = InsertContributions(dev.Id, dev.Login, repo.Id, repo.FullName, repo.Contributions, repo.TalentScore)
		if err != nil {
			continue
		}
	}
	// insert lang
	for lang, size := range dev.Languages {
		err := InsertLanguages("users", dev.Id, dev.Login, lang, size)
		if err != nil {
			continue
		}
	}
	return nil
}
