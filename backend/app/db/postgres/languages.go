package postgres

import (
	"backend/app/models"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InsertLanguages(insType string, objectId int64, objectFullName string, language string, size int64) error {
	if insType != "users" && insType != "repos" {
		return ErrorLangInsertType
	}
	langStored := &models.LanguageStored{
		Type:           insType,
		ObjectId:       objectId,
		ObjectFullName: objectFullName,
		Language:       language,
		Size:           size,
	}
	res := pdb.Create(&langStored)
	if res.Error != nil {
		zap.L().Error("insert repo_languages failed", zap.Error(res.Error))
		zap.L().Debug("insert repo_languages failed",
			zap.String("repo_fullname", objectFullName),
			zap.String("language", language))
		return res.Error
	}
	return nil
}

func GetLanguages(recType string, id int64) (map[string]int64, error) {
	var data []models.LanguageStored
	res := pdb.Find(&data, "type = ? and object_id = ?", recType, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if res.Error != nil {
		zap.L().Debug("postgres.GetLanguages failed", zap.Error(res.Error))
		return nil, res.Error
	}

	languages := make(map[string]int64)
	for _, lang := range data {
		languages[lang.Language] = lang.Size
	}
	return languages, nil
}
