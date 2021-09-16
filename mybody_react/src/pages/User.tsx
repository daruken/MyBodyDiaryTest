import { useState, useEffect } from 'react';
import axios from 'axios';
import '../css/style.css';

interface IUser {
  id: string;
  name: string;
}

const defaultProps:IUser[] = [];

function User() {
  const columns = ["ID", "NAME"];
  const [users, setUsers]: [IUser[], (posts: IUser[]) => void] = useState(defaultProps);

  useEffect(() => {
    axios.get<IUser[]>("/users")
    .then(response => {
        setUsers(response.data);
    });
  }, []);

  return (
    <div>
      <h2>사용자 정보</h2>
  
      <p>사용자 목록</p>
      <div>
        <table>
          <thead>
            <tr>
              {columns.map((column) => (
                <th key={column}>{column}</th>
              ))}
            </tr>
          </thead>
          <tbody>
            {users && users.map(({ id, name }) => (
              <tr key={id + name}>
                <td>{id}</td>
                <td>{name}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

    </div>
  );
}

export default User;
