import React from 'react';
import { Typography } from '@mui/material';

function ErrorDisplay({ error }) {
  return (
    error && (
      <Typography color="error" sx={{ marginTop: 2 }}>
        {error}
      </Typography>
    )
  );
}

export default ErrorDisplay;
