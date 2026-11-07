"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import {
  Search,
  Heart,
  ShoppingCart,
  Plus,
  ChevronDown,
  LogOut,
  Layers,
  FileText,
  CreditCard,
} from "lucide-react";
import { useAuth } from "../context/AuthContext";

export const Navbar: React.FC = () => {
  const router = useRouter();
  const { user, logout } = useAuth();
  const [searchQuery, setSearchQuery] = useState("");
  const [userDropdown, setUserDropdown] = useState(false);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      router.push(`/browse?q=${encodeURIComponent(searchQuery.trim())}`);
    }
  };

  return (
    <header className="sticky top-0 z-50 bg-[#F9F9F6]/95 backdrop-blur-sm border-b border-[#E4E4DC] text-[#141413]">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16 gap-6">
          {/* Brand */}
          <Link href="/" className="flex items-center gap-2.5 flex-shrink-0">
            <span className="font-serif text-xl sm:text-2xl font-bold tracking-tight text-[#141413]">
              Auction <span className="font-sans font-light text-base text-[#706F6A]">Studio</span>
            </span>
          </Link>

          {/* Clean Search */}
          <form onSubmit={handleSearch} className="flex-1 max-w-sm hidden md:block">
            <div className="relative">
              <input
                type="text"
                placeholder="Search auctions, watches, art..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full pl-8 pr-3 py-1.5 text-xs bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] transition-colors text-[#141413] placeholder:text-[#A0A096]"
              />
              <Search className="w-3.5 h-3.5 text-[#8C8C85] absolute left-2.5 top-1/2 -translate-y-1/2" />
            </div>
          </form>

          {/* Primary Nav */}
          <nav className="hidden lg:flex items-center gap-6 text-xs uppercase tracking-wider text-[#706F6A]">
            <Link
              href="/browse?auction=true"
              className="hover:text-[#B83A24] transition-colors flex items-center gap-1.5"
            >
              <span className="w-1.5 h-1.5 rounded-full bg-[#B83A24]" />
              Live Auctions
            </Link>
            <Link href="/browse" className="hover:text-[#141413] transition-colors">
              Explore All
            </Link>
          </nav>

          {/* User & Actions */}
          <div className="flex items-center gap-3">
            <Link
              href="/create"
              className="inline-flex items-center gap-1.5 px-3 py-1.5 text-xs text-[#141413] border border-[#141413] hover:bg-[#141413] hover:text-white transition-colors uppercase tracking-wider font-medium"
            >
              <Plus className="w-3.5 h-3.5" />
              <span>Consign</span>
            </Link>

            <Link
              href="/wishlist"
              className="p-1.5 text-[#706F6A] hover:text-[#141413] transition-colors"
              title="Saved Items"
            >
              <Heart className="w-4 h-4" />
            </Link>

            <Link
              href="/cart"
              className="p-1.5 text-[#706F6A] hover:text-[#141413] transition-colors"
              title="Cart"
            >
              <ShoppingCart className="w-4 h-4" />
            </Link>

            {user ? (
              <div className="relative">
                <button
                  onClick={() => setUserDropdown(!userDropdown)}
                  className="flex items-center gap-2 p-1 border border-[#E4E4DC] hover:border-[#141413] bg-white text-xs transition-colors"
                >
                  <div className="w-6 h-6 bg-[#141413] text-[#F9F9F6] flex items-center justify-center text-[11px] font-bold">
                    {user.username.charAt(0).toUpperCase()}
                  </div>
                  <span className="hidden sm:inline text-xs font-medium text-[#141413]">
                    {user.username}
                  </span>
                  <ChevronDown className="w-3.5 h-3.5 text-[#8C8C85]" />
                </button>

                {userDropdown && (
                  <div
                    className="absolute right-0 mt-1 w-48 bg-white border border-[#141413] shadow-md py-1 z-50 text-xs"
                    onMouseLeave={() => setUserDropdown(false)}
                  >
                    <div className="px-3 py-2 border-b border-[#E4E4DC]">
                      <p className="font-bold text-[#141413] truncate">{user.username}</p>
                      <p className="text-[10px] text-[#706F6A] truncate">{user.email}</p>
                    </div>

                    <Link
                      href="/dashboard"
                      onClick={() => setUserDropdown(false)}
                      className="flex items-center gap-2 px-3 py-2 text-[#141413] hover:bg-[#F2F2EC] transition-colors"
                    >
                      <Layers className="w-3.5 h-3.5 text-[#706F6A]" />
                      <span>Dashboard &amp; Bids</span>
                    </Link>

                    <Link
                      href="/payment/sandbox"
                      onClick={() => setUserDropdown(false)}
                      className="flex items-center gap-2 px-3 py-2 text-[#5C2D91] hover:bg-[#F2F2EC] transition-colors font-medium"
                    >
                      <CreditCard className="w-3.5 h-3.5 text-[#5C2D91]" />
                      <span>Khalti Sandbox</span>
                    </Link>

                    <div className="border-t border-[#E4E4DC] my-1" />

                    <button
                      onClick={() => {
                        setUserDropdown(false);
                        logout();
                      }}
                      className="flex items-center gap-2 w-full text-left px-3 py-2 text-[#B83A24] hover:bg-[#F2F2EC] transition-colors"
                    >
                      <LogOut className="w-3.5 h-3.5" />
                      <span>Sign Out</span>
                    </button>
                  </div>
                )}
              </div>
            ) : (
              <div className="flex items-center gap-2 text-xs">
                <Link
                  href="/login"
                  className="px-3 py-1.5 text-[#141413] hover:text-[#706F6A] transition-colors uppercase tracking-wider"
                >
                  Sign In
                </Link>
                <Link
                  href="/register"
                  className="px-3 py-1.5 bg-[#141413] text-[#F9F9F6] hover:bg-[#2B2B28] transition-colors uppercase tracking-wider"
                >
                  Join
                </Link>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  );
};
