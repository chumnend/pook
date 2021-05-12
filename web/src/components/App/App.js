import { Route, Switch } from 'react-router-dom';

const App = () => {
  return (
    <Switch>
      <Route component={() => <h1>Hello World</h1>} />
    </Switch>
  );
};

export default App;
