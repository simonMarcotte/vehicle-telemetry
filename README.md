# Real-Time Vehicle Telemetry Server

A real-time vehicle telemetry data streaming service that simulates IoT data from electric vehicles. The project leverages Go for backend processing, Apache Kafka for message brokering, PostgreSQL for data storage, and React for the frontend UI. The system simulates electric vehicle telemetry data (such as speed, battery level, and GPS coordinates), streams it via Kafka, processes the data, and stores it in PostgreSQL for analysis.

### Prerequisites

- Docker Desktop Installed for the docker daemon.

### Running the Project

1. **Clone the Repository**:
2. **Build and run services**
    ```
    docker compose up -d --build
    ```
3. **Access Services**
    - React frontend is set up on localhost:3000 to query the simulated data with a UI.
    - Go API is available at localhost:8080 to perform raw queries on the db.

## API Endpoints

### ``` GET /query ```
- Description: Query Vehicle data from the PostgreSQL db.
- Example: ```http://localhost:8080/query?type=all_data```
- Parameters: 
    - ```type=all_data```: Retrieves all vehicle data.
    - ```type=high_speed```: Retrieves data of vehicles travelling over 70 km/h.
    - ```type=low_bat_high_speed```: Retrieves data of vehicles with less than 20% battery and travelling over 70 km/h.

## Troubleshooting
Common Issues:
- Database Not Created: If the vehicle_data table is not created, manually run the create_db.sql script inside the PostgreSQL container.
    - First run: ```psql -U postgres -d vehicle_data_db``` in the postgres docker container
    - Then run: ```\i /docker-entrypoint-initdb.d/create_db.sql```
- Consumer shut down: This can happen if the consumer is still waiting for kafka to start up. Simply run the container again and it should work.

## Future Improvements
- Authenitcation for Go API
- Data Visualization

