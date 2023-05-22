package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	APP appConfig
	DB  dbConfig
	Email email
	SenGridConfig sengridConfig
}

type appConfig struct {
	Port       string
	Name       string
	GenerateDB bool
	Views      string
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type email struct {
	From string
	To   string
}

type sengridConfig struct {
	SengridApiKey string
}

func init() {
	viper.SetDefault("app.port", "9000")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.APP = appConfig{
		Port:       viper.GetString("app.port"),
		Name:       viper.GetString("app.name"),
		GenerateDB: viper.GetBool("app.generateDB"),
		Views:      viper.GetString("app.views"),
		
	}

	cfg.Email = email{
		From:  viper.GetString("email.from"),
		To:    viper.GetString("email.to"),

	}
	cfg.DB = dbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	cfg.SenGridConfig = sengridConfig{
		SengridApiKey: viper.GetString("sengrid.key"),
	}

	return nil
}

/**
 * Retorna informações do app
 */
func GetApp() appConfig {
	return cfg.APP
}

/**
 * Retornar conexão com banco de dados
 */
func GetDB() dbConfig {
	return cfg.DB
}

/**
 * Retornar email informaçẽos
 */
func GetEmail() email{
	return cfg.Email
}

func GetSengridApiKey() string {
	return cfg.SenGridConfig.SengridApiKey
}