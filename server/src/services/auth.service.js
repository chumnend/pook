'use strict';

const jwt = require('jsonwebtoken');
const config = require('../config');
const { User } = require('../models');
const { HttpException } = require('../utils/exceptions');

module.exports = {
  async register(name, email, password) {
    try {
      const user = await User.findOne({ email });
      if(user) {
        throw new HttpException(400, 'User already exists');
      }

      const newUser = await User.create({
        name,
        email,
        password,
      })

      return newUser;
    } catch (err) {
      throw err;
    }
  },

  async login(email, password) {
    try {
      const user = await User.findOne({ email });
      if(!user) {
        throw new HttpException(404, 'User does not exist');
      }

      const isMatch = await user.comparePassword(password);
      if(isMatch) {
        const payload = { id: user.id, name: user.name };
        const token = jwt.sign(payload, config.secret);

        return token;
      } else {
        throw new HttpException(400, 'Invalid email and/or password');
      }
    } catch(err) {
      throw err;
    }
  },
};
