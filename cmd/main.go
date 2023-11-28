package main

import (
	"context"
	"log"
	"os"

	"practice/internal/grpc_server"
	myKafka "practice/internal/kafka"
	"practice/internal/redis"
	"practice/storage"
	"practice/storage/postgres"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Application contains info about application services
type Application struct {
	channel  chan *storage.Car
	redis    *redis.Client
	postgres *postgres.TelematicDB
	kafka    *myKafka.Producer
	gRPC     *grpc_server.Server
}

func main() {
	log.Println("Start telematic service...")

	conf := getConfig() // app settings

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := postgres.New(ctx, os.Getenv("DATABASE_URL")) // connect to db
	if err != nil {
		cancel()
		log.Fatalf("Could not connect to postgres db: %s", err)
	}

	newRedis := redis.New(getRedisConf(conf)) // create redis client

	newKafkaProducer, err := myKafka.NewProducer(getKafkaConf(conf)) // create kafka producer
	if err != nil {
		cancel()
		log.Fatalf("Could not connect to kafka: %s", err)
	}

	telematicCh := make(chan *storage.Car) // telematics transmission channel

	app := &Application{
		channel:  telematicCh,
		redis:    newRedis,
		postgres: db,
		kafka:    newKafkaProducer,
	}

	err = app.refreshTelematicTable(ctx)
	if err != nil {
		cancel()
		log.Fatalf("Could not refresh the table: %s", err)
	}
	app.startApp(ctx)

	// создание консюмера кафки
	log.Println("start consumer")

	config2 := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(config2)
	if err != nil {
		panic(err)
	}
	consumer.SubscribeTopics([]string{"telematicTopic"}, nil)
	go func() {
		for {
			_, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Fatalf("kafka: %s", err)
			}
		}
		consumer.Close()
	}()

	app.startServer(getGRPCConf(conf))
}
