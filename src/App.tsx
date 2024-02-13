import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './layout';
import Calendar from './calendar-view';
import Recipe from './recipe-list';
import ShoppingList from './shopping-list';
import NoPage from './no-page';
import './index.css';
import './App.css';


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
    {/* <div>
      <span id="logo-img">logo</span>
      <span className="navigation-bar">
        <span>Recipe book</span>
        <span>Planning</span>
        <span>Shopping list</span>
        <li>
          <Link to="/new-route">New Route</Link>
        </li>
      </span>
    </div> */}
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={< Recipe />} />
          <Route path="calendar" element={<Calendar />} />
          <Route path="shoppingList" element={<ShoppingList />} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </>
  )
}

export default App
