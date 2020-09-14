'use strict';

const mongoose = require('mongoose');
const bcrypt = require('bcryptjs');

const userSchema = new mongoose.Schema({
  name: { 
    type: String, 
    required: true 
  },
  email: { 
    type: String, 
    required: true, 
    unique: true 
  },
  password: { 
    type: String, 
    required: true 
  },
  created: { 
    type: Date, 
    default: Date.now 
  },
});

userSchema.pre('save', async function(next) {
  try {
    if (!this.isModified('password')) {
      return next();
    }

    let hashedPwd = await bcrypt.hash(this.password, 10);
    this.password = hashedPwd;
  } catch(err) {
    return next(err);
  }
});

userSchema.methods.comparePassword = async function (password, next) {
  try {
    let isMatch = await bcrypt.compare(password, this.password);
    return isMatch;
  } catch (err) {
    return next(err);
  }
};

module.exports = mongoose.model('User', userSchema);
