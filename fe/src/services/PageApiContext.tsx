import {CategoryCreate, Configuration, PageApi, PageCore, PageCreate, PageUrlIdAndTitle} from "../openapi";
import {createContext, ReactNode, useCallback, useContext, useEffect, useState} from "react";
import {useConfig} from "./ConfigContext";
import {useAuth0} from "@auth0/auth0-react";

export interface PageApiData {
    pages: PageUrlIdAndTitle[]
    api: PageApi
    refreshPages: () => Promise<void>;
    add: (pageCreate: PageCreate) => Promise<string|undefined>;
    remove: (urlId: string) => Promise<void>;
}

const PageApiContext = createContext<PageApiData | undefined>(undefined);

interface PageApiProviderProps {
    children: ReactNode;
}

export const PageApiProvider: React.FC<PageApiProviderProps> = ({children,}) => {
    const config = useConfig()!;
    const {getAccessTokenSilently, isAuthenticated} = useAuth0();
    const [pages, setPages] = useState<PageUrlIdAndTitle[]>([])
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

    const refreshPages = useCallback(async () => {
        try {
            const response = await api.pageList();
            setPages(response.data.data);
        } catch (err) {
            console.error(err);
        }
    }, [api]);

    const add = useCallback(
        async (pageCreate: PageCreate) => {
            try {
                const response = await api.pageCreate(pageCreate);

                if (response.status === 201) {
                    await refreshPages()
                    const location = response.headers['location']
                    if (location) {
                        return location
                    }

                    return Promise.reject(new Error("location not found"))
                }

                return Promise.reject(new Error(`request failed with status: ${response.status}`))
            } catch (err) {
                console.log(err);
            }
        },
        [api],
    );

    const remove = useCallback(
        async (urlId: string) => {
            try {
                const response = await api.pageDelete(urlId);

                if (response.status === 200) {
                    await refreshPages()
                }
            } catch (err) {
                console.log(err);
            }
        },
        [api],
    );

    return (
        <PageApiContext.Provider value={{pages, api, refreshPages, add, remove}}>
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