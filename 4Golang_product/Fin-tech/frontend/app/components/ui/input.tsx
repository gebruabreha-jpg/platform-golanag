import { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

export function Input({ label, error, className = "", ...props }: InputProps) {
  return (
    <div className="space-y-1">
      {label && <label className="text-sm font-medium text-zinc-700">{label}</label>}
      <input
        className={`border p-2 w-full rounded focus:outline-none focus:ring-2 focus:ring-black ${error ? "border-red-500" : "border-zinc-300"} ${className}`}
        {...props}
      />
      {error && <p className="text-red-500 text-xs">{error}</p>}
    </div>
  );
}
