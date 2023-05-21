package lib

import (
	"github.com/spf13/viper"
)

// Env is a struct that contains all the environment variables
type Env struct {
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

// NewEnv returns a new Env struct
func NewEnv() *Env {
	env := &Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		panic(err)
	}

	return env
}
