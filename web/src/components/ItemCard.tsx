"use client";

import React, { useState, useEffect } from "react";
import Link from "next/link";
import Image from "next/image";
import { Heart } from "lucide-react";
import { Item } from "../types";
import { formatCurrency, formatTimeRemaining } from "../lib/utils";
import { api } from "../lib/api";
import { useAuth } from "../context/AuthContext";

interface ItemCardProps {
  item: Item;
  highestBidAmount?: string;
  totalBids?: number;
  onWishlistChange?: () => void;
}

export const ItemCard: React.FC<ItemCardProps> = ({
  item,
  highestBidAmount,
  totalBids,
  onWishlistChange,
}) => {
  const { user } = useAuth();
  const [isWishlisted, setIsWishlisted] = useState(false);
  const [timeLeft, setTimeLeft] = useState<number>(0);

  useEffect(() => {
    if (item.valid_to) {
      const targetTime = new Date(item.valid_to).getTime();
      const now = Date.now();
      const diffSec = Math.max(0, Math.floor((targetTime - now) / 1000));
      setTimeLeft(diffSec);

      const interval = setInterval(() => {
        setTimeLeft((prev) => {
          if (prev <= 1) {
            clearInterval(interval);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);

      return () => clearInterval(interval);
    }
  }, [item.valid_to]);

  useEffect(() => {
    if (user && item.id) {
      api.checkWishlisted(item.id).then(setIsWishlisted).catch(() => {});
    }
  }, [user, item.id]);

  const toggleWishlist = async (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    if (!user) {
      window.location.href = "/login";
      return;
    }

    try {
      if (isWishlisted) {
        await api.removeFromWishlist(item.id);
        setIsWishlisted(false);
      } else {
        await api.addToWishlist(item.id);
        setIsWishlisted(true);
      }
      onWishlistChange?.();
    } catch {
      // ignore
    }
  };

  const timerInfo = formatTimeRemaining(timeLeft);
  const displayPrice = highestBidAmount || item.price;
  const isAuction = item.is_auction;
  const isLive = isAuction && item.status === "bidding" && !timerInfo.isExpired;

  return (
    <article className="group flex flex-col bg-white border border-[#E4E4DC] hover:border-[#141413] transition-colors">
      {/* Product Image Frame */}
      <div className="relative aspect-[4/3] w-full bg-[#F2F2EC] overflow-hidden">
        <Link href={`/item/${item.id}`} className="block w-full h-full">
          <Image
            src={
              item.image_url ||
              "https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=600&q=80"
            }
            alt={item.name}
            fill
            sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
            className="object-cover group-hover:scale-105 transition-transform duration-300 ease-out"
          />
        </Link>

        {/* Minimal Live Status Tag */}
        {isLive && (
          <div className="absolute top-2 left-2 flex items-center gap-1.5 px-2 py-0.5 bg-white/95 border border-[#E4E4DC] text-[10px] font-medium text-[#141413]">
            <span className="w-1.5 h-1.5 rounded-full bg-[#B83A24] animate-pulse" />
            <span>{timerInfo.formatted}</span>
          </div>
        )}

        {/* Wishlist Button */}
        <button
          onClick={toggleWishlist}
          className="absolute top-2 right-2 w-7 h-7 bg-white/90 border border-[#E4E4DC] hover:border-[#141413] flex items-center justify-center transition-colors"
          title="Save Lot"
        >
          <Heart
            className={`w-3.5 h-3.5 ${
              isWishlisted ? "fill-[#B83A24] text-[#B83A24]" : "text-[#706F6A]"
            }`}
          />
        </button>
      </div>

      {/* Lot Details */}
      <div className="p-3.5 flex flex-col justify-between flex-1 gap-2">
        <div>
          <Link href={`/item/${item.id}`}>
            <h3 className="font-serif text-sm font-bold text-[#141413] line-clamp-1 group-hover:text-[#B83A24] transition-colors">
              {item.name}
            </h3>
          </Link>
          <p className="text-[11px] text-[#8C8C85] mt-0.5">
            {item.category?.name || "Curated"}
          </p>
        </div>

        {/* Pricing */}
        <div className="pt-2 border-t border-[#F2F2EC] flex items-baseline justify-between text-xs">
          <div>
            <span className="font-mono font-bold text-[#141413]">
              {formatCurrency(displayPrice)}
            </span>
            {isAuction && (
              <span className="text-[10px] text-[#8C8C85] ml-1.5">
                ({totalBids ?? 0} {totalBids === 1 ? "bid" : "bids"})
              </span>
            )}
          </div>

          <Link
            href={`/item/${item.id}`}
            className="text-[11px] font-medium text-[#706F6A] group-hover:text-[#141413] transition-colors"
          >
            {isAuction ? "Bid →" : "Buy →"}
          </Link>
        </div>
      </div>
    </article>
  );
};
