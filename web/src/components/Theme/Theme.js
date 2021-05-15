import CssBaseline from '@material-ui/core/CssBaseline';
import {
  createMuiTheme,
  MuiThemeProvider,
  responsiveFontSizes,
} from '@material-ui/core/styles';
import PropTypes from 'prop-types';

// https://material.io/resources/color/#!/?view.left=0&view.right=0&primary.color=6D4C41&secondary.color=BCAAA4&secondary.text.color=000000&primary.text.color=ffffff
const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#6d4c41',
      light: '#9c786c',
      dark: '#40241a',
    },
    secondary: {
      main: '#bcaaa4',
      light: '#efdcd5',
      dark: '#8c7b75',
    },
    text: {
      light: '#ffffff',
      dark: '#000000',
    },
  },
});

const ThemeProvider = ({ children }) => {
  return (
    <MuiThemeProvider theme={responsiveFontSizes(theme)}>
      <CssBaseline />
      {children}
    </MuiThemeProvider>
  );
};

ThemeProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export default ThemeProvider;
