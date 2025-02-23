import { NavLink } from "react-router-dom";
import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { useCategoryApiCtx } from "../../services/CategoryApiContext";

const MenuBar: React.FC = () => {
  const { categories } = useCategoryApiCtx();
  const { loginWithRedirect, isAuthenticated, logout } = useAuth0();

  let oauthButton;
  let adminSection;
  if (isAuthenticated) {
    adminSection = (
      <li key={"admin"}>
        <NavLink to="/admin">Admin panel</NavLink>
      </li>
    );
    oauthButton = <button onClick={() => logout()}>Log Out</button>;
  } else {
    oauthButton = <button onClick={() => loginWithRedirect()}>Log In</button>;
    adminSection = "";
  }

  return (
    <nav className="navbar">
      <ul>
        <li key={"home"}>
          <NavLink
            to="/"
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Home
          </NavLink>
        </li>
        <li key={"about"}>
          <NavLink
            key={"about"}
            to="/about"
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            About
          </NavLink>
        </li>
        <li key={"contact"}>
          <NavLink
            key={"contact"}
            to="/contact"
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Contact
          </NavLink>
        </li>
        {categories.map((category) => (
          <li key={category.id}>
            <NavLink
              to={`/category/${category.urlId}`}
              className={({ isActive }) => (isActive ? "active" : "")}
            >
              {category.name}
            </NavLink>
          </li>
        ))}
        {adminSection}
        <li key={"login"}>{oauthButton}</li>
      </ul>
    </nav>
  );
};

export default MenuBar;
