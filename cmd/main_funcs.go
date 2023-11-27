package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"strconv"

	"practice/internal/config"
	"practice/internal/generator"
	"practice/internal/grpc_server"
	myKafka "practice/internal/kafka"
	"practice/internal/service"

	"github.com/jackc/pgx/v4"
)

// refreshTelematicTable method recreates telematic table
func (a *Application) refreshTelematicTable(ctx context.Context) error {
	err := service.DropTelematicTable(ctx, a.postgres) // создание новой таблицы
	if err != nil {
		log.Println("Could not drop the table", err)
		return err
	}

	err = service.CreateTelematicTable(ctx, a.postgres) // создание новой таблицы
	if err != nil {
		log.Println("Could not create the table", err)
		return err
	}
	return nil
}

// startApp method starts generator and calls functions for adding telematic data to redis, db and kafka
func (a *Application) startApp(ctx context.Context) {
	amount, err := strconv.Atoi(os.Getenv("TRANSPORT_AMOUNT")) // количество машин
	if err != nil {
		log.Fatal("Transport amount must be an integer")
	}

	for i := 0; i < amount; i++ {
		go generator.GenerateTelematic(i, a.channel)
	}

	key := 0
	go func() {
		for {
			val, ok := <-a.channel
			if ok == false {
				close(a.channel)
				break // exit break loop
			} else {
				val.ID = key
				err := a.redis.AddToRedis(val, key) // добавление в редис
				if err != nil {
					log.Fatalf("Could not add element to redis: %s", err)
				}
				err = service.AddTelematic(ctx, a.postgres, val) // добавление в бд
				if err != nil && !errors.Is(pgx.ErrNoRows, err) {
					log.Fatalf("Could not add element to db: %s", err)
				}
				err = myKafka.Produce(a.kafka, val) // отправка сообщения в кафку
				if err != nil {
					log.Fatalf("Could not add element to kafka: %s", err)
				}
				key++
			}
		}
	}()
}

// startServer method starts gRPC server
func (a *Application) startServer(addr string) {
	lis, err := net.Listen("tcp", addr) // создание gRPC сервера
	if err != nil {
		log.Fatalf("Failed to listen to port: %s", err)
	}
	srv := grpc_server.New(a.postgres)
	if err := srv.Grpc.Serve(lis); err != nil {
		log.Fatalf("Could not start grpc server: %s", err)
	}
}

// getConfig method gets configuration from internal/config/config.yaml
func getConfig() *config.AppConf {
	var conf *config.AppConf
	conf = config.ReadConf(conf)
	return conf
}
