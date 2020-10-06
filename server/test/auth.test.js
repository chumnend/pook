'use strict';

const chai = require('chai');
const chaiHttp = require('chai-http');
const faker = require('faker');
const app = require('../src/app');
const db = require('../src/models');

const expect = chai.expect;
chai.use(chaiHttp);

describe('Auth Routes', () => {
  let user;
  
  before('setting up db', async () => {
    await db.User.deleteMany();

    user = {
      name: faker.name.findName(),
      email: faker.internet.email(),
      password: faker.internet.password(),
    }
    const createdUser = await db.User.create(user);
    user.id = createdUser.id;
  });

  after('cleaning up', async () => {
    await db.User.deleteMany();
  })

  describe('POST /api/users/register', () => {
    it('expects to register a new user', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123456',
        password2: '123456',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(200);
          expect(res.body).to.have.property('token');
          done();
        });
    });
    it('expects to fail if missing name field', (done) => {
      const newUser = {
        email: faker.internet.email(),
        password: '123456',
        password2: '123456',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('name');
          done();
        });
    });
    it('expects to fail if missing email field', (done) => {
      const newUser = {
        name: faker.name.findName(),
        password: '123456',
        password2: '123456',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('email');
          done();
        });
    });
    it('expects to fail if missing password field', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password2: '123456',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('password');
          done();
        });
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
          expect(res.body.error.extra).to.have.property('password2');
          done();
        });
    });
    it('expects to fail if password does not match password2', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123456',
        password2: '654321',
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
    it('expects to fail if email is already taken', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: user.email,
        password: '123456',
        password2: '123456',
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
    it('expects to fail if password is invalid (too short)', (done) => {
      const newUser = {
        name: faker.name.findName(),
        email: faker.internet.email(),
        password: '123',
        password2: '123',
      };

      chai
        .request(app)
        .post('/api/users/register')
        .send(newUser)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('password');
          done();
        });
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
          expect(res.body).to.have.property('token');
          done();
        });
    });
    it('expects to fail if missing email', (done) => {
      const login = {
        password: user.password,
      };

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('email');
          done();
        });
    });
    it('expects to fail if missing password', (done) => {
      const login = {
        email: user.email,
      };

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          expect(res.body.error.extra).to.have.property('password');
          done();
        });
    });
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
        });
    });
    it('expects to fail if password in incorrect', (done) => {
      const login = {
        email: user.email,
        password: user.password + '1',
      };

      chai
        .request(app)
        .post('/api/users/login')
        .send(login)
        .end((err, res) => {
          expect(err).to.be.null;
          expect(res).to.have.status(400);
          expect(res.body).to.have.property('error');
          done();
        });
    });
  });
});
