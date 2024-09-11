import {Link, NavLink} from "react-router-dom";
import React from "react";
import ArticleApi from "../../services/ArticleApi";

const MenuBar: React.FC = () => {

    const categories = ArticleApi.getInstance().getCategories()

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
                <li>
                    <NavLink to="/admin">Admin panel</NavLink>
                </li>
            </ul>
        </nav>
    )
}

export default MenuBar;