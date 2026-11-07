"use client";

import React, { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import {
  Heart,
  ShoppingCart,
  ArrowLeft,
  CreditCard,
} from "lucide-react";
import { ItemDetailResponse } from "../../../types";
import { api } from "../../../lib/api";
import { formatCurrency, formatTimeRemaining } from "../../../lib/utils";
import { useAuth } from "../../../context/AuthContext";
import { trackViewItem } from "../../../lib/gtag";
import { LiveBiddingModal } from "../../../components/LiveBiddingModal";
import { KhaltiPaymentModal } from "../../../components/KhaltiPaymentModal";

export default function ItemDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const { user } = useAuth();

  const [data, setData] = useState<ItemDetailResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isWishlisted, setIsWishlisted] = useState(false);
  const [cartSuccess, setCartSuccess] = useState(false);
  const [showKhaltiModal, setShowKhaltiModal] = useState(false);

  const loadItem = async () => {
    if (!id) return;
    try {
      setLoading(true);
      const res = await api.getItem(id);
      setData(res);

      trackViewItem({
        id: res.item.id,
        name: res.item.name,
        price: parseFloat(res.item.price),
      });

      if (user) {
        api.checkWishlisted(id).then(setIsWishlisted).catch(() => {});
      }
    } catch (err: any) {
      setError(err.message || "Failed to load item");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadItem();
  }, [id, user]);

  const handleToggleWishlist = async () => {
    if (!user) {
      router.push("/login");
      return;
    }
    try {
      if (isWishlisted) {
        await api.removeFromWishlist(id);
        setIsWishlisted(false);
      } else {
        await api.addToWishlist(id);
        setIsWishlisted(true);
      }
    } catch {
      // ignore
    }
  };

  const handleAddToCart = async () => {
    if (!user) {
      router.push("/login");
      return;
    }
    try {
      await api.addToCart(id);
      setCartSuccess(true);
      setTimeout(() => setCartSuccess(false), 2500);
    } catch {
      // ignore
    }
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 text-center text-xs text-[#8C8C85]">
        Loading item...
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="max-w-xl mx-auto px-4 py-20 text-center">
        <h2 className="font-serif text-2xl font-bold text-[#141413] mb-2">Item Not Found</h2>
        <p className="text-xs text-[#706F6A] mb-6">{error || "The requested item could not be found."}</p>
        <Link
          href="/browse"
          className="px-5 py-2.5 bg-[#141413] text-[#F9F9F6] text-xs uppercase tracking-wider"
        >
          Return to Catalogue
        </Link>
      </div>
    );
  }

  const { item, highest_bid, total_bids, time_left_seconds = 0 } = data;
  const isAuction = item.is_auction;
  const timer = formatTimeRemaining(time_left_seconds);
  const isClosed = timer.isExpired || item.status === "sold";
  const currentPrice = highest_bid ? highest_bid.bid_amount : item.price;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-[#F9F9F6] text-[#141413]">
      {/* Breadcrumb */}
      <div className="flex items-center gap-2 mb-6 text-xs text-[#706F6A]">
        <button
          onClick={() => router.back()}
          className="hover:text-[#141413] flex items-center gap-1"
        >
          <ArrowLeft className="w-3.5 h-3.5" />
          <span>Back</span>
        </button>
        <span>/</span>
        <Link href="/browse" className="hover:text-[#141413]">
          Catalogue
        </Link>
        {item.category && (
          <>
            <span>/</span>
            <Link href={`/browse?category=${item.category.slug}`} className="hover:text-[#141413]">
              {item.category.name}
            </Link>
          </>
        )}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
        {/* Left Column: Image & Description */}
        <div className="lg:col-span-7 flex flex-col gap-6">
          <div className="relative aspect-[4/3] w-full bg-[#F2F2EC] border border-[#E4E4DC] overflow-hidden">
            <Image
              src={
                item.image_url ||
                "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=1000&q=80"
              }
              alt={item.name}
              fill
              className="object-cover"
              priority
            />

            <button
              onClick={handleToggleWishlist}
              className="absolute top-3 right-3 w-8 h-8 bg-white border border-[#E4E4DC] hover:border-[#141413] flex items-center justify-center transition-colors"
              title="Save Item"
            >
              <Heart
                className={`w-4 h-4 ${isWishlisted ? "fill-[#B83A24] text-[#B83A24]" : "text-[#706F6A]"}`}
              />
            </button>
          </div>

          {/* Description */}
          <div className="bg-white border border-[#E4E4DC] p-6">
            <h3 className="font-bold text-xs uppercase tracking-wider text-[#141413] mb-3">
              Description
            </h3>
            <p className="text-sm text-[#55544E] leading-relaxed whitespace-pre-line mb-4">
              {item.description || "No description provided."}
            </p>
            <div className="pt-4 border-t border-[#F2F2EC] flex items-center justify-between text-xs text-[#8C8C85]">
              <span>Seller: @{item.posted_by?.username || "Verified"}</span>
              <span>Listed: {new Date(item.created_at).toLocaleDateString()}</span>
            </div>
          </div>
        </div>

        {/* Right Column: Pricing & Actions */}
        <div className="lg:col-span-5 flex flex-col gap-6">
          <div className="bg-white border border-[#E4E4DC] p-6 flex flex-col gap-4">
            <span className="text-[11px] uppercase tracking-wider text-[#706F6A]">
              {item.category?.name || "Curated Lot"}
            </span>

            <h1 className="font-serif text-2xl sm:text-3xl font-bold text-[#141413] leading-tight">
              {item.name}
            </h1>

            {/* Price Box */}
            <div className="p-4 bg-[#F9F9F6] border border-[#E4E4DC] flex items-baseline justify-between">
              <div>
                <span className="text-[11px] text-[#706F6A] block">
                  {isAuction ? "Current Bid" : "Price"}
                </span>
                <span className="font-mono text-3xl font-bold text-[#141413]">
                  {formatCurrency(currentPrice)}
                </span>
              </div>

              {isAuction && !isClosed && (
                <div className="text-right">
                  <span className="text-[11px] text-[#706F6A] block">Time Left</span>
                  <span className="font-mono text-xs font-bold text-[#B83A24]">
                    {timer.formatted}
                  </span>
                </div>
              )}
            </div>

            {/* Actions */}
            <div className="space-y-2 text-xs">
              <button
                onClick={() => setShowKhaltiModal(true)}
                disabled={item.status === "sold"}
                className="w-full py-3 px-4 bg-[#5C2D91] hover:bg-[#4a2278] text-white uppercase tracking-wider font-semibold flex items-center justify-center gap-2 transition-colors disabled:opacity-50"
              >
                <CreditCard className="w-4 h-4" />
                <span>Pay with Khalti &bull; {formatCurrency(currentPrice)}</span>
              </button>

              <button
                onClick={handleAddToCart}
                disabled={item.status === "sold"}
                className="w-full py-2.5 px-4 bg-white hover:bg-[#F2F2EC] text-[#141413] border border-[#E4E4DC] hover:border-[#141413] uppercase tracking-wider font-medium flex items-center justify-center gap-2 transition-colors disabled:opacity-50"
              >
                <ShoppingCart className="w-4 h-4 text-[#706F6A]" />
                <span>{cartSuccess ? "Added to Cart" : "Add to Cart"}</span>
              </button>
            </div>
          </div>

          {/* Live Bidding Console */}
          {isAuction && (
            <LiveBiddingModal
              item={item}
              highestBid={highest_bid}
              totalBids={total_bids}
              timeLeftSeconds={time_left_seconds}
              onBidPlaced={loadItem}
            />
          )}

          {/* Khalti Modal */}
          <KhaltiPaymentModal
            isOpen={showKhaltiModal}
            onClose={() => setShowKhaltiModal(false)}
            itemId={item.id}
            itemName={item.name}
            amount={currentPrice}
            onSuccess={() => {
              setShowKhaltiModal(false);
              loadItem();
            }}
          />
        </div>
      </div>
    </div>
  );
}
