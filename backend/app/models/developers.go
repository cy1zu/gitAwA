package models

import "time"

type Developer struct {
	Id               int              `json:"id"`
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
	Id            int    `json:"id"`
	Type          string `json:"type"`
	Nation        string `json:"nation"`
	Contributions int64  `json:"contributions"`
}

type DeveloperStored struct {
	Id         int       `json:"id"`
	Login      string    `json:"login"`
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Company    string    `json:"company"`
	Blog       string    `json:"blog"`
	Location   string    `json:"location"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	TalentRank float64   `json:"talent_rank,string"`
}
