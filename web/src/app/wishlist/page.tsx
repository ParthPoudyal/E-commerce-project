"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { Trash2 } from "lucide-react";
import { WishlistItemHydrated } from "../../types";
import { api } from "../../lib/api";
import { formatCurrency } from "../../lib/utils";
import { useAuth } from "../../context/AuthContext";

export default function WishlistPage() {
  const { user } = useAuth();
  const [items, setItems] = useState<WishlistItemHydrated[]>([]);
  const [loading, setLoading] = useState(true);

  const loadWishlist = async () => {
    if (!user) {
      setLoading(false);
      return;
    }
    try {
      setLoading(true);
      const res = await api.getWishlist();
      setItems(res);
    } catch {
      // ignore
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadWishlist();
  }, [user]);

  const handleRemove = async (itemId: string) => {
    try {
      await api.removeFromWishlist(itemId);
      setItems((prev) => prev.filter((w) => w.item?.id !== itemId));
    } catch {
      // ignore
    }
  };

  if (!user) {
    return (
      <div className="max-w-md mx-auto px-4 py-20 text-center">
        <h2 className="font-serif text-2xl font-bold text-[#141413] mb-2">
          Sign In to View Saved Items
        </h2>
        <p className="text-xs text-[#706F6A] mb-6">
          Log in to track your saved auction lots.
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
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="pb-4 mb-6 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-3xl font-bold tracking-tight text-[#141413]">
            Saved Items
          </h1>
          <p className="text-xs text-[#706F6A] mt-1">
            {items.length} {items.length === 1 ? "item" : "items"} saved
          </p>
        </div>

        {loading ? (
          <div className="py-20 text-center text-xs text-[#8C8C85]">
            Loading saved items...
          </div>
        ) : items.length === 0 ? (
          <div className="py-16 text-center bg-white border border-[#E4E4DC]">
            <h3 className="font-serif text-lg font-bold text-[#141413] mb-1">
              No saved items
            </h3>
            <p className="text-xs text-[#706F6A] mb-4">
              Click the heart icon on any item to save it here.
            </p>
            <Link
              href="/browse"
              className="px-4 py-2 bg-[#141413] text-white text-xs uppercase tracking-wider"
            >
              Browse Catalogue
            </Link>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
            {items.map((wi) => {
              const itm = wi.item;
              if (!itm) return null;

              return (
                <div
                  key={itm.id}
                  className="flex flex-col bg-white border border-[#E4E4DC]"
                >
                  <div className="relative aspect-[4/3] bg-[#F2F2EC] overflow-hidden">
                    <Image
                      src={
                        itm.image_url ||
                        "https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=600&q=80"
                      }
                      alt={itm.name}
                      fill
                      className="object-cover"
                    />
                    <button
                      onClick={() => handleRemove(itm.id)}
                      className="absolute top-2 right-2 w-7 h-7 bg-white/90 border border-[#E4E4DC] hover:border-[#141413] flex items-center justify-center text-[#706F6A] hover:text-[#B83A24] transition-colors"
                      title="Remove"
                    >
                      <Trash2 className="w-3.5 h-3.5" />
                    </button>
                  </div>

                  <div className="p-4 flex flex-col justify-between flex-1 gap-3">
                    <div>
                      <Link
                        href={`/item/${itm.id}`}
                        className="font-serif font-bold text-base text-[#141413] hover:text-[#B83A24] transition-colors line-clamp-1"
                      >
                        {itm.name}
                      </Link>
                      <span className="font-mono text-sm font-bold text-[#141413] block mt-1">
                        {formatCurrency(itm.price)}
                      </span>
                    </div>

                    <Link
                      href={`/item/${itm.id}`}
                      className="text-xs font-medium text-[#706F6A] hover:text-[#141413] transition-colors"
                    >
                      View Details &rarr;
                    </Link>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
}
