package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
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
	// Set up PostgreSQL connection
	connStr := "user=postgres password=mysecretpassword dbname=vehicle_data_db host=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Configure Kafka consumer
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "measurements",
		GroupID:  "vehicle-consumer-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()

	log.Println("Kafka consumer started...")

	// Consume messages
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message from Kafka: ", err)
		}

		var vehicleData VehicleData
		err = json.Unmarshal(m.Value, &vehicleData)
		if err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Insert data into PostgreSQL
		_, err = db.Exec(
			"INSERT INTO vehicle_data (vehicle_id, speed, speed_unit, battery, longitude, latitude, temperature, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())",
			vehicleData.VehicleID, vehicleData.Speed, vehicleData.SpeedUnit, vehicleData.Battery, vehicleData.Longitude, vehicleData.Latitude, vehicleData.Temperature,
		)
		if err != nil {
			log.Printf("Error inserting data into PostgreSQL: %v", err)
		} else {
			log.Printf("Inserted vehicle data into PostgreSQL: %v", vehicleData)
		}
		log.Println("Received message:", string(m.Value))
	}
}
