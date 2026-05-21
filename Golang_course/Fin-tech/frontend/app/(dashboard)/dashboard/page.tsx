"use client";

import { useAuth } from "@/app/providers/auth-provider";
import { Card } from "@/app/components/ui/card";

export default function DashboardPage() {
  const { user } = useAuth();

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">Dashboard</h1>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card>
          <p className="text-sm text-zinc-500">Balance</p>
          <p className="text-2xl font-bold">$0.00</p>
        </Card>
        <Card>
          <p className="text-sm text-zinc-500">Transactions</p>
          <p className="text-2xl font-bold">0</p>
        </Card>
        <Card>
          <p className="text-sm text-zinc-500">Account</p>
          <p className="text-2xl font-bold">{user?.email}</p>
        </Card>
      </div>
    </div>
  );
}
