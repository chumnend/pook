'use strict';

const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const mongoose = require('mongoose');
const morgan = require('morgan');
const config = require('./config');

// app configuraions
const app = express();
app.use(express.urlencoded({ extended: false }));
app.use(express.json());
app.use(cors());
app.use(helmet());
if (config.env !== 'test') {
  app.use(morgan('common'));
}

// connect to database
mongoose.connect(config.db, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// route configurations
app.get('/status', (req, res) => {
  res.status(200).send('OK');
});

app.all('*', (req, res, next) => {
  const err = new Error('Path Not Found');
  err.status = 404;
  next(err);
});

app.use((err, req, res, next) => {
  return res.status(err.status || 500).json({
    error: {
      message: err.message || 'something went wrong',
      errors: {},
    },
  });
});

// start the app
app.listen(config.port, () => {
  console.log('listening on port', config.port);
});

module.exports = app;
