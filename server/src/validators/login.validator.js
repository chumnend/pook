'use strict';

const validator = require('validator');
const isEmpty = require('is-empty');

module.exports = function validateLogin(data) {
  const errors = {};
  // check for empty strings
  data.email = !isEmpty(data.email) ? data.email : '';
  data.password = !isEmpty(data.password) ? data.password : '';

  // validate each field
  if (validator.isEmpty(data.email)) {
    errors.email = 'Email field required';
  } else if (!validator.isEmail(data.email)) {
    errors.email = 'Email field not valid';
  }

  if (validator.isEmpty(data.password)) {
    errors.password = 'Password field is required';
  }

  return {
    errors,
    isValid: isEmpty(errors),
  };
};
