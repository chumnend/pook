'use strict';

const mongoose = require('mongoose');

const bookingSchema = mongoose.Schema({
  startDate: {
    type: Date,
    required: true,
  },
  endDate: {
    type: Date,
    required: true,
  },
  checkIn: {
    type: String,
    required: true,
  },
  checkOut: {
    type: String,
    required: true,
  },
  occupants: {
    type: Number,
    required: true,
  },
  totalCost: {
    type: Number,
    required: true,
  },
  guest: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'User',
  },
  room: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'Room',
  },
});

module.exports = mongoose.model('Booking', bookingSchema);
