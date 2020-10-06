'use strict';

class HttpError extends Error {
  constructor(status, message, extra = {}) {
    super(message);

    this.status = status;
    this.message = message;
    this.extra = extra;
  }
}

module.exports = HttpError;
