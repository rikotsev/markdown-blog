import {Link, NavLink} from "react-router-dom";
import React from "react";

const MenuBar: React.FC = () => {
    return (
        <nav className="navbar">
            <ul>
                <li>
                    <NavLink to="/" className={({ isActive }) => (isActive ? 'active' : '')}>Home</NavLink>
                </li>
                <li>
                    <NavLink to="/category/tech">Tech</NavLink>
                </li>
                <li>
                    <NavLink to="/category/philosophy">Philosophy</NavLink>
                </li>
                <li>
                    <NavLink to="/category/travel">Travel</NavLink>
                </li>
                <li>
                    <NavLink to="/about" className={({ isActive }) => (isActive ? 'active' : '')}>About</NavLink>
                </li>
                <li>
                    <NavLink to="/contact" className={({ isActive }) => (isActive ? 'active' : '')}>Contact</NavLink>
                </li>
                <li>
                    <NavLink to="/admin">Admin panel</NavLink>
                </li>
            </ul>
        </nav>
    )
}

export default MenuBar;