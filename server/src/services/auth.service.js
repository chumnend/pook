'use strict';

const jwt = require('jsonwebtoken');
const config = require('../config');
const { User } = require('../models');
const { HttpError } = require('../utils');

module.exports = {
  async register(name, email, password) {
    // check if email was already taken
    const user = await User.findOne({ email });
    if (user) {
      throw new HttpError(400, 'Missing or invalid fields', {
        email: 'Email already taken',
      });
    }

    // create new user
    const newUser = await User.create({
      name,
      email,
      password,
    });

    // generate a token
    const token = jwt.sign(
      {
        id: newUser.id,
        name: newUser.name,
        email: newUser.email,
      },
      config.secret,
      { expiresIn: '24h' },
    );

    return token;
  },

  async login(email, password) {
    // find the user by email
    const user = await User.findOne({ email });
    if (!user) {
      throw new HttpError(404, 'Invalid email and/or password');
    }

    // validate the password
    const isMatch = await user.comparePassword(password);
    if (isMatch) {
      // generate a token
      const token = jwt.sign(
        {
          id: user.id,
          name: user.name,
          email: user.email,
        },
        config.secret,
        { expiresIn: '24h' },
      );

      return token;
    } else {
      throw new HttpError(400, 'Invalid email and/or password');
    }
  },

  validate(token) {
    const decoded = jwt.verify(token, config.secret);
    if (
      decoded.id !== undefined &&
      decoded.name !== undefined &&
      decoded.email !== undefined
    ) {
      return decoded;
    }

    return null;
  },
};
