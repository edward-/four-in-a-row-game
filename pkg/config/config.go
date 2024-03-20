package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App   App
		Db    Db
		Cache Cache
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

	Cache struct {
		Host         string
		Port         int
		Password     string
		Db           int
		Protocol     int
		MinIdleConns int
		PoolSize     int
		PoolTimeout  int
	}
)

func GetConfig() Config {
	configFileName := getConfigFileName()
	configFolderLocation := getConfigFolder()

	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFolderLocation)

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
		Cache: Cache{
			Host:         viper.GetString("cache.host"),
			Port:         viper.GetInt("cache.port"),
			Password:     viper.GetString("cache.password"),
			Db:           viper.GetInt("cache.db"),
			Protocol:     viper.GetInt("cache.protocol"),
			MinIdleConns: viper.GetInt("cache.minIdleConns"),
			PoolSize:     viper.GetInt("cache.poolSize"),
			PoolTimeout:  viper.GetInt("cache.poolTimeout"),
		},
	}
}

func getConfigFolder() string {
	configFolder := os.Getenv("CONFIG_FOLDER")
	if configFolder == "" {
		return "./config"
	}
	return configFolder
}

func getConfigFileName() string {
	env := os.Getenv("GO_ENVIRONMENT")

	switch env {
	case "prod":
		return "config.production"
	case "dev":
		return "config.dev"
	case "test":
		return "config.test"
	case "docker":
		return "config.docker"
	default:
		return "config.dev"
	}
}
