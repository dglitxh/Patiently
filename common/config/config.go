package config

import "github.com/spf13/viper"

type Config struct {
	Port  string `mapstructure: "PORT"`
	DBUrl string `mapstructure: "DB_URL`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./common/config/envs")
	viper.SetConfigName("dev")
	viper.setConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
