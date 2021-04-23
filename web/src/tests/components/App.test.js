// import { screen } from '@testing-library/react';
import App from '../../components/App';
import { customRender } from '../test-utils';

test('render <App />', () => {
  customRender(<App />);
});
