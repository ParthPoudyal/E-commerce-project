"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { Trash2, CreditCard } from "lucide-react";
import { CartItemHydrated } from "../../types";
import { api } from "../../lib/api";
import { formatCurrency } from "../../lib/utils";
import { useAuth } from "../../context/AuthContext";
import { KhaltiPaymentModal } from "../../components/KhaltiPaymentModal";

export default function CartPage() {
  const { user } = useAuth();
  const [cartItems, setCartItems] = useState<CartItemHydrated[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedItemForPayment, setSelectedItemForPayment] = useState<{
    id: string;
    name: string;
    price: number;
  } | null>(null);

  const loadCart = async () => {
    if (!user) {
      setLoading(false);
      return;
    }
    try {
      setLoading(true);
      const items = await api.getCart();
      setCartItems(items);
    } catch {
      // ignore
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadCart();
  }, [user]);

  const handleRemove = async (itemId: string) => {
    try {
      await api.removeFromCart(itemId);
      setCartItems((prev) => prev.filter((c) => c.item?.id !== itemId));
    } catch {
      // ignore
    }
  };

  if (!user) {
    return (
      <div className="max-w-md mx-auto px-4 py-20 text-center">
        <h2 className="font-serif text-2xl font-bold text-[#141413] mb-2">
          Sign In to View Cart
        </h2>
        <p className="text-xs text-[#706F6A] mb-6">
          Log in with your account to access your saved items.
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

  const subtotal = cartItems.reduce((acc, curr) => {
    return acc + (curr.item ? parseFloat(curr.item.price) : 0);
  }, 0);

  return (
    <div className="bg-[#F9F9F6] text-[#141413] min-h-screen py-10">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="pb-4 mb-6 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-3xl font-bold tracking-tight text-[#141413]">
            Shopping Cart
          </h1>
          <p className="text-xs text-[#706F6A] mt-1">
            {cartItems.length} {cartItems.length === 1 ? "item" : "items"} in cart
          </p>
        </div>

        {loading ? (
          <div className="py-20 text-center text-xs text-[#8C8C85]">
            Loading cart...
          </div>
        ) : cartItems.length === 0 ? (
          <div className="py-16 text-center bg-white border border-[#E4E4DC]">
            <h3 className="font-serif text-lg font-bold text-[#141413] mb-1">
              Your cart is empty
            </h3>
            <p className="text-xs text-[#706F6A] mb-4">
              Explore active items in the catalogue.
            </p>
            <Link
              href="/browse"
              className="px-4 py-2 bg-[#141413] text-white text-xs uppercase tracking-wider"
            >
              Browse Items
            </Link>
          </div>
        ) : (
          <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
            {/* Items */}
            <div className="lg:col-span-8 space-y-3">
              {cartItems.map((ci) => {
                const itm = ci.item;
                if (!itm) return null;

                return (
                  <div
                    key={itm.id}
                    className="flex items-center gap-4 p-4 bg-white border border-[#E4E4DC]"
                  >
                    <div className="relative w-16 h-16 bg-[#F2F2EC] overflow-hidden flex-shrink-0">
                      <Image
                        src={
                          itm.image_url ||
                          "https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=400&q=80"
                        }
                        alt={itm.name}
                        fill
                        className="object-cover"
                      />
                    </div>

                    <div className="flex-1 min-w-0">
                      <Link
                        href={`/item/${itm.id}`}
                        className="font-serif font-bold text-sm text-[#141413] hover:text-[#B83A24] transition-colors line-clamp-1"
                      >
                        {itm.name}
                      </Link>
                      <span className="font-mono text-sm font-bold text-[#141413] block mt-1">
                        {formatCurrency(itm.price)}
                      </span>
                    </div>

                    <div className="flex items-center gap-3">
                      <button
                        onClick={() =>
                          setSelectedItemForPayment({
                            id: itm.id,
                            name: itm.name,
                            price: parseFloat(itm.price),
                          })
                        }
                        className="px-3 py-1.5 bg-[#5C2D91] hover:bg-[#4a2278] text-white text-xs uppercase font-medium"
                      >
                        Khalti Pay
                      </button>

                      <button
                        onClick={() => handleRemove(itm.id)}
                        className="p-1.5 text-[#8C8C85] hover:text-[#B83A24] transition-colors"
                        title="Remove"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                );
              })}
            </div>

            {/* Summary */}
            <div className="lg:col-span-4 bg-white border border-[#E4E4DC] p-5 text-xs">
              <h3 className="font-bold text-[#141413] uppercase mb-3">
                Order Summary
              </h3>

              <div className="flex justify-between py-2 border-b border-[#F2F2EC] text-[#706F6A]">
                <span>Subtotal</span>
                <span className="font-mono font-bold text-[#141413]">
                  {formatCurrency(subtotal)}
                </span>
              </div>

              <div className="flex justify-between py-2 border-b border-[#F2F2EC] text-[#706F6A]">
                <span>Shipping</span>
                <span>Calculated at checkout</span>
              </div>

              <div className="flex justify-between py-3 font-bold text-sm text-[#141413] mb-4">
                <span>Total</span>
                <span className="font-mono">{formatCurrency(subtotal)}</span>
              </div>

              {cartItems.length > 0 && (
                <button
                  onClick={() => {
                    const first = cartItems[0].item;
                    if (first) {
                      setSelectedItemForPayment({
                        id: first.id,
                        name: first.name,
                        price: parseFloat(first.price),
                      });
                    }
                  }}
                  className="w-full py-2.5 bg-[#5C2D91] hover:bg-[#4a2278] text-white uppercase tracking-wider font-semibold text-xs flex items-center justify-center gap-1.5"
                >
                  <CreditCard className="w-3.5 h-3.5" />
                  <span>Checkout with Khalti</span>
                </button>
              )}
            </div>
          </div>
        )}

        {/* Modal */}
        {selectedItemForPayment && (
          <KhaltiPaymentModal
            isOpen={!!selectedItemForPayment}
            onClose={() => setSelectedItemForPayment(null)}
            itemId={selectedItemForPayment.id}
            itemName={selectedItemForPayment.name}
            amount={selectedItemForPayment.price}
            onSuccess={() => {
              setSelectedItemForPayment(null);
              loadCart();
            }}
          />
        )}
      </div>
    </div>
  );
}
