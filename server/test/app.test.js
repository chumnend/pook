'use strict';

const chai = require('chai');
const chaiHttp = require('chai-http');
const faker = require('faker');
const jwt = require('jsonwebtoken');
const app = require('../src/app');
const db = require('../src/models');
const config = require('../src/config');

const expect = chai.expect;
chai.use(chaiHttp);

describe('Application', () => {
  let user, token;
  
  before('setting up db', async () => {
    await db.User.deleteMany();

    user = {
      name: faker.name.findName(),
      email: faker.internet.email(),
      password: faker.internet.password(),
    }
    const createdUser = await db.User.create(user);
    user.id = createdUser.id;

    const payload = { id: user.id, name: user.name };
    token = jwt.sign(payload, config.secret);
  });

  after('cleaning up', async () => {
    await db.User.deleteMany();
  })

  describe('GET /status', () => {
    it('expects to get status OK', (done) => {
      chai
        .request(app)
        .get('/status')
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200)
          done();
        });
    });
  });

  describe('POST /api/users/register', () => {
    it('expects to register a new user', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123456',
        password2: '123456'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body).to.have.property('_id');
          expect(res.body).to.have.property('name', newUser.name);
          expect(res.body).to.have.property('email', newUser.email);
          done();
        });
    });
    it('expects to fail if missing name field', (done) => {
      const newUser = {
        email: faker.internet.email(),
        password: '123456',
        password2: '123456'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        });
    });
    it('expects to fail if missing email field', (done) => {
      const newUser = {
        name: faker.name.findName(),
        password: '123456',
        password2: '123456'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        });
    });
    it('expects to fail if missing password field', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password2: '123456'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    });
    it('expects to fail if missing password2 field', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123456',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    });
    it('expects to fail if password does not match password2', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123456',
        password2: '654321'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    });
    it('expects to fail if email is already taken', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: user.email,
        password: '123456',
        password2: '123456'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    });
    it('expects to fail if password is invalid (too short)', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123',
        password2: '123'
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    });
  });

  describe('POST /api/users/login', () => {
    it('expects to login a user', (done) => {
      const login = {
        email: user.email,
        password: user.password,
      }

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body).to.equal(token);
          done();
        })
    })
    it('expects to fail if missing email', (done) => {
      const login = {
        password: user.password,
      }

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    })
    it('expects to fail if missing password', (done) => {
      const login = {
        email: user.email,
      }

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    })
    it('expects to fail if user does not exist', (done) => {
      const login = {
        email: faker.internet.email(),
        password: user.password,
      }

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(404);
          expect(res.body).to.have.property('error');
          done();
        })
    })
    it('expects to fail if password in incorrect', (done) => {
      const login = {
        email: user.email,
        password: user.password + '1',
      }

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        })
    })
  });

  describe('GET /not-a-path', () => {
    it('expects to get 404 error', (done) => {
      chai
        .request(app)
        .get('/not-a-path')
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(404);
          expect(res.body).to.have.property('error');
          done();
        })
    })
  });
});
