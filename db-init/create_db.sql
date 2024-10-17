-- Connect to the vehicle_data_db database
\c vehicle_data_db;

-- Create the vehicle_data table if it does not exist
CREATE TABLE IF NOT EXISTS vehicle_data (
    id SERIAL PRIMARY KEY,
    vehicle_id VARCHAR(255),
    speed FLOAT,
    speed_unit TEXT,
    battery FLOAT,
    longitude FLOAT,
    latitude FLOAT,
    temperature FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
