package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ClientAddress  string `mapstructure:"CLIENT_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	CacheHost      string `mapstructure:"CACHE_HOST"`
	CachePort      string `mapstructure:"CACHE_PORT"`
	CacheUser      string `mapstructure:"CACHE_USER"`
	CachePass      string `mapstructure:"CACHE_PASSWORD"`
	CorsDomain     string `mapstructure:"CORS_DOMAIN"`
}

func NewEnv() *Env {
	file := "dev.env"
	if os.Getenv("ENV") == "production" {
		file = "production.env"
	}

	env := Env{}

	viper.AddConfigPath("/etc/discord-clone/")
	viper.SetConfigName(file)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		// if config file is not provided fall back to system variables
		env.AppEnv = viper.GetString("APP_ENV")
		env.ServerAddress = viper.GetString("SERVER_ADDRESS")
		env.ClientAddress = viper.GetString("CLIENT_ADDRESS")
		env.ContextTimeout = viper.GetInt("CONTEXT_TIMOUT")
		env.DBHost = viper.GetString("DBHost")
		env.DBPort = viper.GetString("DBPort")
		env.DBUser = viper.GetString("DB_USER")
		env.DBPassword = viper.GetString("DB_PASSWORD")
		env.DBName = viper.GetString("DB_NAME")
		env.CacheHost = viper.GetString("CACHE_HOST")
		env.CachePort = viper.GetString("CACHE_PORT")
		env.CacheUser = viper.GetString("CACHE_USER")
		env.CachePass = viper.GetString("CachePass")
		env.CorsDomain = viper.GetString("CORS_DOMAIN")
		log.Println("Can't find the env file:", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Can't decode env file: ", err)
	}

	log.Printf("domain: %v", env.CorsDomain)

	if env.AppEnv == "development" {
		log.Println("The app is running in development mode")
	} else {
		log.Println("The app is running in production mode")
	}

	return &env
}
