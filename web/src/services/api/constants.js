// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
/* istanbul ignore next */
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

/** user api routes */
export const API_USER_REGISTER = apiPrefix + '/api/v1/register';
export const API_USER_LOGIN = apiPrefix + '/api/v1/login';
