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

func Init(testing string) {
	if config.Conf.PostgresConfig.Installed == false {
		zap.L().Warn("running on no database mode")
		return
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		config.Conf.PostgresConfig.Host,
		config.Conf.PostgresConfig.User,
		config.Conf.PostgresConfig.Password,
		config.Conf.PostgresConfig.Database,
		config.Conf.PostgresConfig.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		config.Conf.PostgresConfig.Using = false
		zap.L().Error("failed to connect to postgres", zap.Error(err))
		zap.L().Warn("running on no database mode")
		return
	}
	pdb = db
	config.Conf.PostgresConfig.Using = true

	if testing == "test" {
		err = db.Migrator().DropTable(
			&models.DeveloperStored{},
			&models.RepoStored{},
			&models.LanguageStored{},
			&models.ContributionsStored{})
		if err != nil {
			config.Conf.PostgresConfig.Using = false
			panic(err)
		}
	}
	err = pdb.AutoMigrate(
		&models.DeveloperStored{},
		&models.RepoStored{},
		&models.LanguageStored{},
		&models.ContributionsStored{})
	if err != nil {
		config.Conf.PostgresConfig.Using = false
		zap.L().Warn("running on no database mode")
	}

	zap.L().Info("running on database mode")
	return
}
