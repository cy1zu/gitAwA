package postgres

import (
	"backend/app/models"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

func InsertDeveloper(dev *models.Developer) error {
	status, ok := CacheDevelopersSet.Load(dev.Login)
	if !ok || status != DataProcessing {
		return errors.New("developer not fetched")
	}
	developer := models.DeveloperStored{
		GithubId:   dev.Id,
		Login:      dev.Login,
		Type:       dev.Type,
		TalentRank: dev.TalentRank,
		Nation:     dev.Nation,
	}
	if dev.Location == "" {
		developer.Location = dev.Nation
	}
	topLang, topSize := "", int64(0)
	for lang, size := range dev.Languages {
		if size > topSize {
			topLang = lang
			topSize = size
		}
	}
	developer.TopLanguages = strings.ToLower(topLang)
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
			err = InsertContributions(dev.Id, dev.Login, true, repo.Parent.Id, repo.Parent.FullName, repo.StargazersCount, repo.Contributions, repo.TalentScore)
		} else {
			err = InsertContributions(dev.Id, dev.Login, false, repo.Id, repo.FullName, repo.StargazersCount, repo.Contributions, repo.TalentScore)
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

func GetDevelopersList(login string, lang string, nation string) (*[]models.DeveloperStored, error) {
	var developers []models.DeveloperStored
	res := pdb.Where("top_languages like ? AND login like ? AND nation like ?",
		"%"+lang+"%",
		"%"+login+"%",
		"%"+nation+"%").
		Order("talent_rank desc").
		Find(&developers)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return &developers, nil
}
