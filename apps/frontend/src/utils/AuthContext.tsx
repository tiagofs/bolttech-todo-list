import React, { createContext, useState, useEffect, ReactNode } from 'react';
import { useCookies } from 'react-cookie';
import { sessionOptions } from './config';


type AuthContextType = {
  authState: Record<string, string> | null;
  setAuthState: React.Dispatch<React.SetStateAction<Record<string, string> | null>>;
}

type AuthProviderProps = {
  children: ReactNode;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [authState, setAuthState] = useState<Record<string, string> | null>(null);
  console.log('sessionOptions.cookieName: ', sessionOptions.cookieName)
  const [cookies] = useCookies([sessionOptions.cookieName]);

  useEffect(() => {
    const checkAuth = async () => {
      console.log('cookies: ', cookies)
      if (cookies && Object.keys(cookies).length === 0) {
        console.log("setting authState to null");
        setAuthState(null)
      } else {
        setAuthState(cookies);
      }
    };

    checkAuth();
  }, []);

  return (
    <AuthContext.Provider value={{ authState, setAuthState }}>
      {children}
    </AuthContext.Provider>
  );
};