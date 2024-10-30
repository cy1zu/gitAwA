package postgres

import (
	"backend/app/models"
	"backend/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pdb *gorm.DB

func Init(testing string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		config.Conf.PostgresConfig.Host,
		config.Conf.PostgresConfig.User,
		config.Conf.PostgresConfig.Password,
		config.Conf.PostgresConfig.Database,
		config.Conf.PostgresConfig.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect to postgres", zap.Error(err))
		return ErrorInitDatabaseFailed
	}
	pdb = db

	if testing == "test" {
		err = db.Migrator().DropTable(
			&models.DeveloperStored{},
			&models.RepoStored{},
			&models.LanguageStored{},
			&models.ContributionsStored{})
		if err != nil {
			panic(err)
		}
	}

	err = pdb.AutoMigrate(
		&models.DeveloperStored{},
		&models.RepoStored{},
		&models.LanguageStored{},
		&models.ContributionsStored{})
	if err != nil {
		zap.L().Error("data migrate failed", zap.Error(err))
		return err
	}
	err = CacheInit()
	if err != nil {
		zap.L().Error("init cache failed", zap.Error(err))
		return err
	}
	zap.L().Info("running on database mode")
	return nil
}
