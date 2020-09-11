'use strict';

require('dotenv').config();

module.exports = {
  env: process.env.NODE_ENV || 'development',
  port: process.env.PORT || 8080,
  db: (process.env.NODE_ENV = 'test'
    ? process.env.TEST_DATABASE_URL
    : process.env.DATABASE_URL),
};
