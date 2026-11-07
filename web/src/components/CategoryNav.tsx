"use client";

import React from "react";
import Link from "next/link";
import { Category } from "../types";

interface CategoryNavProps {
  categories: Category[];
  activeSlug?: string;
}

export const CategoryNav: React.FC<CategoryNavProps> = ({ categories, activeSlug }) => {
  return (
    <div className="w-full border-b border-[#E4E4DC] bg-white overflow-x-auto no-scrollbar">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center gap-6 py-2.5 text-xs uppercase tracking-wider min-w-max">
          <Link
            href="/browse"
            className={`pb-0.5 transition-colors ${
              !activeSlug
                ? "text-[#141413] font-semibold border-b border-[#141413]"
                : "text-[#706F6A] hover:text-[#141413]"
            }`}
          >
            All Lots
          </Link>

          {categories.map((cat) => {
            const isActive = activeSlug === cat.slug;

            return (
              <Link
                key={cat.id}
                href={`/browse?category=${cat.slug}`}
                className={`pb-0.5 transition-colors ${
                  isActive
                    ? "text-[#141413] font-semibold border-b border-[#141413]"
                    : "text-[#706F6A] hover:text-[#141413]"
                }`}
              >
                {cat.name}
              </Link>
            );
          })}
        </div>
      </div>
    </div>
  );
};
