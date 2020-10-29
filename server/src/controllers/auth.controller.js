'use strict';

const { validateLogin, validateRegister } = require('../validators');
const { authService } = require('../services');
const { HttpError } = require('../utils');

module.exports = {
  async register(req, res, next) {
    try {
      const { errors, isValid } = validateRegister(req.body);
      if (!isValid) {
        throw new HttpError(400, 'Missing or invalid fields', errors);
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
        throw new HttpError(400, 'Missing or invalid fields', errors);
      }

      const token = await authService.login(req.body.email, req.body.password);

      res.cookie('token', token, { httpOnly: true });
      return res.status(200).json({ token });
    } catch (err) {
      return next(err);
    }
  },

  logout(req, res, next) {
    res.clearCookie('token');
    return res.send();
  },

  validate(req, res, next) {
    try {
      const token = req.cookies.token;
      if (!token) {
        return res.status(200).json({ success: false });
      }

      const user = authService.validate(token);
      if (user !== null) {
        return res.status(200).json({
          success: true,
          user,
        });
      }

      return res.status(200).json({ success: false });
    } catch (err) {
      return next(err);
    }
  },
};
