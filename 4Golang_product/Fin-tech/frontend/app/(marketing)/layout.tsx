import Link from "next/link";
import { APP_NAME } from "@/app/lib/constants";

export default function MarketingLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen flex flex-col">
      <header className="h-14 border-b flex items-center justify-between px-6">
        <Link href="/" className="font-bold">{APP_NAME}</Link>
        <div className="flex gap-4 text-sm">
          <Link href="/pricing" className="hover:underline">Pricing</Link>
          <Link href="/login" className="hover:underline">Login</Link>
          <Link href="/register" className="bg-black text-white px-3 py-1 rounded">Get Started</Link>
        </div>
      </header>
      <main className="flex-1">{children}</main>
    </div>
  );
}
