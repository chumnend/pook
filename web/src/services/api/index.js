import * as authHelpers from './auth';
import * as bookHelpers from './book';

const apiHelpers = {
  ...authHelpers,
  ...bookHelpers,
};

export default apiHelpers;
