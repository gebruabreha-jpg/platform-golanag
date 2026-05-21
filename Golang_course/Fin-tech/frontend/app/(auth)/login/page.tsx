"use client";

import { useState } from "react";
import { useAuth } from "@/app/providers/auth-provider";
import { Input } from "@/app/components/ui/input";
import { Button } from "@/app/components/ui/button";
import { Card } from "@/app/components/ui/card";
import Link from "next/link";

export default function LoginPage() {
  const { login } = useAuth();
  const [error, setError] = useState("");
  const [pending, setPending] = useState(false);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError("");
    setPending(true);

    const form = new FormData(e.currentTarget);
    const email = form.get("email") as string;
    const password = form.get("password") as string;

    try {
      await login(email, password);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Login failed");
    } finally {
      setPending(false);
    }
  }

  return (
    <Card>
      <form onSubmit={handleSubmit} className="space-y-4">
        <h1 className="text-xl font-bold text-center">Login</h1>

        {error && <p className="text-red-500 text-sm text-center">{error}</p>}

        <Input name="email" type="email" placeholder="Email" required />
        <Input name="password" type="password" placeholder="Password" required />

        <Button type="submit" loading={pending}>Login</Button>

        <div className="flex justify-between text-sm">
          <Link href="/forgot-password" className="text-zinc-500 hover:underline">Forgot password?</Link>
          <Link href="/register" className="underline">Register</Link>
        </div>
      </form>
    </Card>
  );
}
