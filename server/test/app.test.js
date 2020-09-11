'use strict';

const chai = require('chai');
const chaiHttp = require('chai-http');
const app = require('../src/app');

const expect = chai.expect;
chai.use(chaiHttp);

describe('Application', () => {
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
