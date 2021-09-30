import { useState, useEffect } from 'react';
import Table from '@mui/material/Table';
import TableContainer from '@mui/material/TableContainer';
import TableBody from '@mui/material/TableBody';
import TableHead from '@mui/material/TableHead';
import { TableRow, TableCell } from '@mui/material';
import Paper from '@mui/material/Paper';
import axios from 'axios';

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
      <TableContainer component={Paper} sx={{ width: 650 }}>
        <Table sx={{ width: 650 }} aria-label="simple table">
          <TableHead>
            <tr>
              {columns.map((column) => (
                <th key={column}>{column}</th>
              ))}
            </tr>
          </TableHead>
          <TableBody>
            {users && users.map(({ id, name }) => (
              <TableRow key={id + name}>
                <TableCell align="right">{id}</TableCell>
                <TableCell align="right">{name}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

    </div>
  );
}

export default User;
