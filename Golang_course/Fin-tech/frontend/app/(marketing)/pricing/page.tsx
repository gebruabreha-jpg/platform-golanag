import Link from "next/link";
import { Card } from "@/app/components/ui/card";

const plans = [
  { name: "Free", price: "$0", features: ["1 Account", "Basic transfers", "Email support"] },
  { name: "Pro", price: "$9/mo", features: ["5 Accounts", "Instant transfers", "Priority support", "Analytics"] },
  { name: "Business", price: "$29/mo", features: ["Unlimited accounts", "API access", "Dedicated support", "Custom reports"] },
];

export default function PricingPage() {
  return (
    <div className="py-16 px-4 max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold text-center mb-2">Pricing</h1>
      <p className="text-center text-zinc-500 mb-10">Choose the plan that fits your needs</p>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {plans.map((plan) => (
          <Card key={plan.name} className="flex flex-col justify-between">
            <div>
              <h2 className="text-lg font-bold">{plan.name}</h2>
              <p className="text-2xl font-bold mt-2 mb-4">{plan.price}</p>
              <ul className="space-y-2 text-sm text-zinc-600">
                {plan.features.map((f) => (
                  <li key={f}>✓ {f}</li>
                ))}
              </ul>
            </div>
            <Link href="/register" className="mt-6 block text-center bg-black text-white py-2 rounded hover:bg-zinc-800">
              Get Started
            </Link>
          </Card>
        ))}
      </div>
    </div>
  );
}
