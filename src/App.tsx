import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import './App.css'

function App() {
  // const [count, setCount] = useState(0)

  return (
    <>
      {/* <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount(count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p> */
    /* </div> */}
    <div>
      <span id="logo-img">logo</span>
      <span className="navigation-bar">
        <span>Recipe book</span>
        <span>Planning</span>
        <span>Shopping list</span>
        <li>
          <Link to="/new-route">New Route</Link>
        </li>
      </span>
    </div>
  </>
  )
}

export default App
