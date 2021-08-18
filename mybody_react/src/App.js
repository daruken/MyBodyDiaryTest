import React from 'react';
import { BrowserRouter, Route, Switch } from "react-router-dom";
import './App.css';
import { Home, User, Signin, Nav } from './pages';

function App() {
  return (
    <BrowserRouter>
      <Nav />
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/user" component={User} />
        <Route exact path="/signin" component={Signin} />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
