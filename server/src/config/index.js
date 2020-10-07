'use strict';

require('dotenv').config();

module.exports = {
  env: process.env.NODE_ENV || 'development',
  port: process.env.PORT || 8081,
  secret: process.env.SECRET || 'shhh',
  db:
    process.env.NODE_ENV === 'test'
      ? process.env.TEST_DATABASE_URL || 'mongodb://localhost/hotelio_test'
      : process.env.DATABASE_URL || 'mongodb://localhost/hotelio_dev',
  client: process.env.CLIENT_URL || 'http://localhost:8080',
};
