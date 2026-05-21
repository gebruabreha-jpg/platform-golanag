import type { UserPublic } from "@/types/index.js";

export interface AuthTokenPair {
  accessToken:  string;
  refreshToken: string;
  expiresIn:    string;
}

export interface LoginResult {
  tokens: AuthTokenPair;
  user:   UserPublic;
}

export interface RefreshResult {
  tokens: AuthTokenPair;
}
