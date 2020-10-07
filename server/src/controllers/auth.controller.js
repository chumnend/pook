'use strict';

const { validateLogin, validateRegister } = require('../validators');
const { authService } = require('../services');
const { HttpException } = require('../utils/exceptions');

module.exports = {
  async register(req, res, next) {
    try {
      const { errors, isValid } = validateRegister(req.body);
      if (!isValid) {
        throw new HttpException(400, 'Missing or invalid fields', errors);
      }

      const token = await authService.register(
        req.body.name,
        req.body.email,
        req.body.password,
      );

      res.cookie('token', token, { httpOnly: true });
      return res.status(200).json({ token });
    } catch (err) {
      return next(err);
    }
  },

  async login(req, res, next) {
    try {
      const { errors, isValid } = validateLogin(req.body);
      if (!isValid) {
        throw new HttpException(400, 'Missing or invalid fields', errors);
      }

      const token = await authService.login(req.body.email, req.body.password);

      res.cookie('token', token, { httpOnly: true });
      return res.status(200).json({ token });
    } catch (err) {
      return next(err);
    }
  },

  async logout(req, res, next) {
    res.clearCookie('token');
    return res.json({ success: true });
  },

  validate(req, res, next) {
    try {
      const token = req.cookies.token;
      if (!token) {
        return res.status(200).json({
          isValid: false,
        });
      }

      const user = authService.validate(token);
      console.log(user);
      if (user) {
        return res.status(200).json({
          isValid: true,
          user,
        });
      }

      return res.status(200).json({
        isValid: false,
      });
    } catch (err) {
      return next(err);
    }
  },
};
