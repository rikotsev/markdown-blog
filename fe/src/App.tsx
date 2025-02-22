import React, {useEffect, useState} from 'react';
import './App.css';
import {Auth0Provider} from '@auth0/auth0-react';
import {Config, ConfigContext} from './services/ConfigContext';
import {CategoryApiProvider} from "./services/CategoryApiContext";
import {ArticleApiProvider} from "./services/ArticleApiContext";
import {MarkdownBlog} from "./MarkdownBlog";

function App() {

    const [config, setConfig] = useState<Config | null>(null);

    useEffect(() => {
        fetch("/config.json")
            .then((response) => {
                if (!response.ok) {
                    console.error("failed to fetch config.json", response)
                    return
                }

                response.json()
                    .then((data) => {
                        setConfig(data)
                    })
                    .catch((err) => console.log(err));
            })
            .catch((err) => console.log(err))
    }, []);

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
                <CategoryApiProvider>
                    <ArticleApiProvider>
                        <MarkdownBlog/>
                    </ArticleApiProvider>
                </CategoryApiProvider>
            </Auth0Provider>
        </ConfigContext.Provider>
    );
}

export default App;
