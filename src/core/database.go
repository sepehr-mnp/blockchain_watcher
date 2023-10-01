package core

import (
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenPostgresql(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Tehran")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		log.Error(err)
	}
	log.Info("Postgresql is connected.")
	return db
}
