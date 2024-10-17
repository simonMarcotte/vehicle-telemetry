import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [queryResult, setQueryResult] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleQuery = async (queryType) => {
    setLoading(true);
    setError(null);
    try {
      const res = await axios.get(`http://localhost:8080/query?type=${queryType}`);
      setQueryResult(res.data);
    } catch (err) {
      setError('Error occured: ' + err.message);
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h1>Vehicle Data Queries</h1>
      <div>
        <button onClick={() => handleQuery('all_data')}>Fetch All Data</button>
        <button onClick={() => handleQuery('high_speed')}>Fetch High Speed Vehciles</button>
      </div>

      {loading && <p>Loading...</p>}
      {error && <p>{error}</p>}
      
      <div>
        <h2>Query Results:</h2>
        <ul>
          {queryResult.map((data, index) => (
            <li key={index}>
              <strong>Vehicle ID:</strong> {data.vehicle_id} <br />
              <strong>Speed:</strong> {data.speed} {data.speed_unit} <br />
              <strong>Battery:</strong> {data.battery}% <br />
              <strong>Location:</strong> ({data.latitude}, {data.longitude}) <br />
              <strong>Temperature:</strong> {data.temperature}°C <br />
              <strong>Timestamp:</strong> {data.created_at}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
