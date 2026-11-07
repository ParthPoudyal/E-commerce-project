import React from "react";
import Link from "next/link";

export const Footer: React.FC = () => {
  return (
    <footer className="bg-[#F2F2EC] border-t border-[#E4E4DC] text-[#141413] mt-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          {/* Brand */}
          <div className="space-y-3">
            <Link href="/" className="font-serif text-lg font-bold tracking-tight text-[#141413]">
              Auction <span className="font-sans font-light text-sm text-[#706F6A]">Studio</span>
            </Link>
            <p className="font-serif text-xs text-[#706F6A] leading-relaxed italic max-w-xs">
              Live auction marketplace for verified luxury watches, art, and rare collectibles. Secured by Khalti.
            </p>
          </div>

          {/* Explore */}
          <div className="text-xs">
            <h5 className="font-bold uppercase tracking-wider text-[#141413] mb-3">Explore</h5>
            <ul className="space-y-2 text-[#706F6A]">
              <li><Link href="/browse?auction=true" className="hover:text-[#141413]">Live Auctions</Link></li>
              <li><Link href="/browse?category=watches-jewelry" className="hover:text-[#141413]">Watches &amp; Jewelry</Link></li>
              <li><Link href="/browse?category=art-collectibles" className="hover:text-[#141413]">Art &amp; Design</Link></li>
              <li><Link href="/browse" className="hover:text-[#141413]">All Lots</Link></li>
            </ul>
          </div>

          {/* Consign & Account */}
          <div className="text-xs">
            <h5 className="font-bold uppercase tracking-wider text-[#141413] mb-3">Consign</h5>
            <ul className="space-y-2 text-[#706F6A]">
              <li><Link href="/create" className="hover:text-[#141413]">List an Item</Link></li>
              <li><Link href="/dashboard" className="hover:text-[#141413]">Seller Dashboard</Link></li>
              <li><Link href="/cart" className="hover:text-[#141413]">Shopping Cart</Link></li>
              <li><Link href="/wishlist" className="hover:text-[#141413]">Saved Items</Link></li>
            </ul>
          </div>

          {/* Payments */}
          <div className="text-xs">
            <h5 className="font-bold uppercase tracking-wider text-[#141413] mb-3">Payments</h5>
            <ul className="space-y-2 text-[#706F6A]">
              <li><span className="text-[#5C2D91] font-semibold">Khalti Sandbox Active</span></li>
              <li><Link href="/payment/sandbox" className="hover:text-[#141413]">Test Payment Terminal</Link></li>
              <li><span className="text-[#8C8C85]">Test: 9800000000 / MPIN: 1111</span></li>
            </ul>
          </div>
        </div>

        {/* Bottom Bar */}
        <div className="mt-10 pt-6 border-t border-[#E4E4DC] flex flex-col sm:flex-row items-center justify-between gap-3 text-[11px] text-[#8C8C85]">
          <p>© 2026 Auction Studio. All rights reserved.</p>
          <div className="flex items-center gap-4">
            <span className="hover:text-[#141413] cursor-pointer">Terms</span>
            <span>&bull;</span>
            <span className="hover:text-[#141413] cursor-pointer">Privacy</span>
            <span>&bull;</span>
            <span className="hover:text-[#141413] cursor-pointer">Khalti Escrow</span>
          </div>
        </div>
      </div>
    </footer>
  );
};
