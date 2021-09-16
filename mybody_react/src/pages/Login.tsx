import { useState } from 'react';
import styled from 'styled-components';
import axios from 'axios';

const Login = () => {
  const [Id, SetId] = useState("");
  const [Password, SetPassword] = useState("");

  const idHandler = (e: any) => {
    e.preventDefault();
    SetId(e.target.value);
  };

  const passwordHandler = (e:any) => {
    e.preventDefault();
    SetPassword(e.target.value);
  };

  const submitHandler = (e:any) => {
    e.preventDefault();

    const idLength = Id.length;
    const passwordLength = Password.length;

    if ( idLength === 0 ) {
      alert('ID를 입력해 주세요.');
      return;
    }

    if ( passwordLength === 0 ) {
      alert('비밀번호를 입력해 주세요.');
      return;
    }

    let body = {
      id: Id,
      password: Password,
    };

    axios.post("/login", body)
      .then((res) => {
        console.log(res.data.result)
      })
      .catch((error) => {
        alert('error : ' + error);
      });
  }

  return (<div>
    <h2>로그인</h2>
      <form onSubmit={submitHandler}>
        ID : <Input name="id"
        value={Id}
        onChange={idHandler}
        width="380px"
        height="45px"
        placeholder="ID" />
        <br/>

        비밀번호 : <Input name="password"
        value={Password}
        onChange={passwordHandler}
        width="380px"
        height="45px"
        placeholder="비밀번호"
        type="password" />
        <br/>

        <input type="submit" value="Login" />
      </form>
    </div>
  )
}

const Input = styled.input`
    border: 1px solid gray;
    margin: 2px;
    outline: none;
    border-radius: 0px;
    line-height: 2.0rem;
    font-size: 1.2rem;
    padding-left: 0.5rem;
    padding-right: 0.5rem;
    ::placeholder {
        color: gray;
    }
`;

export default Login;