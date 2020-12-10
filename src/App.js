import { Switch, Route } from 'react-router-dom';
import Landing from './containers/Landing';
import Auth from './containers/Auth';
import NotFound from './containers/NotFound';

const App = () => {
  return (
    <Switch>
      <Route exact path="/register" render={(props) => <Auth {...props} />} />
      <Route
        exact
        path="/login"
        component={(props) => <Auth {...props} login />}
      />
      <Route exact path="/" component={Landing} />
      <Route component={NotFound} />
    </Switch>
  );
};

export default App;
