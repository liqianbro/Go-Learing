package mysqldb

import (
	"Go-Learing/ent"
	"database/sql"
	"log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *ent.Client

type MySQLConfig struct {
	Type         string
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}

func ReadMySQLConfig() MySQLConfig {
	var config MySQLConfig
	config.Dsn = viper.GetString("db.mysql.dsn")
	config.Type = viper.GetString("db.mysql.type")
	config.MaxIdleConns = 10
	config.MaxOpenConns = 100
	return config
}

func InitMySQLClient(config MySQLConfig) error {
	db, err := sql.Open(config.Type, config.Dsn)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(config.Type, db)
	DB = ent.NewClient(ent.Driver(drv))
	return nil
}

func NewViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
		return err
	}
	return nil
}
