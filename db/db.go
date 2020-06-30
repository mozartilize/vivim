package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var Database *sqlx.DB

func GetDb(config *viper.Viper) (*sqlx.DB, error) {
	var err error
	if Database == nil {
		Database, err = sqlx.Open("postgres", config.GetString("database_url"))
		if err != nil {
			return nil, err
		}
		Database.SetMaxOpenConns(config.GetInt("database_max_connections"))
		Database.SetMaxIdleConns(config.GetInt("database_pool_size"))
		Database.SetConnMaxLifetime(time.Duration(config.GetUint("database_pool_recycle") * 1e9))
	}
	return Database, nil
}
