import { ArticleApi, Configuration } from "../openapi";
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import { useConfig } from "./ConfigContext";
import { useAuth0 } from "@auth0/auth0-react";

export interface ArticleApiData {
  api: ArticleApi;
}

const ArticleApiContext = createContext<ArticleApiData | undefined>(undefined);

interface ArticleApiProviderProps {
  children: ReactNode;
}

export const ArticleApiProvider: React.FC<ArticleApiProviderProps> = ({
  children,
}) => {
  const config = useConfig()!;
  const { getAccessTokenSilently, isAuthenticated } = useAuth0();
  const [api, setApi] = useState<ArticleApi>(
    isAuthenticated
      ? new ArticleApi(
          new Configuration({
            basePath: config.API_BASE,
            accessToken: getAccessTokenSilently({
              authorizationParams: { audience: config.AUDIENCE },
            }),
          }),
        )
      : new ArticleApi(
          new Configuration({
            basePath: config.API_BASE,
          }),
        ),
  );

  useEffect(() => {
    if (isAuthenticated) {
      setApi(
        new ArticleApi(
          new Configuration({
            basePath: config.API_BASE,
            accessToken: getAccessTokenSilently({
              authorizationParams: { audience: config.AUDIENCE },
            }),
          }),
        ),
      );
    }
  }, [isAuthenticated, config.API_BASE, config.AUDIENCE, getAccessTokenSilently]);

  return (
    <ArticleApiContext.Provider value={{ api }}>
      {children}
    </ArticleApiContext.Provider>
  );
};

export function useArticleApiCtx() {
  const context = useContext(ArticleApiContext);
  if (!context) {
    throw new Error(
      "useCategoryApiCtx should be used within a CategoryApiContext.Provider",
    );
  }
  return context;
}
