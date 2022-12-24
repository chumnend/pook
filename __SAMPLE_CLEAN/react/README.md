# Pook React App
This is the frontend application for the Pook application

## Development Setup

1) Run `cp env.example .env` and then open the newly created `.env` file and fill out the following fields,

```
NODE_ENV= # production, dev or test
PORT= # the port the React app can run on, should be different than the previous step
BROWSER=none # by default stops the browser from opening when running thr React app

REACT_APP_API_PREFIX= # this should point to the address the Go app runs on
```

2) Install dependencies using, `npm install`

3) Run tests using `npm test` to check component correctness or `npm run coverage` to get code test coverage as well

4) Use `npm run lint` to check code format and `npm run lint:fix` to fix code format
