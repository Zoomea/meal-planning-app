import { Outlet, Link } from "react-router-dom";

const Layout = () => {
    return (
      <>
        <nav>
          <ul>
            <li>
              <Link to="/">Recipe book</Link>
            </li>
            <li>
              <Link to="/calendar">Planning</Link>
            </li>
            <li>
              <Link to="/shoppingList">Shopping list</Link>
            </li>
          </ul>
        </nav>
  
        <Outlet />
      </>
    )
  };
  
  export default Layout;