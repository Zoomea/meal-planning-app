import { Outlet, Link } from "react-router-dom";
import siteLogo from './assets/siteLogo.png';
import recipeLogo from './assets/recipe.png';
import calendarLogo from './assets/calendar.png';
import shoppingLogo from './assets/shopping.png';
import "./App.css";

const Layout = () => {
    return (
        <>
            <div className="navigation-bar">
                <img id="site-logo" src={siteLogo} alt="site logo" />
                <nav>
                    <ul>
                        <li>
                            <a href="/">
                                <img src={recipeLogo} alt="recipe logo" />
                                <Link to="/">Recipe book</Link>
                            </a>
                        </li>
                        <li>
                            <a href="/calendar">
                                <img src={calendarLogo} alt="calendar logo" />
                                <Link to="/calendar">Planning</Link>
                            </a>
                        </li>
                        <li>
                            <a href="/shoppingList">
                                <img src={shoppingLogo} alt="shopping logo" />
                                <Link to="/shoppingList">Shopping list</Link>
                            </a>
                        </li>
                    </ul>
                </nav>
            </div>
            <div className="main-content">
                <Outlet />
            </div>
        </>
    )
};

export default Layout;