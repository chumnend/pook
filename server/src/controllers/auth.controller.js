'use strict';

const { validateLogin, validateRegister } = require('../validators');
const { authService } = require('../services');
const { HttpException } = require('../utils/exceptions');

module.exports = {
  async register(req, res, next) {
    try {
      const { errors, isValid } = validateRegister(req.body);
      if (!isValid) {
        throw new HttpException(400, 'invalid inputs', errors);
      }

      const token = await authService.register(
        req.body.name,
        req.body.email,
        req.body.password,
      );

      return res.status(200).json({ token });
    } catch (err) {
      return next(err);
    }
  },

  async login(req, res, next) {
    try {
      const { errors, isValid } = validateLogin(req.body);
      if (!isValid) {
        throw new HttpException(400, 'invalid inputs', errors);
      }

      const token = await authService.login(req.body.email, req.body.password);

      return res.status(200).json({ token });
    } catch (err) {
      return next(err);
    }
  },
};
