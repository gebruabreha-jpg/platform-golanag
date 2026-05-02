const ALLOWED_ORIGINS = [
  "http://localhost:8080",
  "https://api.yourdomain.com",
];

function getApiUrl(): string {
  const url = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";
  const origin = new URL(url).origin;
  if (!ALLOWED_ORIGINS.includes(origin)) {
    throw new Error(`Untrusted API origin: ${origin}`);
  }
  return url;
}

const API_URL = getApiUrl();

async function request(path: string, options: RequestInit = {}) {
  const res = await fetch(`${API_URL}${path}`, {
    ...options,
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
  });

  const data = await res.json();

  if (!res.ok) {
    throw new Error(data.error || "Something went wrong");
  }

  return data;
}

export const api = {
  get: (path: string) => request(path),
  post: (path: string, body: Record<string, string>) =>
    request(path, { method: "POST", body: JSON.stringify(body) }),
  put: (path: string, body: Record<string, string>) =>
    request(path, { method: "PUT", body: JSON.stringify(body) }),
  delete: (path: string) => request(path, { method: "DELETE" }),
};
