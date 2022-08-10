package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"otus/crud/v1/internal/controller/web"
	"otus/crud/v1/internal/entity"
	"otus/crud/v1/internal/storage"
	"otus/crud/v1/pkg/config"
	"otus/crud/v1/pkg/db"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func main() {
	//flags
	configPath := flag.String("conf-file", "./config.yaml", "config path")
	flag.Parse()
	
	fmt.Println("Try to open config from:", *configPath)

	//configs
	v := viper.New()
	err := config.BindToEnv(v)
	if err != nil {
		panic(err)
	}
	config, err:= config.Read(v, *configPath)
	if err != nil {
		panic(err)	
	}
	
	//postgres
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
	config.PostgresDbParams.Host, config.PostgresDbParams.Port,
	config.PostgresDbParams.User, config.PostgresDbParams.Password,
	config.PostgresDbParams.Database)
	
	db, err := db.GetPostgresDb(dsn)
	if err != nil {
		panic(err)
	}

	//business
	storage := storage.New(db, false)
	usecase := entity.New(&storage)
	router := web.New(&usecase)
	routes := router.GetRoutes()
	
	//metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)
	for _, c := range router.GetPromCollectors() {
		err := reg.Register(c)
		if err != nil {
			panic(err)
		}
	}
	
	routes.Methods(http.MethodGet).
			Path("/metrics").
			Name("Prometheus").
			Handler(promhttp.HandlerFor(reg,  promhttp.HandlerOpts{}))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",config.AppParams.Host, config.AppParams.Port), routes))
}