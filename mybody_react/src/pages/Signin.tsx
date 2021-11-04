import { useState } from 'react';
import Input from '@mui/material/Input';
import Button from '@mui/material/Button';
import axios from 'axios';
import CSS from 'csstype';
import PureModal from 'react-pure-modal';
import 'react-pure-modal/dist/react-pure-modal.min.css';

const Signin = () => {
  const [modal, setModal] = useState(false);

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
        if (res.data.result === 0) {
          setModal(true);
        }
      })
      .catch((error) => {
        alert('error : ' + error);
      });
  }

  const buttonStyle: CSS.Properties = {
    margin: '30px',
    width: '100px'
  }

  return (<div>
    <h2>회원 가입</h2>
      <form onSubmit={submitHandler}>
        <Input name="id"
          value={Id}
          onChange={idHandler}
          placeholder="ID" />
          <br/>

        <Input name="name"
          value={Name}
          onChange={nameHandler}
          placeholder="이름" />
          <br/>

        <Input name="password"
          value={Password}
          onChange={passwordHandler}
          placeholder="비밀번호"
          type="password" />
          <br/>

        <Button variant="contained" style={buttonStyle} type="submit">
          회원 가입</Button>
      </form>

      <PureModal
        header="회원 가입"
        footer={
          <div>
            <button onClick={() => setModal(false)} >Ok</button>
          </div>
        }
        isOpen={modal}
        onClose={() => {
          setModal(false);
          return true;
        }}
      >
        <p>회원 가입이 완료되었습니다.</p>
      </PureModal>

    </div>
  )
}

export default Signin;