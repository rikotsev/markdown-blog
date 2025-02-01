import {Link, NavLink} from "react-router-dom";
import React from "react";
import ArticleApi from "../../services/ArticleApi";
import {useAuth0} from "@auth0/auth0-react";
import admin from "../../pages/admin/Admin";

const MenuBar: React.FC = () => {

    const categories = ArticleApi.getInstance().getCategories()
    const { loginWithRedirect, isAuthenticated, logout } = useAuth0();

    let oauthButton;
    let adminSection;
    if (isAuthenticated) {
        adminSection = <li>
            <NavLink to="/admin">Admin panel</NavLink>
        </li>
        oauthButton = <a href="#" onClick={() => logout()}>Log Out</a>
    }
    else {
        oauthButton = <a href="#" onClick={() => loginWithRedirect()}>Log In</a>
        adminSection = ""
    }

    return (
        <nav className="navbar">
            <ul>
                <li>
                    <NavLink to="/" className={({ isActive }) => (isActive ? 'active' : '')}>Home</NavLink>
                </li>
                <li>
                    <NavLink to="/about" className={({ isActive }) => (isActive ? 'active' : '')}>About</NavLink>
                </li>
                <li>
                    <NavLink to="/contact" className={({ isActive }) => (isActive ? 'active' : '')}>Contact</NavLink>
                </li>
                {categories.map((category) => (
                    <li>
                        <NavLink to={`/category/${category.prettyId}`} className={({ isActive }) => (isActive ? 'active' : '')}>{category.title}</NavLink>
                    </li>
                ))}
                {adminSection}
                <li>
                    {oauthButton}
                </li>
            </ul>
        </nav>
    )
}

export default MenuBar;