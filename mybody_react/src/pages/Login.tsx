import { useState } from 'react';
import Input from '@mui/material/Input';
import Button from '@mui/material/Button';
import axios from 'axios';
import CSS from 'csstype';
import { useHistory } from 'react-router';
import Modal from 'react-modal';
import '../static/css/Style.css';


const Login = () => {
  const [modal, setModal] = useState(false);
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

    axios.post("/gateway/login", body)
      .then((res) => {
        if (res.data.result === 0) {
          setModal(true);

          //history.push("/");
        } else {
          alert('Login failed.');
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
    <h2>로그인</h2>
      <form onSubmit={submitHandler}>
        <Input name="id"
          value={Id}
          onChange={idHandler}
          placeholder="ID" />
          <br/>

        <Input name="password"
          value={Password}
          onChange={passwordHandler}
          placeholder="비밀번호"
          type="password" />
          <br/>

        <Button variant="contained" style={buttonStyle} type="submit">
          Login</Button>
      </form>

      <Modal
        isOpen={modal}
        className="modalContent"
        onRequestClose={() => setModal(false)}
        ariaHideApp={false}
        contentLabel="Login"
      >
        <p>로그인이 완료되었습니다.</p>

        <Button onClick={() => setModal(false)}>Ok</Button>
      </Modal>
    </div>
  )
}

export default Login;