package config

import "github.com/spf13/viper"

type Config struct {
	//postgres
	Port   string `mapstructure:"Port"`
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
	//elastic
	Addresses string `mapstructure:"EL_ADDRESS"`
	Username  string `mapstructure:"EL_USERNAME"`
	Password  string `mapstructure:"EL_PASSWORD"`
}

func LoadConfig() (c Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return c, err
}
