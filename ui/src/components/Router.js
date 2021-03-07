import React from 'react';
import { Switch, Route } from 'react-router-dom';

const Home = () => <h1>Home</h1>
const Page1 = () => <h1>Page 1</h1>;
const Page2 = () => <h1>Page 2</h1>;
const NotFound = () => <h1>NotFound</h1>;

const Router = () => {
  return (
    <Switch>
      <Route exact path="/1" component={Page1} />
      <Route exact path="/2" component={Page2} />
      <Route exact path ="/" component={Home} />
      <Route component={NotFound} />
    </Switch>
  )
}

export default Router;
