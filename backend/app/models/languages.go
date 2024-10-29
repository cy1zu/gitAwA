package models

type LanguageStored struct {
	Id             int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Type           string `json:"type" comment:"user or repo" gorm:"not null;index:idx_type,type:btree"`
	ObjectId       int    `json:"object_id" gorm:"not null;"`
	ObjectFullName string `json:"object_full_name" gorm:"not null;index:idx_object_full_name,type:hash"`
	Language       string `json:"language" gorm:"not null;index:idx_language,type:btree"`
	Size           int64  `json:"size" gorm:"default:0"`
}
