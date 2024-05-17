const express = require('express');
const bodyParser = require('body-parser');
const path = require('path');
const uuid = require('uuid');

const app = express();
const port = 8081;

app.use(bodyParser.json());

app.use(express.static(path.join(__dirname, 'public')));

app.post('/api/v1/register', (req, res) => {
    console.log('Received registration data:', req.body);
    res.send({ status: 'success' });
});

// Запуск сервера
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
