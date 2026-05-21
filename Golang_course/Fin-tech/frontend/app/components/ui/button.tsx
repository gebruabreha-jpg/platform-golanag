import { ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "danger";
  loading?: boolean;
}

export function Button({ variant = "primary", loading, children, disabled, className = "", ...props }: ButtonProps) {
  const base = "px-4 py-2 w-full rounded font-medium transition disabled:opacity-50";
  const variants = {
    primary: "bg-black text-white hover:bg-zinc-800",
    secondary: "bg-zinc-100 text-black hover:bg-zinc-200 border",
    danger: "bg-red-600 text-white hover:bg-red-700",
  };

  return (
    <button className={`${base} ${variants[variant]} ${className}`} disabled={disabled || loading} {...props}>
      {loading ? "Loading..." : children}
    </button>
  );
}
