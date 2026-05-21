"use client";

import { useAuth } from "@/app/providers/auth-provider";
import { Card } from "@/app/components/ui/card";

export default function SettingsPage() {
  const { user } = useAuth();

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">Settings</h1>
      <Card>
        <div className="space-y-3">
          <div>
            <p className="text-sm text-zinc-500">Name</p>
            <p className="font-medium">{user?.name}</p>
          </div>
          <div>
            <p className="text-sm text-zinc-500">Email</p>
            <p className="font-medium">{user?.email}</p>
          </div>
        </div>
      </Card>
    </div>
  );
}
