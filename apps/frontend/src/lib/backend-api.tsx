import { jwtDecode } from "jwt-decode";

interface DecodedAccessToken {
  first_name: string;
  last_name: string;
  email: string;
}

export const doRequest = async (
  endpoint: string,
  {
    method = "GET",
    token = null,
    body = null as string | null,
    callbackError = null as ((error: Error) => void) | null,
    options = {},
  } = {}
): Promise<Record<string, string> | null> => {
  const apiBackendUrl = import.meta.env.VITE_API_BACKEND_URL;
  console.log("apiBackendUrl: ", apiBackendUrl);

  const url = new URL(apiBackendUrl || "");
  url.pathname = new URL(endpoint, url).pathname;
  url.search = new URL(endpoint, apiBackendUrl).search;

  try {
    const headers: HeadersInit = {
      "Content-Type": "application/json",
      Accept: "application/json",
    };
    if (token) headers.Authorization = `Bearer ${token}`;

    const res = await fetch(url.toString(), {
      method,
      headers,
      body,
      ...options,
    });

    if (res.ok) {
      const contentType = res.headers.get("Content-Type");
      if (contentType && contentType.includes("application/json")) {
        return await res.json();
      } else {
        return null;
      }
    }
  } catch (error) {
    if (
      error instanceof Error &&
      callbackError &&
      typeof callbackError === "function"
    ) {
      callbackError(error);
      return null;
    } else {
      console.error(error);
    }
  }
  return null;
};

export const login = async (
  email: string,
  password: string
): Promise<string | null> => {

  const response = await doRequest("/api/v1/auth/login/", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });

  if (!response) {
    return null;
  }

  if (response.access_token) {
    const decodedAccessToken = jwtDecode(
      response.access_token
    ) as DecodedAccessToken;

    const session = JSON.stringify({
      accessToken: response.access_token,
      user: {
        first_name: decodedAccessToken.first_name,
        last_name: decodedAccessToken.last_name,
        email: decodedAccessToken.email,
      },
    });

    return session;
  }
  return null;
};
