import { Link } from "react-router-dom";

function Nav() {
  return (
    <>
      <Link to="/">
        <button>메인 화면으로</button>
      </Link>
      <Link to="/user">
        <button>유저 화면으로</button>
      </Link>
      <Link to="/signin">
        <button>회원 가입 화면으로</button>
      </Link>
    </>
  );
}

export default Nav;