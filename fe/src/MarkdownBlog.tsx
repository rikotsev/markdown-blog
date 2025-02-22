import React, {useEffect, useState} from "react";
import MenuBar from "./components/menu/MenuBar";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Home from "./pages/home/Home";
import About from "./pages/about/About";
import Contact from "./pages/contact/Contact";
import Category from "./pages/category/Category";
import Article from "./pages/article/Article";
import Admin from "./pages/admin/Admin";
import DesignArticle from "./pages/article/DesignArticle";
import {useCategoryApiCtx} from "./services/CategoryApiContext";


export const MarkdownBlog: React.FC = () => {
    //Everything you need to render a proper UI
    const {refreshCategories} = useCategoryApiCtx();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        refreshCategories().then(() => {
            setLoading(false);
        });
    }, [])

    if (loading) {
        return (
            <div style={{ textAlign: "center", marginTop: "50px" }}>
                <div className="spinner" />
                <p>Loading...</p>
            </div>
        )
    }

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
                <Route path="/article/create" element={<DesignArticle/>}/>
                <Route path="/article/:id" element={<DesignArticle/>} />
            </Routes>
        </Router>
    );
}