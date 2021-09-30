import { Link, Button } from "@mui/material";

function Nav() {
  return (
    <>
      <Link href="/">
        <Button>메인 화면</Button>
      </Link>
      <Link href="/login">
        <Button>로그인 화면</Button>
      </Link>
      <Link href="/user">
        <Button>유저 화면</Button>
      </Link>
      <Link href="/signin">
        <Button>회원 가입 화면</Button>
      </Link>
    </>
  );
}

export default Nav;