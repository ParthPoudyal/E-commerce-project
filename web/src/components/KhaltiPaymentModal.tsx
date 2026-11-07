"use client";

import React, { useState } from "react";
import { X, CheckCircle2, AlertCircle, Loader2 } from "lucide-react";
import { formatCurrency } from "../lib/utils";
import { api } from "../lib/api";
import { trackPaymentCompleted } from "../lib/gtag";

interface KhaltiPaymentModalProps {
  isOpen: boolean;
  onClose: () => void;
  itemId: string;
  itemName: string;
  amount: number | string;
  onSuccess: () => void;
}

export const KhaltiPaymentModal: React.FC<KhaltiPaymentModalProps> = ({
  isOpen,
  onClose,
  itemId,
  itemName,
  amount,
  onSuccess,
}) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [step, setStep] = useState<"summary" | "khalti_sandbox" | "success">("summary");

  // Sandbox inputs
  const [mobile, setMobile] = useState("9800000000");
  const [mpin, setMpin] = useState("1111");
  const [otp, setOtp] = useState("987654");
  const [pidx, setPidx] = useState("");

  if (!isOpen) return null;

  const handleInitiate = async () => {
    setLoading(true);
    setError(null);

    try {
      const res = await api.initiateKhaltiPayment(itemId, amount);
      setPidx(res.pidx);

      if (res.payment_url.startsWith("https://test-pay.khalti.com")) {
        window.location.href = res.payment_url;
      } else {
        setStep("khalti_sandbox");
      }
    } catch (err: any) {
      setError(err.message || "Failed to initiate payment");
    } finally {
      setLoading(false);
    }
  };

  const handleSandboxVerify = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const numAmount = typeof amount === "string" ? parseFloat(amount) : amount;
      await api.verifyKhaltiPayment(pidx);
      trackPaymentCompleted(pidx, numAmount, itemId);
      setStep("success");
      setTimeout(() => {
        onSuccess();
      }, 1500);
    } catch (err: any) {
      setError(err.message || "Payment verification failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-[#141413]/50 backdrop-blur-sm animate-fade-in">
      <div className="relative w-full max-w-sm bg-white border border-[#E4E4DC] shadow-xl text-[#141413] text-xs">
        {/* Header */}
        <div className="px-5 py-3.5 bg-[#5C2D91] text-white flex items-center justify-between">
          <div className="flex items-center gap-2">
            <span className="font-bold tracking-tight">Khalti Payment</span>
            <span className="text-[10px] text-purple-200">Sandbox</span>
          </div>
          <button
            onClick={onClose}
            className="w-6 h-6 hover:bg-white/10 flex items-center justify-center text-white transition-colors"
          >
            <X className="w-4 h-4" />
          </button>
        </div>

        {/* Modal Body */}
        <div className="p-5">
          {error && (
            <div className="mb-4 p-2.5 bg-red-50 border border-red-200 text-red-700 flex items-center gap-2">
              <AlertCircle className="w-4 h-4 flex-shrink-0" />
              <span>{error}</span>
            </div>
          )}

          {step === "summary" && (
            <div className="space-y-4">
              <div className="p-3.5 bg-[#F9F9F6] border border-[#E4E4DC]">
                <p className="font-serif font-bold text-sm text-[#141413] mb-2 line-clamp-1">
                  {itemName}
                </p>
                <div className="pt-2 border-t border-[#E4E4DC] flex items-center justify-between">
                  <span className="text-[#706F6A]">Amount</span>
                  <span className="font-mono font-bold text-base text-[#141413]">
                    {formatCurrency(amount)}
                  </span>
                </div>
              </div>

              <button
                onClick={handleInitiate}
                disabled={loading}
                className="w-full py-2.5 bg-[#5C2D91] hover:bg-[#4a2278] text-white font-medium uppercase tracking-wider transition-colors disabled:opacity-50"
              >
                {loading ? "Connecting..." : "Continue to Payment"}
              </button>
            </div>
          )}

          {step === "khalti_sandbox" && (
            <form onSubmit={handleSandboxVerify} className="space-y-3.5">
              <div className="text-center pb-2 border-b border-[#E4E4DC]">
                <h4 className="font-serif font-bold text-sm text-[#141413]">
                  Pay {formatCurrency(amount)}
                </h4>
                <p className="text-[10px] text-[#8C8C85] mt-0.5">
                  Test credentials pre-filled
                </p>
              </div>

              <div>
                <label className="block text-[11px] text-[#706F6A] mb-1">
                  Mobile Number
                </label>
                <input
                  type="text"
                  value={mobile}
                  onChange={(e) => setMobile(e.target.value)}
                  className="w-full px-2.5 py-1.5 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] font-mono text-[#141413]"
                  required
                />
              </div>

              <div className="grid grid-cols-2 gap-2">
                <div>
                  <label className="block text-[11px] text-[#706F6A] mb-1">
                    MPIN
                  </label>
                  <input
                    type="password"
                    value={mpin}
                    onChange={(e) => setMpin(e.target.value)}
                    className="w-full px-2.5 py-1.5 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] font-mono text-[#141413]"
                    required
                  />
                </div>
                <div>
                  <label className="block text-[11px] text-[#706F6A] mb-1">
                    OTP
                  </label>
                  <input
                    type="text"
                    value={otp}
                    onChange={(e) => setOtp(e.target.value)}
                    className="w-full px-2.5 py-1.5 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] font-mono text-[#141413]"
                    required
                  />
                </div>
              </div>

              <button
                type="submit"
                disabled={loading}
                className="w-full py-2.5 bg-[#5C2D91] hover:bg-[#4a2278] text-white font-medium uppercase tracking-wider transition-colors disabled:opacity-50 mt-2"
              >
                {loading ? "Verifying..." : `Authorize Payment`}
              </button>
            </form>
          )}

          {step === "success" && (
            <div className="py-6 text-center">
              <CheckCircle2 className="w-10 h-10 text-emerald-600 mx-auto mb-2" />
              <h4 className="font-serif text-base font-bold text-[#141413] mb-1">
                Payment Successful
              </h4>
              <p className="text-xs text-[#706F6A]">
                The transaction has been confirmed.
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
