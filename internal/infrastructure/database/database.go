package database

import (
	"auth-service/internal/config"
	"auth-service/internal/util/log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB ...
type DB struct {
	CampaignDB *gorm.DB
	config     *config.Configuration
	logger     *log.Logger
}

func (b *DB) ConnectMSSLDatabase(host string) *gorm.DB {
	b.logger.Infof("[Database] Database %v connecting", host)

	var gormConfig gorm.Config
	if !b.config.IsDebugDB {
		gormConfig.Logger = logger.Discard
	}
	db, err := gorm.Open(sqlserver.Open(host), &gormConfig)

	if b.config.IsDebugDB {
		db = db.Debug()
	}

	if err != nil {
		b.logger.Fatalf("[Database] Error to connect database %v", host)
	}
	b.logger.Infof("[Database] Database %v connected", host)

	return db
}

// NewServerDB ...
func NewServerDB(c *config.Configuration, l *log.Logger) *DB {
	g := &DB{
		config: c,
		logger: l,
	}

	g.CampaignDB = g.ConnectMSSLDatabase(c.DBConCampaign)

	return g
}
