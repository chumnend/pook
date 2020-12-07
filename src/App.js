import { Switch, Route } from 'react-router-dom';
import Landing from './containers/Landing';
import NotFound from './containers/NotFound';

const App = () => {
  return (
    <Switch>
      <Route exact path="/" component={Landing} />
      <Route component={NotFound} />
    </Switch>
  );
};

export default App;
