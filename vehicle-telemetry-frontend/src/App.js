import React, { useState } from 'react';
import axios from 'axios';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Container, Box, CssBaseline } from '@mui/material';
import QueryForm from './components/QueryForm'
import QueryResults from './components/QueryResults';
import ErrorDisplay from './components/ErrorDisplay';

const theme = createTheme({
  palette: {
    mode: 'dark',
    background: {
      default: '#121212',
      paper: '#121212',
    },
    text: {
      primary: '#ffffff',
    },
  },
});

function App() {
  const [queryResult, setQueryResult] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [queryType, setQueryType] = useState('');
  const [queryMade, setQueryMade] = useState(false);

  const handleQuery = async () => {
    setLoading(true);
    setError(null);
    setQueryMade(false);
    try {
      const res = await axios.get(`http://localhost:8080/query?type=${queryType}`);
      setQueryResult(res.data || []); 
      setQueryMade(true);
    } catch (err) {
      setError('Error occurred: ' + err.message);
      setQueryMade(true);
    } finally {
      setLoading(false);
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Container>
        <Box sx={{ marginBottom: 4 }}>
          <h1>Vehicle Telemetry Data Queries</h1>
        </Box>

        {/* Query Form */}
        <QueryForm
          queryType={queryType}
          setQueryType={setQueryType}
          handleQuery={handleQuery}
          loading={loading}
        />

        {/* Only show "Query Results" title and content if a query was made */}
        {queryMade && (
          <Box sx={{ marginTop: 4 }}>
            <ErrorDisplay error={error} />
            <QueryResults queryResult={queryResult} />
          </Box>
        )}
      </Container>
    </ThemeProvider>
  );
}

export default App;