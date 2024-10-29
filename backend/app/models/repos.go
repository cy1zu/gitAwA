package models

import "time"

type Repos struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	FullName        string           `json:"full_name"`
	Private         bool             `json:"private"`
	Owner           *MiniDeveloper   `json:"owner"`
	Description     string           `json:"description"`
	Fork            bool             `json:"fork"`
	Languages       map[string]int64 `json:"languages"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	PushedAt        time.Time        `json:"pushed_at"`
	Size            int              `json:"size"`
	StargazersCount int              `json:"stargazers_count"`
	Parent          *MiniRepo        `json:"parent"`
	Contributions   float64          `json:"contributions,string"`
	TalentScore     float64          `json:"talent_score,string"`
}

type MiniRepo struct {
	Id              int              `json:"id"`
	FullName        string           `json:"full_name"`
	Description     string           `json:"description"`
	Contributors    *[]MiniDeveloper `json:"contributors"`
	Size            int              `json:"size"`
	StargazersCount int              `json:"stargazers_count"`
}

type RepoStored struct {
	Id                    int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GithubId              int       `json:"github_id" gorm:"not null;unique"`
	Name                  string    `json:"name" gorm:"not null"`
	FullName              string    `json:"full_name" gorm:"not null"`
	Private               bool      `json:"private" gorm:"not null"`
	OwnerId               int       `json:"owner_id" gorm:"not null"`
	OwnerLogin            string    `json:"owner_login" gorm:"not null;index:idx_owner_login,type:hash"`
	Description           *string   `json:"description"`
	Fork                  bool      `json:"fork"`
	CreatedAt             time.Time `json:"created_at" gorm:"not null;index:idx_created_at,type:btree"`
	UpdatedAt             time.Time `json:"updated_at" gorm:"not null;index:idx_updated_at,type:btree"`
	PushedAt              time.Time `json:"pushed_at"`
	Size                  int       `json:"size" gorm:"default:0"`
	StargazersCount       int       `json:"stargazers_count" gorm:"default:0;index:idx_stargazers_count,type:btree"`
	ParentId              int       `json:"parent_id"`
	ParentFullName        *string   `json:"parent_full_name"`
	ParentStargazersCount int       `json:"parent_stargazers_count"`
}
