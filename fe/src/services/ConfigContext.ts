import { createContext, useContext } from "react";

export interface Config {
  API_BASE: string;
  AUTH_0_DOMAIN: string;
  AUTH_0_CLIENT_ID: string;
  AUDIENCE: string;
}

export const ConfigContext = createContext<Config | null>(null);

export const useConfig = () => {
  return useContext(ConfigContext);
};
