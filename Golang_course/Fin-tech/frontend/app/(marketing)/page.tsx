import Link from "next/link";

export default function LandingPage() {
  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] text-center px-4">
      <h1 className="text-5xl font-bold tracking-tight mb-4">
        Modern Banking,<br />Simplified.
      </h1>
      <p className="text-lg text-zinc-500 max-w-md mb-8">
        Send, receive, and manage your money with ease. Built for the future of finance.
      </p>
      <div className="flex gap-3">
        <Link href="/register" className="bg-black text-white px-6 py-3 rounded-lg font-medium hover:bg-zinc-800">
          Get Started Free
        </Link>
        <Link href="/pricing" className="border px-6 py-3 rounded-lg font-medium hover:bg-zinc-100">
          View Pricing
        </Link>
      </div>
    </div>
  );
}
