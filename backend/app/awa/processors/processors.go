package processors

import (
	"backend/app/awa/fetchers"
	"backend/app/models"
	"go.uber.org/zap"
)

func FinalDevelopers(dev *ParsedDeveloper) (models.Developer, error) {
	if dev == nil {
		return models.Developer{}, ErrorGetDeveloperFailed
	}
	data := models.Developer{
		Id:         dev.Id,
		Login:      dev.Login,
		Type:       dev.Type,
		Name:       dev.Name,
		Company:    dev.Company,
		Blog:       dev.Blog,
		Location:   dev.Location,
		Email:      dev.Email,
		CreatedAt:  dev.CreatedAt,
		TalentRank: 0,
	}
	if dev.AllRepos != nil {
		repos := make([]models.Repos, 0, len(*dev.AllRepos))
		languages := map[string]int64{}
		for _, repo := range *dev.AllRepos {
			parsed, err := calculateUserContributions(dev.Login, &repo)
			if err != nil {
				zap.L().Error("calculateUserContributions failed", zap.Error(err))
				zap.L().Debug("calculateUserContributions failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				return models.Developer{}, ErrorCalculateContributionsFailed
			}
			if parsed.Contributions != 0 {
				repos = append(repos, parsed)
			}
			if parsed.Fork == false {
				for lang, langSize := range repo.Languages {
					point := float64(langSize) / float64(repo.Size)
					if point >= LanguagesProportion {
						languages[lang] += langSize
					}
				}
			}
			data.TalentRank += parsed.TalentScore
		}
		data.Languages = languages
		data.ContributedRepos = &repos
	}
	return data, nil
}

func ParseDevelopersData(dev *fetchers.DeveloperFull) (*ParsedDeveloper, error) {
	if dev == nil {
		zap.L().Error("ParseDevelopersData failed: dev is nil")
		return &ParsedDeveloper{}, nil
	}
	if len(dev.AllRepos) != 0 {
		removed := removeNoContributionsRepo(dev.AllRepos)
		dev.AllRepos = removed
	}
	data := &ParsedDeveloper{
		Login:     dev.Login,
		Id:        dev.Id,
		Type:      dev.Type,
		Name:      dev.Name,
		Company:   dev.Company,
		Blog:      dev.Blog,
		Location:  dev.Location,
		Email:     dev.Email,
		CreatedAt: dev.CreatedAt,
		UpdatedAt: dev.UpdatedAt,
		AllRepos:  nil,
	}
	reposList := make([]ParsedRepos, 0, len(dev.AllRepos))

	for _, repo := range dev.AllRepos {
		if repo.Fork == false {
			detail, err := fetchers.GetReposDetail(repo.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposDetail failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposDetail failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			lang, err := fetchers.GetReposLanguages(repo.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposLanguages failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposLanguages failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			cons, err := fetchers.GetReposContributors(repo.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposContributors failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposContributors failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			reposList = append(reposList, ParsedRepos{
				Id:       detail.Id,
				Name:     detail.Name,
				FullName: detail.FullName,
				Private:  detail.Private,
				Owner: &models.MiniDeveloper{
					Login: detail.Owner.Login,
					Id:    detail.Owner.Id,
					Type:  detail.Owner.Type,
				},
				Description:     detail.Description,
				Fork:            detail.Fork,
				Languages:       lang,
				Contributors:    cons,
				CreatedAt:       detail.CreatedAt,
				UpdatedAt:       detail.UpdatedAt,
				PushedAt:        detail.PushedAt,
				Size:            detail.Size,
				StargazersCount: detail.StargazersCount,
				Parent:          nil,
			})
		} else if repo.Fork == true {
			detail, err := fetchers.GetReposDetail(repo.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposDetail failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposDetail failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			lang, err := fetchers.GetReposLanguages(detail.Parent.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposLanguages failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposLanguages failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			cons, err := fetchers.GetReposContributors(detail.Parent.FullName)
			if err != nil {
				zap.L().Error("fetchers.GetReposContributors failed", zap.Error(err))
				zap.L().Debug("fetchers.GetReposContributors failed", zap.Error(err),
					zap.String("githubId", dev.Login),
					zap.String("repoId", repo.FullName))
				continue
			}
			reposList = append(reposList, ParsedRepos{
				Id:       detail.Id,
				Name:     detail.Name,
				FullName: detail.FullName,
				Private:  detail.Private,
				Owner: &models.MiniDeveloper{
					Login: detail.Owner.Login,
					Id:    detail.Owner.Id,
					Type:  detail.Owner.Type,
				},
				Description:     detail.Description,
				Fork:            detail.Fork,
				Languages:       lang,
				Contributors:    cons,
				CreatedAt:       detail.CreatedAt,
				UpdatedAt:       detail.UpdatedAt,
				PushedAt:        detail.PushedAt,
				Size:            detail.Size,
				StargazersCount: detail.StargazersCount,
				Parent: &models.MiniRepo{
					Id:              detail.Parent.Id,
					FullName:        detail.Parent.FullName,
					Description:     detail.Parent.Description,
					Contributors:    cons,
					Size:            detail.Parent.Size,
					StargazersCount: detail.Parent.StargazersCount,
				},
			})
		}
	}
	data.AllRepos = &reposList
	return data, nil
}

func removeNoContributionsRepo(dev []fetchers.ReposFull) []fetchers.ReposFull {
	parsed := make([]fetchers.ReposFull, 0, len(dev))
	for _, repo := range dev {
		if repo.StargazersCount != 0 || repo.Fork == true {
			parsed = append(parsed, repo)
		}
	}
	return parsed
}

func calculateUserContributions(githubId string, repo *ParsedRepos) (models.Repos, error) {
	var devContribution, allContribution float64
	devContribution = 0
	allContribution = 1
	if repo.Fork == true {
		for _, dev := range *repo.Contributors {
			allContribution += float64(dev.Contributions)
			if dev.Login == githubId {
				devContribution = float64(dev.Contributions)
			}
		}
		data := models.Repos{
			Id:              repo.Id,
			Name:            repo.Name,
			FullName:        repo.FullName,
			Private:         repo.Private,
			Owner:           repo.Owner,
			Description:     repo.Description,
			Fork:            repo.Fork,
			Languages:       repo.Languages,
			CreatedAt:       repo.CreatedAt,
			UpdatedAt:       repo.UpdatedAt,
			PushedAt:        repo.PushedAt,
			Size:            repo.Size,
			StargazersCount: repo.StargazersCount + repo.Parent.StargazersCount,
			Parent:          repo.Parent,
			Contributions:   devContribution / allContribution,
		}
		data.TalentScore = float64(data.StargazersCount) * data.Contributions
		return data, nil
	} else {
		data := models.Repos{
			Id:              repo.Id,
			Name:            repo.Name,
			FullName:        repo.FullName,
			Private:         repo.Private,
			Owner:           repo.Owner,
			Description:     repo.Description,
			Fork:            repo.Fork,
			Languages:       repo.Languages,
			CreatedAt:       repo.CreatedAt,
			UpdatedAt:       repo.UpdatedAt,
			PushedAt:        repo.PushedAt,
			Size:            repo.Size,
			StargazersCount: repo.StargazersCount,
			Parent:          repo.Parent,
			Contributions:   1.0,
		}
		data.TalentScore = float64(data.StargazersCount)
		return data, nil
	}
}
