import { useState } from 'react';

const Home = () => {
  const [count, setCount] = useState(0);

  return (
    <div>
      <h2>Home Page</h2>

      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  )
}

export default Home;