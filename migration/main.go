package main

import (
	"flag"
	"fmt"
	"otus/crud/v1/internal/storage"
	"otus/crud/v1/pkg/config"
	"otus/crud/v1/pkg/db"

	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("conf-file", "./config.yaml", "config path")
	flag.Parse()
	fmt.Println("Try to open config from:", *configPath)
	viper.SetConfigFile(*configPath)
	viper.BindEnv("db.name", "DB_NAME")
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.user", "DB_USER")
	viper.BindEnv("db.pass", "DB_PASS")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config config.Config
	viper.Unmarshal(&config)
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
	config.PostgresDbParams.Host, config.PostgresDbParams.Port,
	config.PostgresDbParams.User, config.PostgresDbParams.Password,
	config.PostgresDbParams.Database)
	
	db, err := db.GetPostgresDb(dsn)
	if err != nil {
		panic(err)
	}
	storage.New(db, true)
	//usecase := entity.New(&storage)
	//router := web.New(&usecase)

	//log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",config.AppParams.Host, config.AppParams.Port), router.GetRoutes()))
}
