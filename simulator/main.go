package main

import (
	"github.com.fabioqmarsiaj/simulator/infra/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	//route := route2.Route{
	//	ID:       "1",
	//	ClientID: "1",
	//}
	//
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//fmt.Print(stringjson[1])

	producer := kafka.NewKafkaProducer()
	err := kafka.Publish("Olá", "readtest", producer)
	if err != nil {
		return
	}

	for {
		_ = 1
	}
}
