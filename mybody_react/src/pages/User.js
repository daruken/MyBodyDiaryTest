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
  }

  getUsers = async() => {
    const response = await axios.get('/users');
    const user = response.data;

    this.setState({user});
  };

  render() {
    return <div>
      <h2>User Page</h2>
      
      <p>User list</p>
      {
      this.state.user.map((user) => {
        return <ul key={user.id}>
          <p>ID : {user.id}</p>
          <p>E-Mail : {user.email}</p>
        </ul>
      })
      }
    </div>
  }
}

export default User;