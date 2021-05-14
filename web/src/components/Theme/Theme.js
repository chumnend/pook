import blue from '@material-ui/core/colors/blue';
import green from '@material-ui/core/colors/green';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

export const theme = createMuiTheme({
  palette: {
    primary: {
      main: blue[500],
    },
    secondary: {
      main: green[500],
    },
  },
});

const ThemeProvider = ({ children }) => {
  return <MuiThemeProvider theme={theme}>{children}</MuiThemeProvider>;
};

ThemeProvider.propTypes = {
  children: PropTypes.array.isRequired,
};

export default ThemeProvider;
