package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertRepo(repo models.Repos) error {
	status, ok := CacheReposSet.Load(repo.FullName)
	if ok && (status == DataStored || status == DataProcessing) {
		return nil
	}
	CacheReposSet.Store(repo.FullName, DataProcessing)
	stored := models.RepoStored{
		GithubId:        repo.Id,
		Name:            repo.Name,
		FullName:        repo.FullName,
		Private:         repo.Private,
		OwnerId:         repo.Owner.Id,
		OwnerLogin:      repo.Owner.Login,
		Fork:            repo.Fork,
		Size:            repo.Size,
		StargazersCount: repo.StargazersCount,
	}
	if repo.Fork == true && repo.Parent != nil {
		stored.ParentId = repo.Parent.Id
		stored.ParentFullName = repo.Parent.FullName
		stored.ParentStargazersCount = repo.Parent.StargazersCount
	}
	res2 := pdb.Create(&stored)
	if res2.Error != nil {
		zap.L().Error("insert repo failed", zap.Error(res2.Error))
		zap.L().Debug("insert repo failed", zap.String("repo_fullname", repo.FullName))
		CacheReposSet.Delete(repo.FullName)
		return res2.Error
	}
	for lang, size := range repo.Languages {
		err := InsertLanguages("repos", repo.Id, repo.FullName, lang, size)
		if err != nil {
			continue
		}
	}
	CacheReposSet.Store(repo.FullName, DataStored)
	CacheRepos[repo.FullName] = stored
	return nil
}

func GetRepoByRepoId(repoId int64) (models.RepoStored, error) {
	repo := models.RepoStored{}
	res := pdb.Take(&repo, "github_id = ?", repoId)
	if res.Error != nil {
		return repo, res.Error
	}
	return repo, nil
}
