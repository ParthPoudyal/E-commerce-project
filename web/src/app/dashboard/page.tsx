"use client";

import React, { useEffect, useState, Suspense } from "react";
import Link from "next/link";
import Image from "next/image";
import { useSearchParams } from "next/navigation";
import { Plus } from "lucide-react";
import { Bid, PurchaseHydrated } from "../../types";
import { api } from "../../lib/api";
import { formatCurrency } from "../../lib/utils";
import { useAuth } from "../../context/AuthContext";

function DashboardContent() {
  const searchParams = useSearchParams();
  const initialTab = (searchParams.get("tab") as any) || "bids";
  const { user } = useAuth();

  const [activeTab, setActiveTab] = useState<"bids" | "purchases">(
    initialTab === "purchases" ? "purchases" : "bids"
  );
  const [bids, setBids] = useState<Bid[]>([]);
  const [purchases, setPurchases] = useState<PurchaseHydrated[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!user) return;

    async function loadData() {
      setLoading(true);
      try {
        const [bidsRes, purchRes] = await Promise.all([
          api.getUserBids().catch(() => []),
          api.getPurchases().catch(() => []),
        ]);
        setBids(bidsRes);
        setPurchases(purchRes);
      } finally {
        setLoading(false);
      }
    }
    loadData();
  }, [user]);

  if (!user) {
    return (
      <div className="max-w-md mx-auto px-4 py-20 text-center">
        <h2 className="font-serif text-2xl font-bold text-[#141413] mb-2">
          Sign In to View Dashboard
        </h2>
        <p className="text-xs text-[#706F6A] mb-6">
          Log in to track your bids and past purchases.
        </p>
        <Link
          href="/login"
          className="px-6 py-2 bg-[#141413] text-white text-xs uppercase tracking-wider font-medium"
        >
          Sign In
        </Link>
      </div>
    );
  }

  return (
    <div className="bg-[#F9F9F6] text-[#141413] min-h-screen py-10">
      <div className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* User Header */}
        <div className="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4 pb-4 mb-6 border-b border-[#E4E4DC]">
          <div>
            <h1 className="font-serif text-3xl font-bold tracking-tight text-[#141413]">
              {user.username}
            </h1>
            <p className="text-xs text-[#706F6A] mt-0.5">{user.email}</p>
          </div>

          <Link
            href="/create"
            className="px-3.5 py-1.5 bg-[#141413] text-white text-xs uppercase tracking-wider font-medium inline-flex items-center gap-1.5 self-start sm:self-auto"
          >
            <Plus className="w-3.5 h-3.5" />
            <span>List an Item</span>
          </Link>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-8 text-xs">
          <div className="p-4 bg-white border border-[#E4E4DC]">
            <span className="text-[#8C8C85] uppercase text-[10px] block">Bids Placed</span>
            <span className="text-2xl font-bold font-mono text-[#141413] mt-1 block">
              {bids.length}
            </span>
          </div>

          <div className="p-4 bg-white border border-[#E4E4DC]">
            <span className="text-[#8C8C85] uppercase text-[10px] block">Items Won</span>
            <span className="text-2xl font-bold font-mono text-[#141413] mt-1 block">
              {purchases.length}
            </span>
          </div>

          <div className="p-4 bg-white border border-[#E4E4DC]">
            <span className="text-[#8C8C85] uppercase text-[10px] block">Total Spent</span>
            <span className="text-2xl font-bold font-mono text-[#141413] mt-1 block">
              {formatCurrency(
                purchases.reduce((acc, curr) => acc + curr.purchase.price_paid, 0)
              )}
            </span>
          </div>
        </div>

        {/* Tabs */}
        <div className="flex items-center gap-6 border-b border-[#E4E4DC] mb-6 text-xs uppercase tracking-wider">
          <button
            onClick={() => setActiveTab("bids")}
            className={`pb-2.5 transition-colors ${
              activeTab === "bids"
                ? "border-b border-[#141413] text-[#141413] font-semibold"
                : "text-[#706F6A] hover:text-[#141413]"
            }`}
          >
            My Bids ({bids.length})
          </button>
          <button
            onClick={() => setActiveTab("purchases")}
            className={`pb-2.5 transition-colors ${
              activeTab === "purchases"
                ? "border-b border-[#141413] text-[#141413] font-semibold"
                : "text-[#706F6A] hover:text-[#141413]"
            }`}
          >
            Purchases ({purchases.length})
          </button>
        </div>

        {/* Content */}
        {loading ? (
          <div className="py-20 text-center text-xs text-[#8C8C85]">
            Loading...
          </div>
        ) : activeTab === "bids" ? (
          bids.length === 0 ? (
            <div className="py-16 text-center bg-white border border-[#E4E4DC]">
              <p className="text-sm font-semibold text-[#141413] mb-1">No active bids</p>
              <p className="text-xs text-[#706F6A] mb-4">
                Explore ongoing live auctions to place bids.
              </p>
              <Link
                href="/browse?auction=true"
                className="px-4 py-2 bg-[#141413] text-white text-xs uppercase tracking-wider"
              >
                Browse Auctions
              </Link>
            </div>
          ) : (
            <div className="space-y-2.5 text-xs">
              {bids.map((b) => (
                <div
                  key={b.id}
                  className="flex items-center justify-between p-3.5 bg-white border border-[#E4E4DC]"
                >
                  <div>
                    <span className="font-bold text-[#141413] block">
                      Lot #{b.item_id.substring(0, 8)}
                    </span>
                    <span className="text-[11px] text-[#8C8C85]">
                      {new Date(b.bid_time).toLocaleString()}
                    </span>
                  </div>

                  <div className="flex items-center gap-4">
                    <span className="font-mono text-sm font-bold text-[#141413]">
                      {formatCurrency(b.bid_amount)}
                    </span>

                    <Link
                      href={`/item/${b.item_id}`}
                      className="text-xs text-[#706F6A] hover:text-[#141413]"
                    >
                      View &rarr;
                    </Link>
                  </div>
                </div>
              ))}
            </div>
          )
        ) : purchases.length === 0 ? (
          <div className="py-16 text-center bg-white border border-[#E4E4DC]">
            <p className="text-sm font-semibold text-[#141413] mb-1">No purchases yet</p>
            <p className="text-xs text-[#706F6A] mb-4">
              Items purchased or won at auction will appear here.
            </p>
            <Link
              href="/browse"
              className="px-4 py-2 bg-[#141413] text-white text-xs uppercase tracking-wider"
            >
              Browse Catalogue
            </Link>
          </div>
        ) : (
          <div className="space-y-3 text-xs">
            {purchases.map((p) => {
              const itm = p.item;
              return (
                <div
                  key={p.purchase.id}
                  className="flex items-center justify-between p-3.5 bg-white border border-[#E4E4DC]"
                >
                  <div className="flex items-center gap-3">
                    {itm?.image_url && (
                      <div className="relative w-12 h-12 bg-[#F2F2EC] overflow-hidden flex-shrink-0">
                        <Image
                          src={itm.image_url}
                          alt={itm.name}
                          fill
                          className="object-cover"
                        />
                      </div>
                    )}
                    <div>
                      <h4 className="font-serif font-bold text-sm text-[#141413]">
                        {itm?.name || `Lot #${p.purchase.item_id.substring(0, 8)}`}
                      </h4>
                      <span className="text-[11px] text-[#8C8C85]">
                        {new Date(p.purchase.purchased_at).toLocaleDateString()} &bull; Paid via Khalti
                      </span>
                    </div>
                  </div>

                  <span className="font-mono text-sm font-bold text-[#141413]">
                    {formatCurrency(p.purchase.price_paid)}
                  </span>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
}

export default function DashboardPage() {
  return (
    <Suspense
      fallback={
        <div className="p-16 text-center text-xs text-[#8C8C85]">
          Loading dashboard...
        </div>
      }
    >
      <DashboardContent />
    </Suspense>
  );
}
