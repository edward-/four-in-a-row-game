package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

func GetConfig() Config {
	configFileName := getConfigFileName()

	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	return Config{
		App: App{
			Port: viper.GetInt("app.server.port"),
		},
		Db: Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
	}
}

func getConfigFileName() string {
	env := os.Getenv("GO_ENVIRONMENT")

	switch env {
	case "prod":
		return "config.production"
	case "development":
		return "config.development"
	case "test":
		return "config.test"
	default:
		return "config.development"
	}
}