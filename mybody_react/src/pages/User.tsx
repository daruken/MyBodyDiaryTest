import React, { Component, useState, useEffect } from 'react';
import axios from 'axios';

interface IUser {
  id: string;
  name: string;
}

const defaultProps:IUser[] = [];

function User() {
  const [users, setUsers]: [IUser[], (posts: IUser[]) => void] = useState(defaultProps);
  
  useEffect(() => {
    axios.get<IUser[]>("/users")
    .then(response => {
        console.log(response.data);
        setUsers(response.data);
    });
  }, []);

  return (
    <div>
    <h2>사용자 정보</h2>
    
    <p>사용자 목록</p>
    {
      users.map(user => (
        <ul key={user.id}>
          <p>ID : {user.id}</p>
          <p>이름 : {user.name}</p>
        </ul>
      ))
    }
    </div>
  );
}

export default User;
