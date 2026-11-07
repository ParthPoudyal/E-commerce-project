"use client";

import React, { useState, Suspense } from "react";
import { useSearchParams, useRouter } from "next/navigation";
import { formatCurrency } from "../../../lib/utils";

function SandboxContent() {
  const searchParams = useSearchParams();
  const router = useRouter();

  const pidx = searchParams.get("pidx") || "khalti_test_default";
  const amount = searchParams.get("amount") || "5000";

  const [mobile, setMobile] = useState("9800000000");
  const [mpin, setMpin] = useState("1111");
  const [otp, setOtp] = useState("987654");
  const [loading, setLoading] = useState(false);

  const handlePay = (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    setTimeout(() => {
      router.push(`/payment/verify?pidx=${encodeURIComponent(pidx)}`);
    }, 1000);
  };

  return (
    <div className="min-h-[80vh] flex items-center justify-center px-4 py-16 bg-[#F9F9F6] text-[#141413]">
      <div className="w-full max-w-sm bg-white border border-[#E4E4DC] p-8 text-xs">
        <div className="text-center pb-5 mb-5 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-2xl font-bold text-[#141413]">
            Khalti Sandbox
          </h1>
          <p className="font-mono text-xl font-bold text-[#5C2D91] mt-2">
            {formatCurrency(amount)}
          </p>
        </div>

        <form onSubmit={handlePay} className="space-y-4">
          <div>
            <label className="block text-[#706F6A] mb-1">
              Test Phone
            </label>
            <input
              type="text"
              value={mobile}
              onChange={(e) => setMobile(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413] font-mono"
              required
            />
          </div>

          <div className="grid grid-cols-2 gap-3">
            <div>
              <label className="block text-[#706F6A] mb-1">
                MPIN
              </label>
              <input
                type="password"
                value={mpin}
                onChange={(e) => setMpin(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413] font-mono"
                required
              />
            </div>

            <div>
              <label className="block text-[#706F6A] mb-1">
                OTP
              </label>
              <input
                type="text"
                value={otp}
                onChange={(e) => setOtp(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413] font-mono"
                required
              />
            </div>
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-2.5 bg-[#5C2D91] hover:bg-[#4a2278] text-white font-medium uppercase tracking-wider transition-colors disabled:opacity-50 mt-2"
          >
            {loading ? "Processing..." : `Authorize Payment`}
          </button>
        </form>
      </div>
    </div>
  );
}

export default function SandboxPage() {
  return (
    <Suspense
      fallback={
        <div className="p-16 text-center text-xs text-[#8C8C85]">
          Loading terminal...
        </div>
      }
    >
      <SandboxContent />
    </Suspense>
  );
}
