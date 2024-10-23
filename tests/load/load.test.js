import http from 'k6/http';
import { check, sleep } from 'k6';

// Test options configuration
export let options = {
  stages: [
    { duration: '30s', target: 5000 }, // Increases to 5000 VUs in 30 seconds
    { duration: '30s', target: 10000 },  // Maintains 10000 VUs for 1 minute
    { duration: '30s', target: 0 },  // Reduces to 0 VUs in 30 seconds
  ],
};

export default function () {
  // Sends a GET request to the healthcheck endpoint
  let res = http.get('http://localhost:8080/');

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  // Pauses for a short period between requests
  sleep(1);
}
