package main

import "practice/internal/config"

// getRedisConf returns valid redis address
func getRedisConf(conf *config.AppConf) string {
	return conf.Redis.Host + ":" + conf.Redis.Port
}

// getKafkaConf returns valid kafka address
func getKafkaConf(conf *config.AppConf) string {
	return conf.Kafka.Host + ":" + conf.Kafka.Port
}

// getGRPCConf returns valid gRPC address
func getGRPCConf(conf *config.AppConf) string {
	return ":" + conf.GRPC.Port
}
