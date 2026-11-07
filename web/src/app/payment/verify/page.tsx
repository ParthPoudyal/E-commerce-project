"use client";

import React, { useEffect, useState, Suspense } from "react";
import { useSearchParams, useRouter } from "next/navigation";
import Link from "next/link";
import { CheckCircle2, AlertCircle, Loader2, ArrowRight } from "lucide-react";
import { api } from "../../../lib/api";
import { trackPaymentCompleted } from "../../../lib/gtag";

function VerifyContent() {
  const searchParams = useSearchParams();
  const router = useRouter();

  const pidx = searchParams.get("pidx");
  const [status, setStatus] = useState<"verifying" | "success" | "failed">("verifying");
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  useEffect(() => {
    if (!pidx) {
      setStatus("failed");
      setErrorMessage("No payment index (pidx) provided in callback.");
      return;
    }

    async function verify() {
      try {
        const tx = await api.verifyKhaltiPayment(pidx!);
        trackPaymentCompleted(pidx!, parseFloat(tx.amount), tx.item_id);
        setStatus("success");
      } catch (err: any) {
        setStatus("failed");
        setErrorMessage(err.message || "Failed to verify transaction with Khalti API.");
      }
    }

    verify();
  }, [pidx]);

  return (
    <div className="min-h-[75vh] flex items-center justify-center px-4 py-16 bg-[#F9F9F6] text-[#141413]">
      <div className="w-full max-w-md bg-white border border-[#141413] shadow-xl p-8 text-center font-mono text-xs">
        {status === "verifying" && (
          <div>
            <div className="w-12 h-12 border border-[#141413] border-t-transparent animate-spin mx-auto mb-4" />
            <h2 className="font-serif text-xl font-bold text-[#141413] mb-1">
              Verifying Khalti Settlement
            </h2>
            <p className="text-[#706F6A] mb-4 italic">
              Executing handshake with Khalti sandbox API endpoint...
            </p>
            <span className="text-[10px] text-[#8C8C85] block">PIDX: {pidx}</span>
          </div>
        )}

        {status === "success" && (
          <div>
            <div className="w-12 h-12 bg-emerald-700 text-white flex items-center justify-center mx-auto mb-4 font-bold">
              ✓
            </div>
            <span className="text-[10px] text-[#5C2D91] font-bold uppercase tracking-wider block mb-1">
              [ ESCROW TRANSACTION CAPTURED ]
            </span>
            <h2 className="font-serif text-2xl font-bold text-[#141413] mb-2">
              Settlement Completed
            </h2>
            <p className="text-xs text-[#706F6A] mb-6 italic leading-relaxed">
              Your Khalti transaction has been authenticated. The specimen is registered as sold to your collector account.
            </p>
            <div className="flex flex-col gap-2">
              <Link
                href="/dashboard?tab=purchases"
                className="w-full py-2.5 px-4 bg-[#141413] hover:bg-[#B83A24] text-[#F9F9F6] uppercase tracking-wider font-semibold transition-colors flex items-center justify-center gap-1.5"
              >
                <span>Inspect Settlement Receipt</span>
                <ArrowRight className="w-3.5 h-3.5" />
              </Link>
              <Link
                href="/"
                className="w-full py-2.5 px-4 bg-[#F2F2EC] hover:bg-[#E4E4DC] text-[#141413] border border-[#E4E4DC] uppercase tracking-wider transition-colors"
              >
                Return to Floor
              </Link>
            </div>
          </div>
        )}

        {status === "failed" && (
          <div>
            <div className="w-12 h-12 bg-[#B83A24] text-white flex items-center justify-center mx-auto mb-4 font-bold">
              !
            </div>
            <h2 className="font-serif text-xl font-bold text-[#141413] mb-1">
              Settlement Incomplete
            </h2>
            <p className="text-xs text-[#B83A24] mb-6 italic">{errorMessage}</p>
            <Link
              href="/"
              className="inline-block px-5 py-2.5 bg-[#141413] text-[#F9F9F6] uppercase tracking-wider"
            >
              Return to Marketplace
            </Link>
          </div>
        )}
      </div>
    </div>
  );
}

export default function PaymentVerifyPage() {
  return (
    <Suspense
      fallback={
        <div className="p-16 text-center font-mono text-xs text-[#8C8C85]">
          [ VERIFYING SETTLEMENT STATUS... ]
        </div>
      }
    >
      <VerifyContent />
    </Suspense>
  );
}
