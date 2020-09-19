import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import styled from 'styled-components';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import Landing from './views/Landing';
import Register from './views/Register';
import Login from './views/Login';
import NotFound from './views/NotFound';

const StyledApp = styled.div`
  width: 100vw;
  height: 100vh;
  display: grid;
  grid-template-areas: 'header' 'main' 'footer';
  grid-template-columns: auto;
  grid-template-rows: 5rem 1fr 3rem;
`;

const StyledHeader = styled.header`
  grid-area: header;
`;

const StyledMain = styled.main`
  grid-area: main;
`;

const StyledFooter = styled.footer`
  grid-area: footer;
`;

function App() {
  return (
    <BrowserRouter>
      <StyledApp>
        <StyledHeader>
          <Navbar />
        </StyledHeader>
        <StyledMain>
          <Switch>
            <Route exact path="/register" component={Register} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/" component={Landing} />
            <Route component={NotFound} />
          </Switch>
        </StyledMain>
        <StyledFooter>
          <Footer />
        </StyledFooter>
      </StyledApp>
    </BrowserRouter>
  );
}

export default App;
