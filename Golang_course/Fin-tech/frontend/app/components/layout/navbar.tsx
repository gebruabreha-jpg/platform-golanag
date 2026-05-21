"use client";

import { useAuth } from "@/app/providers/auth-provider";

export function Navbar() {
  const { user } = useAuth();

  return (
    <header className="h-14 border-b flex items-center justify-between px-6 bg-white">
      <span className="text-sm text-zinc-500">Welcome back</span>
      <span className="text-sm font-medium">{user?.name || user?.email}</span>
    </header>
  );
}
