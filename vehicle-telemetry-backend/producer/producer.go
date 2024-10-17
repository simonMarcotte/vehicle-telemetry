package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type VehicleData struct {
	VehicleID   string  `json:"vehicle_id"`
	Speed       float64 `json:"speed"`
	SpeedUnit   string  `json:"speed_unit"`
	Battery     float64 `json:"battery"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Temperature float64 `json:"temperature"`
	CreatedAt   string  `json:"created_at"` // Make sure this field is defined
}

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"}, // Ensure the broker address is correct
		Topic:    "measurements",
		Balancer: &kafka.LeastBytes{},
	})

	// Handle the case where writer is not initialized
	if writer == nil {
		log.Fatal("Kafka writer is not initialized")
	}

	defer func() {
		err := writer.Close()
		if err != nil {
			log.Printf("Error closing Kafka writer: %v\n", err)
		}
	}()

	for {
		vehicleData := generateVehicleData()

		data, err := json.Marshal(vehicleData)
		if err != nil {
			log.Printf("Error marshaling vehicle data: %v\n", err)
			continue
		}

		// Send the message to Kafka
		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Value: data,
			})
		if err != nil {
			log.Printf("Error writing message to Kafka: %v\n", err)
		} else {
			fmt.Printf("Sent vehicle data: %s\n", data)
		}

		time.Sleep(2 * time.Second)
	}
}

func generateVehicleData() VehicleData {
	return VehicleData{
		VehicleID:   fmt.Sprintf("EV-%s", uuid.New().String()),
		Speed:       rand.Float64() * 100,
		SpeedUnit:   "km/h",
		Battery:     0 + rand.Float64()*(100),
		Longitude:   -180 + rand.Float64()*360,
		Latitude:    -90 + rand.Float64()*180,
		Temperature: -10 + rand.Float64()*40,
	}
}
