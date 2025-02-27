package api

import (
	"log"

	"github.com/sachatarba/course-db/internal/config"
	handler "github.com/sachatarba/course-db/internal/delivery/rest"
	"github.com/sachatarba/course-db/internal/orm"
	postrgres_adapter "github.com/sachatarba/course-db/internal/postrgres"
	redis_adapter "github.com/sachatarba/course-db/internal/redis"
	"github.com/sachatarba/course-db/internal/server"
)

type ApiServer struct {
	conf config.Config
}

func (api *ApiServer) Run() {
	conf := config.NewConfFromEnv()

	postgresConnector := postrgres_adapter.PostgresConnector{
		Conf: conf.PostgresConf,
	}
	redisConnector := redis_adapter.RedisConnector{
		Conf: conf.RedisConf,
	}

	db, err := postgresConnector.Connect()
	if err != nil {
		log.Print("Cant connect postgres", err)
		return
	}

	rdb := redisConnector.Connect()
	if rdb == nil {
		log.Print("Cant connect redis", err)
		return
	}

	postgresMigrator := postrgres_adapter.PostgresMigrator{
		DB:     db,
		Tables: orm.TablesORM,
	}

	err = postgresMigrator.Migrate()
	if err != nil {
		log.Print("Cant connect migrate", err)
		return
	}

	paymentConfig := conf.PaymentConfig
	paymentHandler := handler.NewPaymentHandler(paymentConfig)

	director := ApiServicesDirector{
		Builder: &ApiServicesBuilder{
			Postgres: db,
			Redis:    rdb,
		},
	}

	services, err := director.NewServices()
	if err != nil {
		log.Print("Cant create services", err)
		return
	}

	server := server.Server{
		PaymentHandler: paymentHandler,
		Handler: handler.NewHandler(*services),
		Conf:    &config.ServerConfig{},
	}

	server.Run()
}
