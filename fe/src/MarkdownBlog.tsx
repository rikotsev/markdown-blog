import React, {useEffect, useState} from "react";
import MenuBar from "./components/menu/MenuBar";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Category from "./pages/category/Category";
import Article from "./pages/article/Article";
import Admin from "./pages/admin/Admin";
import DesignArticle from "./pages/article/DesignArticle";
import {useCategoryApiCtx} from "./services/CategoryApiContext";
import {Page} from "./pages/page/Page";
import DesignPage from "./pages/page/DesignPage";
import {usePageApiCtx} from "./services/PageApiContext";

export const MarkdownBlog: React.FC = () => {
    //Everything you need to render a proper UI
    const {refreshCategories} = useCategoryApiCtx();
    const {refreshPages} = usePageApiCtx();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        Promise.all([refreshCategories(), refreshPages()]).then(() => {
            setLoading(false);
        }).catch((err) => {
            console.error('failed to load initial data: ', err)
        })
    }, [refreshCategories, refreshPages]);

    if (loading) {
        return (
            <div style={{textAlign: "center", marginTop: "50px"}}>
                <div className="spinner"/>
                <p>Loading...</p>
            </div>
        );
    }

    return (
        <Router>
            <MenuBar/>
            <Routes>
                <Route path="/" element={<Page/>}/>
                <Route path="/page/:id" element={<Page/>}/>
                <Route path="/category/:category" element={<Category/>}/>
                <Route path="/category/:category/articles/:id" element={<Article/>}/>
                <Route path="/admin" element={<Admin/>}/>
                <Route path="/article/create" element={<DesignArticle/>}/>
                <Route path="/article/:id" element={<DesignArticle/>}/>
                <Route path="/page/create" element={<DesignPage/>}/>
                <Route path="/page/design/:id" element={<DesignPage/>}/>
            </Routes>
        </Router>
    );
};
