import { NavLink } from "react-router-dom";
import React, {useEffect, useState} from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { useCategoryApiCtx } from "../../services/CategoryApiContext";
import {usePageApiCtx} from "../../services/PageApiContext";
import {PageUrlIdAndTitle} from "../../openapi";

const MenuBar: React.FC = () => {
  const { categories } = useCategoryApiCtx();
  const { pages, api } = usePageApiCtx();
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
        {pages.map((page) => (
            <li key={page.urlId}>
              <NavLink
                  to={`/page/${page.urlId}`}
                  className={({ isActive }) => (isActive ? "active" : "")}
              >
                {page.title}
              </NavLink>
            </li>
        ))}
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
