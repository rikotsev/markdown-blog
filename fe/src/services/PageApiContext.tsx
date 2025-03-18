import {Configuration, PageApi} from "../openapi";
import {createContext, ReactNode, useContext, useEffect, useState} from "react";
import {useConfig} from "./ConfigContext";
import {useAuth0} from "@auth0/auth0-react";

export interface PageApiData {
    api: PageApi
}

const PageApiContext = createContext<PageApiData | undefined>(undefined);

interface PageApiProviderProps {
    children: ReactNode;
}

export const PageApiProvider: React.FC<PageApiProviderProps> = ({children,}) => {
    const config = useConfig()!;
    const {getAccessTokenSilently, isAuthenticated} = useAuth0();
    const [api, setApi] = useState<PageApi>(
        isAuthenticated
            ? new PageApi(
                new Configuration({
                    basePath: config.API_BASE,
                    accessToken: getAccessTokenSilently({
                        authorizationParams: {audience: config.AUDIENCE},
                    }),
                }),
            )
            : new PageApi(
                new Configuration({
                    basePath: config.API_BASE,
                }),
            ),
    );

    useEffect(() => {
        if (isAuthenticated) {
            setApi(
                new PageApi(
                    new Configuration({
                        basePath: config.API_BASE,
                        accessToken: getAccessTokenSilently({
                            authorizationParams: {audience: config.AUDIENCE},
                        }),
                    }),
                ),
            );
        }
    }, [
        isAuthenticated,
        config.API_BASE,
        config.AUDIENCE,
        getAccessTokenSilently,
    ]);

    return (
        <PageApiContext.Provider value={{api}}>
            {children}
        </PageApiContext.Provider>
    );
};

export function usePageApiCtx() {
    const context = useContext(PageApiContext);
    if (!context) {
        throw new Error(
            "usePageApiCtx should be used within a PageApiContext.Provider",
        );
    }
    return context;
}