package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "vehicle_data_db"
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

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Route to handle queries
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		queryType := r.URL.Query().Get("type")
		var rows *sql.Rows
		switch queryType {
		case "all_data":
			rows, err = db.Query("SELECT vehicle_id, speed, speed_unit, battery, longitude, latitude, temperature, created_at FROM vehicle_data")
		case "high_speed":
			rows, err = db.Query("SELECT vehicle_id, speed, speed_unit, battery, longitude, latitude, temperature, created_at FROM vehicle_data WHERE speed >= 70")
		case "low_bat_high_speed":
			rows, err = db.Query("SELECT vehicle_id, speed, speed_unit, battery, longitude, latitude, temperature, created_at FROM vehicle_data WHERE battery <= 20 AND speed >= 70")
		default:
			http.Error(w, "Invalid query type", http.StatusBadRequest)
			return
		}

		if err != nil {
			log.Printf("Error executing query: %v", err)
			http.Error(w, "Error executing query", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []VehicleData
		for rows.Next() {
			var data VehicleData
			if err := rows.Scan(&data.VehicleID, &data.Speed, &data.SpeedUnit, &data.Battery, &data.Longitude, &data.Latitude, &data.Temperature, &data.CreatedAt); err != nil {
				http.Error(w, "Error scanning result", http.StatusInternalServerError)
				return
			}
			results = append(results, data)
		}

		fmt.Println("Query of type", queryType, "completed successfully")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux)))
}
