import {createContext, ReactNode, useCallback, useContext, useEffect, useState} from "react";
import {Category, CategoryApi, CategoryCreate, Configuration} from "../openapi";
import {useConfig} from "./ConfigContext";
import {useAuth0} from "@auth0/auth0-react";


export interface CategoryApiData {
    categories: Category[],
    api: CategoryApi | undefined,
    withAuth: () => void
    refreshCategories: () => Promise<void>
    add: (newCategory: CategoryCreate) => Promise<void>
    remove: (urlId: string) => Promise<void>
}

const CategoryApiContext = createContext<CategoryApiData | undefined>(undefined);

interface CategoryApiProviderProps {
    children: ReactNode
}

export const CategoryApiProvider: React.FC<CategoryApiProviderProps> = ({children}) => {
    const config = useConfig()!
    const { getAccessTokenSilently, isAuthenticated } = useAuth0();
    const [categories, setCategories] = useState<Category[]>([]);
    const [api, setApi] = useState<CategoryApi | undefined>(
        (
            isAuthenticated ?
                new CategoryApi(new Configuration({
                    basePath: config.API_BASE,
                    accessToken: getAccessTokenSilently({authorizationParams:{audience:config.AUDIENCE}})
                })) :
                new CategoryApi(new Configuration({
                    basePath: config.API_BASE
                }))
        )
    );

    useEffect(() => {
        if (isAuthenticated) {
            setApi(new CategoryApi(new Configuration({
                basePath: config.API_BASE,
                accessToken: getAccessTokenSilently({authorizationParams:{audience:config.AUDIENCE}})
            })));
        }
    }, [isAuthenticated])


    const withAuth = useCallback(() => {
        setApi(new CategoryApi(new Configuration({
            basePath: config.API_BASE,
            accessToken: getAccessTokenSilently({authorizationParams:{audience:config.AUDIENCE}})
        })))
    }, [])

    const refreshCategories = useCallback(async () => {
        try {
            const response = await api!.categoryList()
            setCategories(response.data.data)
        }
        catch(err) {
            console.error(err)
        }
    }, [api]);

    const add = useCallback(async (newCategory: CategoryCreate) => {
        try {
            console.log(api);
            const response = await api!.categoryCreate(newCategory);

            if (response.status == 201) {
                setCategories((categories) => [...categories, response.data])
            }
        }
        catch(err) {
            console.log(err);
        }
    }, [api]);

    const remove = useCallback(async (urlId: string) => {
        try {
            const response = await api!.categoryDelete(urlId);

            if (response.status == 200) {
                setCategories((categories) => categories.filter((category) => category.urlId !== urlId))
            }
        }
        catch(err) {
            console.log(err);
        }
    }, [api]);


    return (
        <CategoryApiContext.Provider value={{ categories, api, withAuth, refreshCategories, add, remove }}>
            {children}
        </CategoryApiContext.Provider>
    )
};

export function useCategoryApiCtx() {
    const context = useContext(CategoryApiContext);
    if (!context) {
        throw new Error('useCategoryApiCtx should be used within a CategoryApiContext.Provider');
    }
    return context;
}