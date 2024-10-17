
import React from 'react';
import { List, ListItem, ListItemText, Typography } from '@mui/material';

function QueryResults({ queryResult = []}) {
  const resultCount = queryResult.length;
  return (
    <div>
      <Typography variant="h5">
        Query Results: {resultCount} {resultCount === 1 ? 'result' : 'results'}
      </Typography>
      <List>
        {resultCount > 0 ? (
          queryResult.map((data, index) => (
            <ListItem key={index} sx={{ borderBottom: '1px solid gray' }}>
              <ListItemText
                primary={`Vehicle ID: ${data.vehicle_id}`}
                secondary={
                  <>
                    Speed: {data.speed} {data.speed_unit}<br />
                    Battery: {data.battery}%<br />
                    Location: ({data.latitude}, {data.longitude})<br />
                    Temperature: {data.temperature}°C<br />
                    Timestamp: {data.created_at}
                  </>
                }
              />
            </ListItem>
          ))
        ) : (
          <Typography>No results found.</Typography>
        )}
      </List>
    </div>
  );
}

export default QueryResults;