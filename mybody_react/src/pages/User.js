import React, { Component } from 'react';
import axios from 'axios';

class User extends Component {
  state = {
    user : []
  };

  componentDidMount() {
    this.getUsers();
  }

  getUsers = async() => {
    const response = await axios.get('/users');
    const user = response.data;

    this.setState({user});
  };

  render() {
    return <div>
      <h2>사용자 정보</h2>
      
      <p>사용자 목록</p>
      {
      this.state.user.map((user) => {
        return <ul key={user.id}>
          <p>ID : {user.id}</p>
          <p>이름 : {user.name}</p>
        </ul>
      })
      }
    </div>
  }
}

export default User;