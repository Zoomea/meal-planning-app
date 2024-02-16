import { useEffect } from "react";
import { Outlet, Link, useLocation } from "react-router-dom";
import siteLogo from './assets/siteLogo.png';
import recipeLogo from './assets/recipe.png';
import calendarLogo from './assets/calendar.png';
import shoppingLogo from './assets/shopping.png';
import "./App.css";

const Layout = () => {
    const location = useLocation();

    useEffect(() => {
        // Change the body background color based on the route
        document.body.style.backgroundColor = location.pathname === '/calendar' ? 'white' : '#EAEFC5';
    
        // Clean up the effect when the component unmounts
        return () => {
          document.body.style.backgroundColor = '#EAEFC5'; // Reset to default background color or remove this line if not needed
        };
      }, [location.pathname]);

    return (
        <>
            <div className="navigation-bar bg-green-500 flex">
                <nav className="w-3/4 flex justify-between mx-auto items-center">
                    <img id="site-logo" src={siteLogo} alt="site logo" />
                    <ul className="flex">
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
            <div className="w-full mt-48 main-content">
                <Outlet />
            </div>
        </>
    )
};

export default Layout;