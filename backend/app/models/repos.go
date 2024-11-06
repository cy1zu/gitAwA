package models

type Repos struct {
	Id              int64            `json:"id"`
	Name            string           `json:"name"`
	FullName        string           `json:"full_name"`
	Private         bool             `json:"private"`
	Owner           *MiniDeveloper   `json:"owner"`
	Fork            bool             `json:"fork"`
	Languages       map[string]int64 `json:"languages"`
	Size            int64            `json:"size"`
	StargazersCount int64            `json:"stargazers_count"`
	Parent          *MiniRepo        `json:"parent"`
	Contributions   float64          `json:"contributions,string"`
	TalentScore     float64          `json:"talent_score,string"`
}

type MiniRepo struct {
	Id              int64            `json:"id"`
	FullName        string           `json:"full_name"`
	Contributors    *[]MiniDeveloper `json:"contributors"`
	Size            int64            `json:"size"`
	StargazersCount int64            `json:"stargazers_count"`
}

type RepoStored struct {
	Id                    int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GithubId              int64  `json:"github_id" gorm:"not null;unique;index:idx_github_id,type:hash"`
	Name                  string `json:"name" gorm:"not null"`
	FullName              string `json:"full_name" gorm:"not null"`
	Private               bool   `json:"private" gorm:"not null"`
	OwnerId               int64  `json:"owner_id" gorm:"not null"`
	OwnerLogin            string `json:"owner_login" gorm:"not null;index:idx_owner_login,type:hash"`
	Fork                  bool   `json:"fork" gorm:"not null"`
	Size                  int64  `json:"size" gorm:"default:0"`
	StargazersCount       int64  `json:"stargazers_count" gorm:"default:0;index:idx_stargazers_count,type:btree"`
	ParentId              int64  `json:"parent_id"`
	ParentFullName        string `json:"parent_full_name"`
	ParentStargazersCount int64  `json:"parent_stargazers_count"`
}
