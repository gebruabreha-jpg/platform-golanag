import { APP_NAME } from "@/app/lib/constants";

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50">
      <div className="w-full max-w-sm">
        <h2 className="text-center text-2xl font-bold mb-6">{APP_NAME}</h2>
        {children}
      </div>
    </div>
  );
}
