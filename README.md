# Pook: An app for creating storybooks
Pook is a storybook creator application using React and Go. The idea if this app is to allow users to create fun novels in thier browser and share them with others.

## Demo
TBD

## Development Setup
### Prerequisites
- Go v1.15
- npm
- Postgresql

### Configuration
1) Clone this repo using `git clone https://github.com/chumnend/pook.git`

2) Create a postgresql database locally or online. The database will be connected through using a connection string. It should have the form of 
`postgresql://<username>:<password>@<address>/<dbname>`

3) Run `cp env.example .env` and then open the `.env` and fill out the following fields,
```
PORT= # the port the app will run on
SECRET_KEY= # string used for hashing
DATABASE_URL= # database string used to connect to postgresql database
DATABASE_TEST_URL= # database string to database used for integration tests
```

4) This step is only needed if you plan to work on the React app. Go into the web folder ie. `cd web/` and copy run `cp env.example .env` and then open the `.env` and fill out the following fields,

```
NODE_ENV= # production, dev or test
PORT= # the port the React app can run on, should be different than the previous step
BROWSER=none # by default stops the browser from opening when running thr React app

REACT_APP_API_PREFIX= # this should point to the address the Go app runs on
```

5) Now the apps are ready to run. Go back to the root folder and run `make` this will build the React and Go assets and start the app on the given port. You can build assets on thier own using the `make build` command and just serve currently built assets using `make serve`.

## Deployment
Not currently deployed.
