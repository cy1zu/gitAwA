package postgres

import (
	"backend/app/models"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InsertDeveloper(dev *models.Developer) error {
	status, ok := CacheDevelopersSet.Load(dev.Login)
	if !ok || status != DataProcessing {
		return nil
	}
	developer := models.DeveloperStored{
		GithubId:   dev.Id,
		Login:      dev.Login,
		Name:       dev.Name,
		Type:       dev.Type,
		Company:    dev.Company,
		Blog:       dev.Blog,
		Location:   dev.Location,
		Email:      dev.Email,
		CreatedAt:  dev.CreatedAt,
		TalentRank: dev.TalentRank,
	}
	res := pdb.Create(&developer)
	if res.Error != nil {
		zap.L().Error("insert user failed", zap.Error(res.Error))
		CacheDevelopersSet.Delete(dev.Login)
		return res.Error
	}
	for _, repo := range *dev.ContributedRepos {
		err := InsertRepo(repo)
		if err != nil {
			zap.L().Debug("InsertRepo failed", zap.Error(err),
				zap.String("repo fullname", repo.FullName))
			continue
		}
		// insert contributions
		if repo.Fork == true {
			err = InsertContributions(dev.Id, dev.Login, true, repo.Parent.Id, repo.Parent.FullName, repo.Contributions, repo.TalentScore)
		} else {
			err = InsertContributions(dev.Id, dev.Login, false, repo.Id, repo.FullName, repo.Contributions, repo.TalentScore)
		}
		if err != nil {
			zap.L().Debug("InsertContributions failed", zap.Error(err),
				zap.String("dev login", dev.Login),
				zap.String("repo fullname", repo.FullName))
			continue
		}
	}
	// insert lang
	for lang, size := range dev.Languages {
		err := InsertLanguages("users", dev.Id, dev.Login, lang, size)
		if err != nil {
			zap.L().Debug("InsertLanguages failed", zap.Error(err),
				zap.String("dev login", dev.Login),
				zap.String("languages", lang),
				zap.Int64("size", size))
			continue
		}
	}
	return nil
}

func GetDeveloper(githubLogin string) (models.DeveloperStored, error) {
	developer := models.DeveloperStored{}
	res := pdb.Take(&developer, "github_login = ?", githubLogin)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return developer, ErrorDeveloperNotStored
	}
	if res.Error != nil {
		return developer, res.Error
	}
	return developer, nil
}
