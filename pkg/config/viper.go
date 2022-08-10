package config

import (
	"github.com/spf13/viper"
)

func Read(v *viper.Viper, path string) (*Config, error) {
	v.SetConfigFile(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = v.Unmarshal(&config)
	return &config, err
}

func BindToEnv(v *viper.Viper) error {
	if err := v.BindEnv("app.host", "APP_HOST"); err != nil {
		return err;
	}
	if err := v.BindEnv("app.port", "APP_PORT"); err != nil {
		return err;
	} 
	if err := v.BindEnv("db.name", "DB_NAME"); err != nil {
		return err;
	} 
	if err := v.BindEnv("db.host", "DB_HOST"); err != nil {
		return err;
	} 
	if err := v.BindEnv("db.port", "DB_PORT"); err != nil {
		return err;
	} 
	if err := v.BindEnv("db.user", "DB_USER"); err != nil {
		return err;
	} 
	if err := v.BindEnv("db.pass", "DB_PASS"); err != nil {
		return err;
	}
	return nil;
}