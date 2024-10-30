package api

import (
	"backend/app/models"
	"time"
)

type DeveloperApi struct {
	Id            int64                        `json:"id"`
	Login         string                       `json:"login"`
	Type          string                       `json:"type"`
	Name          string                       `json:"name"`
	Company       string                       `json:"company"`
	Blog          string                       `json:"blog"`
	Location      string                       `json:"location"`
	Email         string                       `json:"email"`
	CreatedAt     time.Time                    `json:"created_at"`
	Languages     map[string]int64             `json:"languages"`
	Contributions []models.ContributionsStored `json:"contributions"`
	TalentRank    float64                      `json:"talent_rank,string"`
}
