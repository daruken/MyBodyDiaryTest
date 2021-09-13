import { BrowserRouter, Route, Switch } from "react-router-dom";
import './App.css';
import { Home, Signin, Nav, UserList } from './pages';

function App() {
  return (
    <BrowserRouter>
      <Nav />
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/user" component={UserList} />
        <Route exact path="/signin" component={Signin} />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
