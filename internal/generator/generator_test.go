package generator

import (
	"fmt"
	"log"
	"practice/storage"
	"time"

	//"github.com/golang/protobuf/protoc-gen-go/generator"
	"math/rand"
	"testing"
	//"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestGenerator(t *testing.T) {
	kafkaCh := make(chan *storage.Car)
	log.Println("Starting generator...")
	go Generate(rand.Intn(3), kafkaCh)

	go func() {
		for {
			val, ok := <-kafkaCh
			if ok == false {
				break // exit break loop
			} else {
				fmt.Println(val, ok)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(kafkaCh)
}
