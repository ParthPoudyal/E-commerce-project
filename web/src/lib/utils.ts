import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatCurrency(amount: number | string | undefined): string {
  if (amount === undefined || amount === null) return "NPR 0.00";
  const num = typeof amount === "string" ? parseFloat(amount) : amount;
  return new Intl.NumberFormat("en-NP", {
    style: "currency",
    currency: "NPR",
    minimumFractionDigits: 2,
  }).format(num);
}

export function formatTimeRemaining(seconds: number): {
  days: number;
  hours: number;
  minutes: number;
  seconds: number;
  isExpired: boolean;
  formatted: string;
} {
  if (seconds <= 0) {
    return { days: 0, hours: 0, minutes: 0, seconds: 0, isExpired: true, formatted: "Ended" };
  }

  const days = Math.floor(seconds / (3600 * 24));
  const hours = Math.floor((seconds % (3600 * 24)) / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const secs = Math.floor(seconds % 60);

  let formatted = "";
  if (days > 0) {
    formatted = `${days}d ${hours}h ${minutes}m`;
  } else if (hours > 0) {
    formatted = `${hours}h ${minutes}m ${secs}s`;
  } else {
    formatted = `${minutes}m ${secs}s`;
  }

  return { days, hours, minutes, seconds: secs, isExpired: false, formatted };
}

