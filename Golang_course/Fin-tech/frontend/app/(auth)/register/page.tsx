"use client";

import { useState } from "react";
import { useAuth } from "@/app/providers/auth-provider";
import { Input } from "@/app/components/ui/input";
import { Button } from "@/app/components/ui/button";
import { Card } from "@/app/components/ui/card";
import Link from "next/link";

export default function RegisterPage() {
  const { register } = useAuth();
  const [error, setError] = useState("");
  const [pending, setPending] = useState(false);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError("");
    setPending(true);

    const form = new FormData(e.currentTarget);
    const name = form.get("name") as string;
    const email = form.get("email") as string;
    const password = form.get("password") as string;

    if (password.length < 8) {
      setError("Password must be at least 8 characters");
      setPending(false);
      return;
    }

    try {
      await register(name, email, password);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Registration failed");
    } finally {
      setPending(false);
    }
  }

  return (
    <Card>
      <form onSubmit={handleSubmit} className="space-y-4">
        <h1 className="text-xl font-bold text-center">Create Account</h1>

        {error && <p className="text-red-500 text-sm text-center">{error}</p>}

        <Input name="name" type="text" placeholder="Full Name" required />
        <Input name="email" type="email" placeholder="Email" required />
        <Input name="password" type="password" placeholder="Password (min 8 chars)" minLength={8} required />

        <Button type="submit" loading={pending}>Create Account</Button>

        <p className="text-sm text-center">
          Have an account? <Link href="/login" className="underline">Login</Link>
        </p>
      </form>
    </Card>
  );
}
