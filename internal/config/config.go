package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// AppConf struct contains info about redis, kafka and gRPC configurations
type AppConf struct {
	Redis struct {
		Port string
		Host string
	}
	Kafka struct {
		Port string
		Host string
	}
	GRPC struct {
		Port string
		Host string
	}
}

// ReadConf gets configuration info from config.yaml
func ReadConf(cfg *AppConf) *AppConf {
	f, err := os.Open("/app/internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Could not open config file: %s", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Could not decode config file: %s", err)
	}
	return cfg
}
