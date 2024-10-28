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
