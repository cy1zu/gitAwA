package models

import "time"

type DeveloperApi struct {
	Id            int64                 `json:"id"`
	Login         string                `json:"login"`
	Type          string                `json:"type"`
	Name          string                `json:"name"`
	Company       string                `json:"company"`
	Blog          string                `json:"blog"`
	Location      string                `json:"location"`
	Email         string                `json:"email"`
	CreatedAt     time.Time             `json:"created_at"`
	Languages     map[string]int64      `json:"languages"`
	Contributions []ContributionsStored `json:"contributions"`
	TalentRank    float64               `json:"talent_rank,string"`
}

type Developer struct {
	Id               int64            `json:"id"`
	Login            string           `json:"login"`
	Type             string           `json:"type"`
	Name             string           `json:"name"`
	Company          string           `json:"company"`
	Blog             string           `json:"blog"`
	Location         string           `json:"location"`
	Email            string           `json:"email"`
	CreatedAt        time.Time        `json:"created_at"`
	Languages        map[string]int64 `json:"languages"`
	ContributedRepos *[]Repos         `json:"contributed_repos"`
	TalentRank       float64          `json:"talent_rank,string"`
}

type MiniDeveloper struct {
	Login         string `json:"login"`
	Id            int64  `json:"id"`
	Type          string `json:"type"`
	Nation        string `json:"nation"`
	Contributions int64  `json:"contributions"`
}

type DeveloperStored struct {
	Id         int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GithubId   int64     `json:"github_id" gorm:"not null;unique"`
	Login      string    `json:"login" gorm:"not null;index:idx_github_login,type:hash"`
	Type       string    `json:"type" gorm:"not null"`
	Name       string    `json:"name"`
	Company    string    `json:"company"`
	Blog       string    `json:"blog"`
	Location   string    `json:"location"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;index:idx_created_at,type:btree"`
	TalentRank float64   `json:"talent_rank,string" gorm:"default:0;index:idx_talent_rank,type:btree"`
}
