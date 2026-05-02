import { Card } from "@/app/components/ui/card";

export default function BillingPage() {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">Billing</h1>
      <Card>
        <p className="text-sm text-zinc-500">Current Plan</p>
        <p className="text-xl font-bold">Free</p>
        <p className="text-sm text-zinc-400 mt-2">Upgrade to unlock more features.</p>
      </Card>
    </div>
  );
}
