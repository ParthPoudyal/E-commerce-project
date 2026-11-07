"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { User } from "../types";
import { api } from "../lib/api";

interface AuthContextType {
  user: User | null;
  token: string | null;
  isLoading: boolean;
  login: (token: string, user: User) => void;
  logout: () => void;
  refreshUser: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  token: null,
  isLoading: true,
  login: () => {},
  logout: () => {},
  refreshUser: async () => {},
});

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const savedToken = localStorage.getItem("auctionhub_token");
    const savedUser = localStorage.getItem("auctionhub_user");

    if (savedToken) {
      setToken(savedToken);
      if (savedUser) {
        try {
          setUser(JSON.parse(savedUser));
        } catch {
          // ignore error
        }
      }
      // verify token with server
      api
        .getMe()
        .then((fetchedUser) => {
          setUser(fetchedUser);
          localStorage.setItem("auctionhub_user", JSON.stringify(fetchedUser));
        })
        .catch(() => {
          // Token expired or invalid
          logout();
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else {
      setIsLoading(false);
    }
  }, []);

  const login = (newToken: string, newUser: User) => {
    setToken(newToken);
    setUser(newUser);
    localStorage.setItem("auctionhub_token", newToken);
    localStorage.setItem("auctionhub_user", JSON.stringify(newUser));
  };

  const logout = () => {
    setToken(null);
    setUser(null);
    localStorage.removeItem("auctionhub_token");
    localStorage.removeItem("auctionhub_user");
  };

  const refreshUser = async () => {
    try {
      const u = await api.getMe();
      setUser(u);
      localStorage.setItem("auctionhub_user", JSON.stringify(u));
    } catch {
      // ignore
    }
  };

  return (
    <AuthContext.Provider value={{ user, token, isLoading, login, logout, refreshUser }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
