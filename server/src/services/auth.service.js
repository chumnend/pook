'use strict';

const jwt = require('jsonwebtoken');
const config = require('../config');
const { User } = require('../models');
const { HttpException } = require('../utils/exceptions');

module.exports = {
  async register(name, email, password) {
    const user = await User.findOne({ email });
    if (user) {
      throw new HttpException(400, 'User already exists', {
        email: 'Email already in use',
      });
    }

    const newUser = await User.create({
      name,
      email,
      password,
    });

    const payload = {
      id: newUser.id,
      name: newUser.name,
      email: newUser.email,
    };
    const token = jwt.sign(payload, config.secret);

    return token;
  },

  async login(email, password) {
    const user = await User.findOne({ email });
    if (!user) {
      throw new HttpException(404, 'User does not exist', {
        email: 'Email was not found',
      });
    }

    const isMatch = await user.comparePassword(password);
    if (isMatch) {
      const payload = {
        id: user.id,
        name: user.name,
        email: user.email,
      };
      const token = jwt.sign(payload, config.secret);

      return token;
    } else {
      throw new HttpException(400, 'Invalid email and/or password', {
        password: 'Incorrect password',
      });
    }
  },
};
