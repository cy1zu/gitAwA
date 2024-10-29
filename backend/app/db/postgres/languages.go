package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertLanguages(insType string, objectId int, objectFullName string, language string, size int64) error {
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
