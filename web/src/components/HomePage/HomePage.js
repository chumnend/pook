import Container from '@material-ui/core/Container';

import Header from '../Header';

const HomePage = () => {
  return (
    <Container maxWidth="sm">
      <Header isAuth />
      <div>
        <input type="text" placeholder="filter" />
      </div>
      <div>Boards go here</div>
    </Container>
  );
};

export default HomePage;
