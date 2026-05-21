"use client";

import { useState } from "react";
import { api } from "@/app/lib/api";
import { Input } from "@/app/components/ui/input";
import { Button } from "@/app/components/ui/button";
import { Card } from "@/app/components/ui/card";
import Link from "next/link";

export default function ForgotPasswordPage() {
  const [sent, setSent] = useState(false);
  const [error, setError] = useState("");
  const [pending, setPending] = useState(false);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError("");
    setPending(true);

    const form = new FormData(e.currentTarget);
    const email = form.get("email") as string;

    try {
      await api.post("/forgot-password", { email });
      setSent(true);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Failed to send reset email");
    } finally {
      setPending(false);
    }
  }

  if (sent) {
    return (
      <Card>
        <div className="text-center space-y-2">
          <h1 className="text-xl font-bold">Check your email</h1>
          <p className="text-sm text-zinc-500">We sent a password reset link to your email.</p>
          <Link href="/login" className="text-sm underline">Back to Login</Link>
        </div>
      </Card>
    );
  }

  return (
    <Card>
      <form onSubmit={handleSubmit} className="space-y-4">
        <h1 className="text-xl font-bold text-center">Forgot Password</h1>

        {error && <p className="text-red-500 text-sm text-center">{error}</p>}

        <Input name="email" type="email" placeholder="Email" required />
        <Button type="submit" loading={pending}>Send Reset Link</Button>

        <p className="text-sm text-center">
          <Link href="/login" className="underline">Back to Login</Link>
        </p>
      </form>
    </Card>
  );
}
