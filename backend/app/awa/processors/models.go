package processors

import (
	"backend/app/models"
	"time"
)

type ParsedDeveloper struct {
	Login     string         `json:"login"`
	Id        int            `json:"id"`
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Company   string         `json:"company"`
	Blog      string         `json:"blog"`
	Location  string         `json:"location"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	AllRepos  *[]ParsedRepos `json:"all_repos"`
}

type ParsedRepos struct {
	Id              int                     `json:"id"`
	Name            string                  `json:"name"`
	FullName        string                  `json:"full_name"`
	Private         bool                    `json:"private"`
	Owner           *models.MiniDeveloper   `json:"owner"`
	Description     string                  `json:"description"`
	Fork            bool                    `json:"fork"`
	Languages       map[string]int64        `json:"languages"`
	Contributors    *[]models.MiniDeveloper `json:"contributors"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	PushedAt        time.Time               `json:"pushed_at"`
	Size            int                     `json:"size"`
	StargazersCount int                     `json:"stargazers_count"`
	Parent          *models.MiniRepo        `json:"parent"`
}
