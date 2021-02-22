package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configurations struct {
	ServerAddress              string
	DBHost                     string
	DBName                     string
	DBUser                     string
	DBPass                     string
	DBPort                     string
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int
}

func NewConfigurations(logger *logrus.Logger) *Configurations {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:8500")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "pomodorogo-server")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("ACCESS_TOKEN_PRIVATE_KEY_PATH", "./access-private.pem")
	viper.SetDefault("ACCESS_TOKEN_PUBLIC_KEY_PATH", "./access-public.pem")
	viper.SetDefault("REFRESH_TOKEN_PRIVATE_KEY_PATH", "./refresh-private.pem")
	viper.SetDefault("REFRESH_TOKEN_PUBLIC_KEY_PATH", "./refresh-public.pem")
	viper.SetDefault("JWT_EXPIRATION", 30)

	configs := &Configurations{
		ServerAddress:              viper.GetString("SERVER_ADDRESS"),
		DBHost:                     viper.GetString("DB_HOST"),
		DBName:                     viper.GetString("DB_NAME"),
		DBUser:                     viper.GetString("DB_USER"),
		DBPass:                     viper.GetString("DB_PASSWORD"),
		DBPort:                     viper.GetString("DB_PORT"),
		AccessTokenPrivateKeyPath:  viper.GetString("ACCESS_TOKEN_PRIVATE_KEY_PATH"),
		AccessTokenPublicKeyPath:   viper.GetString("ACCESS_TOKEN_PUBLIC_KEY_PATH"),
		RefreshTokenPrivateKeyPath: viper.GetString("REFRESH_TOKEN_PRIVATE_KEY_PATH"),
		RefreshTokenPublicKeyPath:  viper.GetString("REFRESH_TOKEN_PUBLIC_KEY_PATH"),
		JwtExpiration:              viper.GetInt("JWT_EXPIRATION"),
	}

	logger.Debug("server port: ", configs.ServerAddress)
	logger.Debug("db host: ", configs.DBHost)
	logger.Debug("db name: ", configs.DBName)
	logger.Debug("db port: ", configs.DBPort)
	logger.Debug("jwt expiration: ", configs.JwtExpiration)

	return configs
}
