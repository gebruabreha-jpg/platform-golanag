import { ReactNode } from "react";

export function Card({ children, className = "" }: { children: ReactNode; className?: string }) {
  return (
    <div className={`border rounded-xl p-6 bg-white shadow-sm ${className}`}>
      {children}
    </div>
  );
}
