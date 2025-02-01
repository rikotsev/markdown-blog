import React, {useEffect, useState} from 'react';
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
import {Auth0Provider} from '@auth0/auth0-react';
import {Config, ConfigContext} from './services/ConfigContext';

function App() {

    const [config, setConfig] = useState<Config | null>(null);
    const [error, setError] = useState("");

    useEffect(() => {
        async function loadConfig() {
            try {
                const response = await fetch("/config.json");
                if (!response.ok) {
                    throw new Error(`Failed to fetch config.json: ${response.status}`);
                }
                const data: Config = await response.json();
                setConfig(data);
            } catch (err) {
                // @ts-ignore
                setError(err.message);
            }
        }

        loadConfig();
    }, []);

    if (error) {
        return <div>Error loading config: {error}</div>
    }

    if (!config) {
        return <div>Loading config...</div>
    }

    return (
        <ConfigContext.Provider value={config}>
            <Auth0Provider
                domain={config.AUTH_0_DOMAIN}
                clientId={config.AUTH_0_CLIENT_ID}
                authorizationParams={{
                    redirect_uri: window.location.origin,
                    audience: config.AUDIENCE,
                    scope: ""
                }}
            >
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
                    </Routes>
                </Router>
            </Auth0Provider>
        </ConfigContext.Provider>
    );
}

export default App;
