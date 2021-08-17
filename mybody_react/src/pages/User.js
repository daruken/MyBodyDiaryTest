import React, {Component} from 'react';
import axios from 'axios';

class User extends Component {
  constructor(props) {
    super(props);
  };

  state = {
    user : []
  };

  componentDidMount() {
    console.log('in componentDidMount');
    this.getUsers();
    console.log('after user : ' + JSON.stringify(this.state.user));
  }

  getUsers = async() => {
    const response = await axios.get('/users');
    const user = response.data;
    console.log('in user : ' + JSON.stringify(user));

    this.setState({user});
  };

  render() {
    return <div>
      <h2>User Page</h2>
      
      <p>User list</p>
      {
      this.state.user.map((user) => {
        return [
          <p>ID : { user.id }</p>,
          <p>Email : { user.email }</p>
        ];
      })
      }
    </div>
  }
}

export default User;