"use client";

import React, { useEffect, useState, Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { Search } from "lucide-react";
import { Item, Category } from "../../types";
import { api } from "../../lib/api";
import { ItemCard } from "../../components/ItemCard";
import { CategoryNav } from "../../components/CategoryNav";

function BrowseContent() {
  const searchParams = useSearchParams();
  const categorySlug = searchParams.get("category");
  const searchQuery = searchParams.get("q");
  const isAuctionParam = searchParams.get("auction");

  const [categories, setCategories] = useState<Category[]>([]);
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [filterType, setFilterType] = useState<"all" | "auction" | "buynow">(
    isAuctionParam === "true" ? "auction" : "all"
  );
  const [sortBy, setSortBy] = useState<"newest" | "price_asc" | "price_desc">("newest");
  const [localSearch, setLocalSearch] = useState<string>(searchQuery || "");

  useEffect(() => {
    api.getCategories().then(setCategories).catch(() => {});
  }, []);

  useEffect(() => {
    async function fetchItems() {
      setLoading(true);
      try {
        let result: Item[] = [];

        if (searchQuery) {
          result = await api.searchItems(searchQuery);
        } else if (categorySlug) {
          const cat = categories.find((c) => c.slug === categorySlug);
          if (cat) {
            result = await api.getItems({ category_id: cat.id });
          } else {
            result = await api.getItems();
          }
        } else {
          result = await api.getItems();
        }

        if (filterType === "auction") {
          result = result.filter((i) => i.is_auction);
        } else if (filterType === "buynow") {
          result = result.filter((i) => !i.is_auction);
        }

        if (sortBy === "price_asc") {
          result.sort((a, b) => parseFloat(a.price) - parseFloat(b.price));
        } else if (sortBy === "price_desc") {
          result.sort((a, b) => parseFloat(b.price) - parseFloat(a.price));
        } else {
          result.sort(
            (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
          );
        }

        setItems(result);
      } finally {
        setLoading(false);
      }
    }

    fetchItems();
  }, [categorySlug, searchQuery, filterType, sortBy, categories]);

  return (
    <div className="bg-[#F9F9F6] text-[#141413] min-h-screen pb-20">
      <CategoryNav categories={categories} activeSlug={categorySlug || undefined} />

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4 pb-4 mb-6 border-b border-[#E4E4DC]">
          <div>
            <h1 className="font-serif text-3xl font-bold tracking-tight text-[#141413]">
              {categorySlug
                ? categories.find((c) => c.slug === categorySlug)?.name || "Catalogue"
                : searchQuery
                ? `Results for "${searchQuery}"`
                : "Explore Catalogue"}
            </h1>
            <p className="text-xs text-[#8C8C85] mt-1">
              {items.length} {items.length === 1 ? "item" : "items"} available
            </p>
          </div>

          <form
            onSubmit={(e) => {
              e.preventDefault();
              if (localSearch.trim()) {
                window.location.href = `/browse?q=${encodeURIComponent(localSearch.trim())}`;
              }
            }}
            className="relative w-full sm:w-64"
          >
            <input
              type="text"
              placeholder="Search..."
              value={localSearch}
              onChange={(e) => setLocalSearch(e.target.value)}
              className="w-full pl-8 pr-3 py-1.5 text-xs bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
            />
            <Search className="w-3.5 h-3.5 text-[#8C8C85] absolute left-2.5 top-1/2 -translate-y-1/2" />
          </form>
        </div>

        {/* Filter Bar */}
        <div className="flex flex-wrap items-center justify-between gap-4 mb-8 text-xs">
          <div className="flex items-center gap-1">
            <button
              onClick={() => setFilterType("all")}
              className={`px-3 py-1.5 transition-colors uppercase tracking-wider text-[11px] border ${
                filterType === "all"
                  ? "border-[#141413] bg-[#141413] text-white"
                  : "border-transparent text-[#706F6A] hover:text-[#141413]"
              }`}
            >
              All
            </button>
            <button
              onClick={() => setFilterType("auction")}
              className={`px-3 py-1.5 transition-colors uppercase tracking-wider text-[11px] border flex items-center gap-1.5 ${
                filterType === "auction"
                  ? "border-[#141413] bg-[#141413] text-white"
                  : "border-transparent text-[#706F6A] hover:text-[#141413]"
              }`}
            >
              <span className="w-1.5 h-1.5 rounded-full bg-[#B83A24]" />
              Live Auctions
            </button>
            <button
              onClick={() => setFilterType("buynow")}
              className={`px-3 py-1.5 transition-colors uppercase tracking-wider text-[11px] border ${
                filterType === "buynow"
                  ? "border-[#141413] bg-[#141413] text-white"
                  : "border-transparent text-[#706F6A] hover:text-[#141413]"
              }`}
            >
              Buy It Now
            </button>
          </div>

          <div className="flex items-center gap-2">
            <span className="text-xs text-[#8C8C85]">Sort:</span>
            <select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value as any)}
              className="bg-white border border-[#E4E4DC] px-2 py-1 text-xs text-[#141413] focus:outline-none focus:border-[#141413]"
            >
              <option value="newest">Newest</option>
              <option value="price_asc">Price: Low to High</option>
              <option value="price_desc">Price: High to Low</option>
            </select>
          </div>
        </div>

        {/* Grid */}
        {loading ? (
          <div className="py-20 text-center text-xs text-[#8C8C85]">
            Loading items...
          </div>
        ) : items.length === 0 ? (
          <div className="py-20 text-center bg-white border border-[#E4E4DC]">
            <h3 className="font-serif text-lg font-bold text-[#141413] mb-1">
              No items found
            </h3>
            <p className="text-xs text-[#706F6A] mb-4">
              Try adjusting your filters or search terms.
            </p>
            <button
              onClick={() => {
                setFilterType("all");
                window.location.href = "/browse";
              }}
              className="px-4 py-2 text-xs uppercase tracking-wider bg-[#141413] text-white"
            >
              Reset Filters
            </button>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
            {items.map((itm) => (
              <ItemCard key={itm.id} item={itm} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default function BrowsePage() {
  return (
    <Suspense
      fallback={
        <div className="p-16 text-center text-xs text-[#8C8C85]">
          Loading catalogue...
        </div>
      }
    >
      <BrowseContent />
    </Suspense>
  );
}
