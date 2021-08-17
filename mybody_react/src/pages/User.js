import React from 'react';
import axios from 'axios';

async function getAllUsers() {
  const response = await axios.get('/users');

  return response.data;
}

const User = () => {
  return (
    <div>
      Hi Users!

    </div>
  )
}

export default User;