package models

type DeveloperApi struct {
	Login         string                `json:"login"`
	Type          string                `json:"type"`
	Location      string                `json:"location"`
	Languages     map[string]int64      `json:"languages"`
	Contributions []ContributionsStored `json:"contributions"`
	TalentRank    float64               `json:"talent_rank,string"`
	Nation        string                `json:"nation"`
}

type Developer struct {
	Id               int64            `json:"id"`
	Login            string           `json:"login"`
	Name             string           `json:"name"`
	Type             string           `json:"type"`
	Location         string           `json:"location"`
	Languages        map[string]int64 `json:"languages"`
	ContributedRepos *[]Repos         `json:"contributed_repos"`
	TalentRank       float64          `json:"talent_rank,string"`
	Nation           string           `json:"nation"`
}

type MiniDeveloper struct {
	Login         string `json:"login"`
	Id            int64  `json:"id"`
	Type          string `json:"type"`
	Nation        string `json:"nation"`
	Contributions int64  `json:"contributions"`
}

type DeveloperStored struct {
	Id           int64   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GithubId     int64   `json:"github_id" gorm:"not null;unique"`
	Login        string  `json:"login" gorm:"not null;index:idx_github_login,type:hash"`
	Name         string  `json:"name"`
	Type         string  `json:"type" gorm:"not null"`
	Location     string  `json:"location"`
	TalentRank   float64 `json:"talent_rank,string" gorm:"default:0;index:idx_talent_rank,type:btree"`
	Nation       string  `json:"nation" gorm:"not null;index:idx_nation,type:hash"`
	TopLanguages string  `json:"top_languages" gorm:"default:''"`
}
