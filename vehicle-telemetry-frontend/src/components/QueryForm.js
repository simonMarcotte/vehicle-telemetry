import React from 'react';
import { FormControl, InputLabel, MenuItem, Select, Button, CircularProgress } from '@mui/material';

function QueryForm({ queryType, setQueryType, handleQuery, loading }) {
  return (
    <div>
      <FormControl fullWidth variant="outlined" sx={{ marginBottom: 2 }}>
        <InputLabel>Select Query</InputLabel>
        <Select
          value={queryType}
          onChange={(e) => setQueryType(e.target.value)}
          label="Select Query"
        >
          <MenuItem value="all_data">Fetch All Data</MenuItem>
          <MenuItem value="high_speed">Fetch High Speed Vehicles</MenuItem>
          <MenuItem value="low_bat_high_speed">Fetch High Speed Low Battery Vehicles</MenuItem>
        </Select>
      </FormControl>

      <Button
        variant="contained"
        color="primary"
        onClick={handleQuery}
        disabled={!queryType || loading}
      >
        {loading ? <CircularProgress size={24} /> : 'Get Results'}
      </Button>
    </div>
  );
}

export default QueryForm;
