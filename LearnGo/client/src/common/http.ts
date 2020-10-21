import axios from 'axios';

export default axios.create({
    baseURL: 'http://localhost:3030/api',
    headers: {
        'Content-type': 'application/json',
        // 'Access-Control-Allow-Origin': '*',
    },
});
