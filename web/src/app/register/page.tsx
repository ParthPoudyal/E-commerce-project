"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { AlertCircle } from "lucide-react";
import { api } from "../../lib/api";
import { useAuth } from "../../context/AuthContext";

export default function RegisterPage() {
  const router = useRouter();
  const { login } = useAuth();

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [address, setAddress] = useState("");
  const [phoneNo, setPhoneNo] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const res = await api.register({
        username,
        email,
        password,
        address,
        phone_no: phoneNo,
      });
      login(res.token, res.user);
      router.push("/");
    } catch (err: any) {
      setError(err.message || "Registration failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-[85vh] flex items-center justify-center px-4 py-16 bg-[#F9F9F6] text-[#141413]">
      <div className="w-full max-w-sm bg-white border border-[#E4E4DC] p-8 text-xs">
        <div className="text-center pb-5 mb-5 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-2xl font-bold text-[#141413]">
            Create Account
          </h1>
          <p className="text-xs text-[#706F6A] mt-1">
            Join the auction community
          </p>
        </div>

        {error && (
          <div className="mb-4 p-2.5 bg-red-50 border border-red-200 text-xs text-red-700 flex items-center gap-2">
            <AlertCircle className="w-4 h-4 flex-shrink-0" />
            <span>{error}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-3.5">
          <div>
            <label className="block font-medium text-[#141413] mb-1">
              Username *
            </label>
            <input
              type="text"
              placeholder="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              required
            />
          </div>

          <div>
            <label className="block font-medium text-[#141413] mb-1">
              Email *
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
              Password *
            </label>
            <input
              type="password"
              placeholder="Minimum 6 characters"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              required
              minLength={6}
            />
          </div>

          <div className="grid grid-cols-2 gap-3">
            <div>
              <label className="block font-medium text-[#141413] mb-1">
                City
              </label>
              <input
                type="text"
                placeholder="Kathmandu"
                value={address}
                onChange={(e) => setAddress(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              />
            </div>

            <div>
              <label className="block font-medium text-[#141413] mb-1">
                Phone
              </label>
              <input
                type="tel"
                placeholder="9800000000"
                value={phoneNo}
                onChange={(e) => setPhoneNo(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              />
            </div>
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-2.5 px-4 bg-[#141413] hover:bg-[#2B2B28] text-white uppercase tracking-wider font-semibold transition-colors disabled:opacity-50 mt-3"
          >
            {loading ? "Registering..." : "Register"}
          </button>
        </form>

        <div className="mt-5 text-center text-xs text-[#706F6A]">
          Already have an account?{" "}
          <Link href="/login" className="font-semibold text-[#141413] hover:underline">
            Sign In
          </Link>
        </div>
      </div>
    </div>
  );
}
