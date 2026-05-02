"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import { APP_NAME } from "@/app/lib/constants";

const links = [
  { href: "/dashboard", label: "Dashboard" },
  { href: "/billing", label: "Billing" },
  { href: "/settings", label: "Settings" },
];

export function Sidebar() {
  const pathname = usePathname();
  const { logout } = useAuth();

  return (
    <aside className="w-56 h-screen border-r bg-zinc-50 flex flex-col justify-between p-4">
      <div>
        <h1 className="text-lg font-bold mb-6">{APP_NAME}</h1>
        <nav className="space-y-1">
          {links.map((link) => (
            <Link
              key={link.href}
              href={link.href}
              className={`block px-3 py-2 rounded text-sm ${pathname === link.href ? "bg-black text-white" : "hover:bg-zinc-200"}`}
            >
              {link.label}
            </Link>
          ))}
        </nav>
      </div>
      <button onClick={logout} className="text-sm text-red-600 hover:underline text-left">
        Logout
      </button>
    </aside>
  );
}
