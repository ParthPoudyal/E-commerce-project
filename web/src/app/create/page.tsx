"use client";

import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Loader2, AlertCircle, Upload } from "lucide-react";
import { Category } from "../../types";
import { api } from "../../lib/api";
import { useAuth } from "../../context/AuthContext";

export default function CreateListingPage() {
  const router = useRouter();
  const { user } = useAuth();

  const [categories, setCategories] = useState<Category[]>([]);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [categoryId, setCategoryId] = useState("");
  const [imageUrl, setImageUrl] = useState("");
  const [isAuction, setIsAuction] = useState(true);
  const [price, setPrice] = useState("");
  const [durationDays, setDurationDays] = useState(7);
  const [loading, setLoading] = useState(false);
  const [uploadingImage, setUploadingImage] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    setUploadingImage(true);
    setError(null);
    try {
      const res = await api.uploadImage(file);
      setImageUrl(res.url);
    } catch (err: any) {
      setError(err.message || "Failed to upload image");
    } finally {
      setUploadingImage(false);
    }
  };

  useEffect(() => {
    api.getCategories().then((cats) => {
      setCategories(cats);
      if (cats.length > 0) setCategoryId(cats[0].id);
    }).catch(() => {});
  }, []);

  const sampleImages = [
    { label: "Watch", url: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=800&q=80" },
    { label: "Camera", url: "https://images.unsplash.com/photo-1516035069371-29a1b244cc32?auto=format&fit=crop&w=800&q=80" },
    { label: "Art", url: "https://images.unsplash.com/photo-1579783902614-a3fb3927b675?auto=format&fit=crop&w=800&q=80" },
  ];

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!user) {
      router.push("/login");
      return;
    }

    if (!name || !price) {
      setError("Please provide an item name and price.");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const created = await api.createItem({
        name,
        description,
        category_id: categoryId || undefined,
        image_url: imageUrl || sampleImages[0].url,
        price: parseFloat(price),
        is_auction: isAuction,
        starting_bid: isAuction ? parseFloat(price) : undefined,
        duration_days: isAuction ? durationDays : undefined,
      });

      router.push(`/item/${created.id}`);
    } catch (err: any) {
      setError(err.message || "Failed to create listing");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-[#F9F9F6] text-[#141413] min-h-screen py-10">
      <div className="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="pb-4 mb-6 border-b border-[#E4E4DC]">
          <h1 className="font-serif text-3xl font-bold tracking-tight text-[#141413]">
            List an Item
          </h1>
          <p className="text-xs text-[#706F6A] mt-1">
            Post an item for live auction or immediate fixed-price checkout.
          </p>
        </div>

        {error && (
          <div className="mb-6 p-3 bg-red-50 border border-red-200 text-xs text-red-700 flex items-center gap-2">
            <AlertCircle className="w-4 h-4 flex-shrink-0" />
            <span>{error}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6 text-xs">
          {/* Details */}
          <div className="bg-white border border-[#E4E4DC] p-6 space-y-4">
            <div>
              <label className="block text-[#141413] font-semibold mb-1">
                Item Title *
              </label>
              <input
                type="text"
                placeholder="e.g. 1972 Vintage Rolex Submariner"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
                required
              />
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label className="block text-[#141413] font-semibold mb-1">
                  Category
                </label>
                <select
                  value={categoryId}
                  onChange={(e) => setCategoryId(e.target.value)}
                  className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
                >
                  {categories.map((c) => (
                    <option key={c.id} value={c.id}>
                      {c.name}
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <div className="flex items-center justify-between mb-1">
                  <label className="font-semibold text-[#141413]">Image URL</label>
                  <label className="text-xs text-[#B83A24] hover:underline cursor-pointer flex items-center gap-1">
                    <Upload className="w-3 h-3" />
                    <span>{uploadingImage ? "Uploading..." : "Upload File"}</span>
                    <input
                      type="file"
                      accept="image/*"
                      onChange={handleFileUpload}
                      className="hidden"
                      disabled={uploadingImage}
                    />
                  </label>
                </div>
                <input
                  type="url"
                  placeholder="Paste URL or upload"
                  value={imageUrl}
                  onChange={(e) => setImageUrl(e.target.value)}
                  className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
                />
              </div>
            </div>

            <div>
              <label className="block text-[#141413] font-semibold mb-1">
                Description
              </label>
              <textarea
                rows={4}
                placeholder="Condition, origin, accessories, etc."
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
              />
            </div>
          </div>

          {/* Pricing & Type */}
          <div className="bg-white border border-[#E4E4DC] p-6 space-y-4">
            <label className="block font-semibold text-[#141413]">Sale Format</label>
            <div className="grid grid-cols-2 gap-3">
              <button
                type="button"
                onClick={() => setIsAuction(true)}
                className={`py-2.5 px-3 border text-center transition-colors font-medium ${
                  isAuction
                    ? "border-[#141413] bg-[#141413] text-white"
                    : "border-[#E4E4DC] bg-white text-[#706F6A] hover:border-[#141413]"
                }`}
              >
                Live Auction
              </button>

              <button
                type="button"
                onClick={() => setIsAuction(false)}
                className={`py-2.5 px-3 border text-center transition-colors font-medium ${
                  !isAuction
                    ? "border-[#141413] bg-[#141413] text-white"
                    : "border-[#E4E4DC] bg-white text-[#706F6A] hover:border-[#141413]"
                }`}
              >
                Buy It Now
              </button>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-2">
              <div>
                <label className="block text-[#141413] font-semibold mb-1">
                  {isAuction ? "Starting Price (NPR) *" : "Price (NPR) *"}
                </label>
                <input
                  type="number"
                  placeholder="10000"
                  value={price}
                  onChange={(e) => setPrice(e.target.value)}
                  className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413] font-mono"
                  required
                />
              </div>

              {isAuction && (
                <div>
                  <label className="block text-[#141413] font-semibold mb-1">
                    Duration
                  </label>
                  <select
                    value={durationDays}
                    onChange={(e) => setDurationDays(parseInt(e.target.value))}
                    className="w-full px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413]"
                  >
                    <option value={1}>1 Day</option>
                    <option value={3}>3 Days</option>
                    <option value={5}>5 Days</option>
                    <option value={7}>7 Days</option>
                  </select>
                </div>
              )}
            </div>
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full py-3 px-4 bg-[#141413] hover:bg-[#2B2B28] text-white uppercase tracking-wider font-semibold transition-colors disabled:opacity-50 text-xs"
          >
            {loading ? "Publishing..." : "Publish Item"}
          </button>
        </form>
      </div>
    </div>
  );
}
