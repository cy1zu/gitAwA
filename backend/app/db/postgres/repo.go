package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertRepo(repo models.Repos) error {
	stored := models.RepoStored{
		GithubId:        repo.Id,
		Name:            repo.Name,
		FullName:        repo.FullName,
		Private:         repo.Private,
		OwnerId:         repo.Owner.Id,
		OwnerLogin:      repo.Owner.Login,
		Description:     &repo.Description,
		Fork:            repo.Fork,
		CreatedAt:       repo.CreatedAt,
		UpdatedAt:       repo.UpdatedAt,
		PushedAt:        repo.PushedAt,
		Size:            repo.Size,
		StargazersCount: repo.StargazersCount,
	}
	if repo.Fork == true && repo.Parent != nil {
		stored.ParentId = repo.Parent.Id
		stored.ParentFullName = &repo.Parent.FullName
		stored.ParentStargazersCount = repo.Parent.StargazersCount
	}
	res2 := pdb.Create(&stored)
	if res2.Error != nil {
		zap.L().Error("insert repo failed", zap.Error(res2.Error))
		zap.L().Debug("insert repo failed", zap.String("repo_fullname", repo.FullName))
		return res2.Error
	}
	for lang, size := range repo.Languages {
		err := InsertLanguages("repos", repo.Id, repo.FullName, lang, size)
		if err != nil {
			continue
		}
	}
	return nil
}
