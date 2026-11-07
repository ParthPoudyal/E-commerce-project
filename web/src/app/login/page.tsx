"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Loader2, AlertCircle } from "lucide-react";
import { api } from "../../lib/api";
import { useAuth } from "../../context/AuthContext";

export default function LoginPage() {
  const router = useRouter();
  const { login } = useAuth();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const res = await api.login({ email, password });
      login(res.token, res.user);
      router.push("/");
    } catch (err: any) {
      setError(err.message || "Invalid email or password");
    } finally {
      setLoading(false);
    }
  };

  const handleFillDemo = (demoEmail: string) => {
    setEmail(demoEmail);
    setPassword("password123");
  };

  return (
    <div className="min-h-[80vh] flex items-center justify-center px-4 py-16 bg-[#F9F9F6] text-[#141413]">
      <div className="w-full max-w-sm bg-white border border-[#E4E4DC] p-8 text-xs">
        <div className="text-center pb-5 mb-5 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-2xl font-bold text-[#141413]">
            Sign In
          </h1>
          <p className="text-xs text-[#706F6A] mt-1">
            Access your bids and account
          </p>
        </div>

        {error && (
          <div className="mb-4 p-2.5 bg-red-50 border border-red-200 text-xs text-red-700 flex items-center gap-2">
            <AlertCircle className="w-4 h-4 flex-shrink-0" />
            <span>{error}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block font-medium text-[#141413] mb-1">
              Email
            </label>
            <input
              type="email"
              placeholder="name@domain.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              required
            />
          </div>

          <div>
            <label className="block font-medium text-[#141413] mb-1">
              Password
            </label>
            <input
              type="password"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              required
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-2.5 px-4 bg-[#141413] hover:bg-[#2B2B28] text-white uppercase tracking-wider font-semibold transition-colors disabled:opacity-50 mt-2"
          >
            {loading ? "Signing in..." : "Sign In"}
          </button>
        </form>

        {/* Demo Accounts */}
        <div className="mt-6 pt-4 border-t border-[#E4E4DC] text-center">
          <span className="text-[11px] text-[#8C8C85] block mb-2">
            Demo Accounts:
          </span>
          <div className="flex justify-center gap-2 text-[11px]">
            <button
              type="button"
              onClick={() => handleFillDemo("parth@ecom.np")}
              className="px-2.5 py-1 bg-[#F2F2EC] hover:bg-[#E4E4DC] text-[#141413] transition-colors"
            >
              @parth
            </button>
            <button
              type="button"
              onClick={() => handleFillDemo("viraj@ecom.np")}
              className="px-2.5 py-1 bg-[#F2F2EC] hover:bg-[#E4E4DC] text-[#141413] transition-colors"
            >
              @viraj
            </button>
          </div>
        </div>

        <div className="mt-5 text-center text-xs text-[#706F6A]">
          Don&apos;t have an account?{" "}
          <Link href="/register" className="font-semibold text-[#141413] hover:underline">
            Register
          </Link>
        </div>
      </div>
    </div>
  );
}
