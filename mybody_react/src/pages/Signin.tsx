import { useState } from 'react';
import styled from 'styled-components';
import axios from 'axios';

const Signin = () => {
  const [Id, SetId] = useState("");
  const [Name, SetName] = useState("");
  const [Password, SetPassword] = useState("");

  const idHandler = (e: any) => {
    e.preventDefault();
    SetId(e.target.value);
  };

  const nameHandler = (e: any) => {
    e.preventDefault();
    SetName(e.target.value);
  };

  const passwordHandler = (e:any) => {
    e.preventDefault();
    SetPassword(e.target.value);
  };

  const submitHandler = (e:any) => {
    e.preventDefault();

    const idLength = Id.length;
    const nameLength = Name.length;
    const passwordLength = Password.length;

    if ( idLength === 0 ) {
      alert('ID를 입력해 주세요.');
      return;
    }

    if ( nameLength === 0 ) {
      alert('이름을 입력해 주세요.');
      return;
    }

    if ( passwordLength === 0 ) {
      alert('비밀번호를 입력해 주세요.');
      return;
    }

    if ( idLength > 32 ) {
      alert('ID가 길이 제한을 초과하였습니다.');
      return;
    }

    if ( nameLength > 32 ) {
      alert('이름이 길이 제한을 초과하였습니다.');
      return;
    }

    if ( passwordLength > 32 ) {
      alert('비밀번호가 길이 제한을 초과하였습니다.');
      return;
    }

    let body = {
      id: Id,
      name: Name,
      password: Password,
    };

    axios.post("/users", body)
      .then((res) => {
        console.log(res.data.result)
      })
      .catch((error) => {
        alert('error : ' + error);
      });
  }

  return (<div>
    <h2>회원 가입</h2>
      <form onSubmit={submitHandler}>
        ID : <Input name="id"
        value={Id}
        onChange={idHandler}
        width="380px"
        height="45px"
        placeholder="ID" />
        <br/>

        이름 : <Input name="name"
        value={Name}
        onChange={nameHandler}
        width="380px"
        height="45px"
        placeholder="이름" />
        <br/>

        비밀번호 : <Input name="password"
        value={Password}
        onChange={passwordHandler}
        width="380px"
        height="45px"
        placeholder="비밀번호"
        type="password" />
        <br/>

        <input type="submit" value="회원가입" />
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

export default Signin;