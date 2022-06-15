package config

import "github.com/spf13/viper"

type Config struct {
	PORT   int    `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
