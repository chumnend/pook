import React from 'react';
import { Switch, Route } from 'react-router-dom';

const Test = () => <p>Test</p>;
const TestHello = () => <p>TestHello</p>;

const Router = () => {
  return (
    <Switch>
      <Route exact path="/hello" component={TestHello} />
      <Route exact path="/" component={Test} />
    </Switch>
  )
}

export default Router;
