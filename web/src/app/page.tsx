"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { ArrowRight } from "lucide-react";
import { Item, Category } from "../types";
import { api } from "../lib/api";
import { formatCurrency, formatTimeRemaining } from "../lib/utils";
import { ItemCard } from "../components/ItemCard";
import { CategoryNav } from "../components/CategoryNav";

export default function HomePage() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [featuredAuction, setFeaturedAuction] = useState<Item | null>(null);
  const [liveAuctions, setLiveAuctions] = useState<Item[]>([]);
  const [buyNowItems, setBuyNowItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [timeLeft, setTimeLeft] = useState<number>(0);

  useEffect(() => {
    async function loadData() {
      try {
        setLoading(true);
        const [cats, auctions, allItems] = await Promise.all([
          api.getCategories().catch(() => []),
          api.getAuctions(12).catch(() => []),
          api.getItems({ limit: 16 }).catch(() => []),
        ]);

        setCategories(cats);
        setLiveAuctions(auctions);

        if (auctions.length > 0) {
          setFeaturedAuction(auctions[0]);
          if (auctions[0].valid_to) {
            const diff = Math.max(
              0,
              Math.floor((new Date(auctions[0].valid_to).getTime() - Date.now()) / 1000)
            );
            setTimeLeft(diff);
          }
        }

        const buyNow = allItems.filter((i) => !i.is_auction && i.status === "available");
        setBuyNowItems(buyNow.slice(0, 8));
      } finally {
        setLoading(false);
      }
    }
    loadData();
  }, []);

  // Timer for hero auction
  useEffect(() => {
    if (timeLeft > 0) {
      const interval = setInterval(() => {
        setTimeLeft((prev) => Math.max(0, prev - 1));
      }, 1000);
      return () => clearInterval(interval);
    }
  }, [timeLeft]);

  const heroTimer = formatTimeRemaining(timeLeft);

  return (
    <div className="flex flex-col gap-10 pb-20 bg-[#F9F9F6] text-[#141413]">
      {/* Category Bar */}
      <CategoryNav categories={categories} />

      {/* Hero Section */}
      <section className="border-b border-[#E4E4DC] pb-12 pt-4">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            {/* Left Column: Heading & Value Proposition */}
            <div className="lg:col-span-7 flex flex-col items-start gap-4">
              <span className="text-[11px] uppercase tracking-wider text-[#706F6A] font-medium">
                Live Online Marketplace
              </span>

              <h1 className="font-serif text-4xl sm:text-5xl font-bold tracking-tight text-[#141413] leading-[1.12]">
                Live Auctions for <br />
                <span className="italic font-normal text-[#55544E]">Rare &amp; Collectible</span> Objects.
              </h1>

              <p className="font-serif text-sm sm:text-base text-[#706F6A] max-w-lg leading-relaxed italic">
                Bid live in real time on fine watches, art, and rare electronics — with instant settlement powered by Khalti.
              </p>

              <div className="flex items-center gap-3 pt-2 text-xs uppercase tracking-wider">
                <Link
                  href="/browse?auction=true"
                  className="px-5 py-2.5 bg-[#141413] text-[#F9F9F6] hover:bg-[#2B2B28] transition-colors flex items-center gap-1.5 font-medium"
                >
                  <span>Explore Live Floor</span>
                  <ArrowRight className="w-3.5 h-3.5" />
                </Link>

                <Link
                  href="/create"
                  className="px-5 py-2.5 bg-white border border-[#E4E4DC] hover:border-[#141413] text-[#141413] transition-colors"
                >
                  Consign Item
                </Link>
              </div>
            </div>

            {/* Right Column: Featured Lot */}
            <div className="lg:col-span-5">
              {featuredAuction && (
                <div className="bg-white border border-[#E4E4DC] p-4 flex flex-col">
                  <div className="flex items-center justify-between text-[11px] pb-2 mb-3 border-b border-[#E4E4DC]">
                    <span className="font-bold text-[#B83A24] flex items-center gap-1.5">
                      <span className="w-1.5 h-1.5 rounded-full bg-[#B83A24] animate-pulse" />
                      Featured Lot
                    </span>
                    <span className="text-[#8C8C85] font-mono">
                      Closes in {heroTimer.formatted}
                    </span>
                  </div>

                  <div className="relative aspect-[4/3] w-full bg-[#F2F2EC] overflow-hidden mb-3">
                    <Image
                      src={
                        featuredAuction.image_url ||
                        "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=800&q=80"
                      }
                      alt={featuredAuction.name}
                      fill
                      className="object-cover"
                    />
                  </div>

                  <div className="flex items-end justify-between pt-1">
                    <div>
                      <h3 className="font-serif text-base font-bold text-[#141413] line-clamp-1">
                        {featuredAuction.name}
                      </h3>
                      <span className="font-mono text-base font-bold text-[#141413] block mt-0.5">
                        {formatCurrency(featuredAuction.price)}
                      </span>
                    </div>

                    <Link
                      href={`/item/${featuredAuction.id}`}
                      className="px-3.5 py-1.5 bg-[#141413] hover:bg-[#B83A24] text-[#F9F9F6] text-xs uppercase tracking-wider transition-colors"
                    >
                      Bid Now
                    </Link>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </section>

      {/* Live Auctions Section */}
      <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full">
        <div className="flex items-baseline justify-between pb-3 mb-6 border-b border-[#E4E4DC]">
          <div className="flex items-baseline gap-3">
            <h2 className="font-serif text-2xl font-bold text-[#141413]">
              Live Auctions
            </h2>
            <span className="text-xs text-[#8C8C85]">
              Real-time competitive bidding
            </span>
          </div>

          <Link
            href="/browse?auction=true"
            className="text-xs uppercase tracking-wider text-[#706F6A] hover:text-[#141413] flex items-center gap-1"
          >
            <span>View All ({liveAuctions.length})</span>
            <ArrowRight className="w-3.5 h-3.5" />
          </Link>
        </div>

        {liveAuctions.length === 0 ? (
          <div className="py-12 text-center bg-white border border-[#E4E4DC] text-xs text-[#706F6A]">
            No active live auctions right now.
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
            {liveAuctions.slice(0, 8).map((itm) => (
              <ItemCard key={itm.id} item={itm} />
            ))}
          </div>
        )}
      </section>

      {/* Buy It Now Section */}
      <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full">
        <div className="flex items-baseline justify-between pb-3 mb-6 border-b border-[#E4E4DC]">
          <div className="flex items-baseline gap-3">
            <h2 className="font-serif text-2xl font-bold text-[#141413]">
              Buy It Now
            </h2>
            <span className="text-xs text-[#8C8C85]">
              Instant acquisition with Khalti
            </span>
          </div>

          <Link
            href="/browse"
            className="text-xs uppercase tracking-wider text-[#706F6A] hover:text-[#141413] flex items-center gap-1"
          >
            <span>Full Catalogue</span>
            <ArrowRight className="w-3.5 h-3.5" />
          </Link>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
          {buyNowItems.slice(0, 8).map((itm) => (
            <ItemCard key={itm.id} item={itm} />
          ))}
        </div>
      </section>
    </div>
  );
}
