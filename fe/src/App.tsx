import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import MenuBar from "./components/menu/MenuBar";
import Home from "./pages/home/Home";
import About from "./pages/about/About";
import Contact from "./pages/contact/Contact";
import Category from "./pages/category/Category";
import Article from "./pages/article/Article";
import Admin from "./pages/admin/Admin";
import DesignArticle from "./pages/article/DesignArticle";

function App() {
    return (
        <Router>
            <MenuBar/>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/about" element={<About/>}/>
                <Route path="/contact" element={<Contact/>}/>
                <Route path="/category/:category" element={<Category/>}/>
                <Route path="/category/:category/articles/:id" element={<Article/>}/>
                <Route path="/admin" element={<Admin/>}/>
                <Route path="/article/create" element={<DesignArticle />}/>
            </Routes>
        </Router>
    );
}

export default App;
