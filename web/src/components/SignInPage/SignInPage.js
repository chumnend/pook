import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import Checkbox from '@material-ui/core/Checkbox';
import Container from '@material-ui/core/Container';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Grid from '@material-ui/core/Grid';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { useState } from 'react';
import { Link as RouterLink, useHistory } from 'react-router-dom';

import { useAuth } from '../AuthProvider';
import { HOME_ROUTE, NOT_FOUND_ROUTE, REGISTER_ROUTE } from '../Router';

const useStyles = makeStyles((theme) => ({
  authCard: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  gridItem: {
    textAlign: 'center',
    padding: theme.spacing(0.5),
  },
}));

const SignInPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);

  const auth = useAuth();
  const history = useHistory();
  const classes = useStyles();

  const validate = () => {
    return email.length > 0 && password.length > 0;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const success = await auth.login(email, password);
    if (success) {
      history.push(HOME_ROUTE);
    } else {
      alert(auth.error);
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.authCard}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Grid className={classes.gridItem} item xs={12} md={7}>
          <Link variant="body2" component={RouterLink} to={HOME_ROUTE}>
            Back to Home
          </Link>
        </Grid>
        <Typography component="h1" variant="h5">
          Welcome Back!
        </Typography>
        <form className={classes.form} noValidate onSubmit={handleSubmit}>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="password"
            name="password"
            label="Password"
            type="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me"
            checked={rememberMe}
            onChange={() => setRememberMe(!rememberMe)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            disabled={!validate()}
          >
            Sign In
          </Button>
          <Grid container>
            <Grid className={classes.gridItem} item xs={12} md={5}>
              <Link variant="body2" component={RouterLink} to={NOT_FOUND_ROUTE}>
                Forgot password?
              </Link>
            </Grid>
            <Grid className={classes.gridItem} item xs={12} md={7}>
              <Link variant="body2" component={RouterLink} to={REGISTER_ROUTE}>
                {"Don't have an account?"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
    </Container>
  );
};

export default SignInPage;
