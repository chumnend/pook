'use strict';

const jwt = require('jsonwebtoken');
const config = require('../config');
const { HttpException } = require('../utils/exceptions');

module.exports = (req, res, next) => {
  try {
    // read token from cookie
    const token = req.cookies.token;
    if (!token) {
      const err = new HttpException(401, 'Unauthorized. Missing token');
      return next(err);
    }

    const decoded = jwt.verify(token, config.secret);
    req.user = decoded;
    next();
  } catch (err) {
    res.clearCookie('cookie');
    return next(err);
  }
};
