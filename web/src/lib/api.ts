import {
  User,
  Item,
  Bid,
  ItemDetailResponse,
  Category,
  CartItemHydrated,
  WishlistItemHydrated,
  PurchaseHydrated,
  PaymentTransaction,
} from "../types";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";

function getAuthHeader(): Record<string, string> {
  if (typeof window === "undefined") return {};
  const token = localStorage.getItem("auctionhub_token");
  if (!token) return {};
  return { Authorization: `Bearer ${token}` };
}

async function handleResponse<T>(res: Response): Promise<T> {
  const json = await res.json().catch(() => ({}));
  if (!res.ok || json.success === false) {
    const errorMsg = json.error || json.message || `Request failed with status ${res.status}`;
    throw new Error(errorMsg);
  }
  return json.data !== undefined ? json.data : json;
}

export const api = {
  // Auth
  async register(data: { username: string; email: string; password: string; address?: string; phone_no?: string }) {
    const res = await fetch(`${API_BASE}/auth/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    return handleResponse<{ token: string; user: User }>(res);
  },

  async login(data: { email: string; password: string }) {
    const res = await fetch(`${API_BASE}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    return handleResponse<{ token: string; user: User }>(res);
  },

  async getMe() {
    const res = await fetch(`${API_BASE}/auth/me`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<User>(res);
  },

  // Categories
  async getCategories() {
    const res = await fetch(`${API_BASE}/categories`);
    return handleResponse<Category[]>(res);
  },

  async getCategoryBySlug(slug: string) {
    const res = await fetch(`${API_BASE}/categories/${slug}`);
    return handleResponse<Category>(res);
  },

  // Items
  async getItems(params?: { category_id?: string; is_auction?: boolean; limit?: number; offset?: number }) {
    const query = new URLSearchParams();
    if (params?.category_id) query.set("category_id", params.category_id);
    if (params?.is_auction !== undefined) query.set("is_auction", String(params.is_auction));
    if (params?.limit) query.set("limit", String(params.limit));
    if (params?.offset) query.set("offset", String(params.offset));

    const res = await fetch(`${API_BASE}/items?${query.toString()}`);
    return handleResponse<Item[]>(res);
  },

  async getAuctions(limit = 20, offset = 0) {
    const res = await fetch(`${API_BASE}/items/auctions?limit=${limit}&offset=${offset}`);
    return handleResponse<Item[]>(res);
  },

  async getItem(id: string) {
    const res = await fetch(`${API_BASE}/items/${id}`);
    return handleResponse<ItemDetailResponse>(res);
  },

  async searchItems(q: string) {
    const res = await fetch(`${API_BASE}/items/search?q=${encodeURIComponent(q)}`);
    return handleResponse<Item[]>(res);
  },

  async createItem(data: {
    name: string;
    description: string;
    category_id?: string;
    image_url: string;
    price: number | string;
    is_auction: boolean;
    starting_bid?: number | string;
    duration_days?: number;
  }) {
    const res = await fetch(`${API_BASE}/items`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...getAuthHeader(),
      },
      body: JSON.stringify(data),
    });
    return handleResponse<Item>(res);
  },

  async deleteItem(id: string) {
    const res = await fetch(`${API_BASE}/items/${id}`, {
      method: "DELETE",
      headers: { ...getAuthHeader() },
    });
    return handleResponse<{ message: string }>(res);
  },

  async uploadImage(file: File) {
    const formData = new FormData();
    formData.append("image", file);
    const token = typeof window !== "undefined" ? localStorage.getItem("auctionhub_token") : null;
    const res = await fetch(`${API_BASE}/upload`, {
      method: "POST",
      headers: token ? { Authorization: `Bearer ${token}` } : {},
      body: formData,
    });
    return handleResponse<{ url: string; filename: string }>(res);
  },

  // Bids
  async placeBid(itemId: string, amount: number | string) {
    const res = await fetch(`${API_BASE}/bids`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...getAuthHeader(),
      },
      body: JSON.stringify({ item_id: itemId, amount: String(amount) }),
    });
    return handleResponse<{ bid: Bid; is_new_highest: boolean }>(res);
  },

  async getBidHistory(itemId: string) {
    const res = await fetch(`${API_BASE}/bids/item/${itemId}`);
    return handleResponse<Bid[]>(res);
  },

  async getUserBids() {
    const res = await fetch(`${API_BASE}/bids/user`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<Bid[]>(res);
  },

  // Cart
  async getCart() {
    const res = await fetch(`${API_BASE}/cart`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<CartItemHydrated[]>(res);
  },

  async addToCart(itemId: string) {
    const res = await fetch(`${API_BASE}/cart`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...getAuthHeader(),
      },
      body: JSON.stringify({ item_id: itemId }),
    });
    return handleResponse<any>(res);
  },

  async removeFromCart(itemId: string) {
    const res = await fetch(`${API_BASE}/cart/${itemId}`, {
      method: "DELETE",
      headers: { ...getAuthHeader() },
    });
    return handleResponse<{ message: string }>(res);
  },

  // Wishlist
  async getWishlist() {
    const res = await fetch(`${API_BASE}/wishlist`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<WishlistItemHydrated[]>(res);
  },

  async addToWishlist(itemId: string) {
    const res = await fetch(`${API_BASE}/wishlist`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...getAuthHeader(),
      },
      body: JSON.stringify({ item_id: itemId }),
    });
    return handleResponse<any>(res);
  },

  async removeFromWishlist(itemId: string) {
    const res = await fetch(`${API_BASE}/wishlist/${itemId}`, {
      method: "DELETE",
      headers: { ...getAuthHeader() },
    });
    return handleResponse<{ message: string }>(res);
  },

  async checkWishlisted(itemId: string) {
    const res = await fetch(`${API_BASE}/wishlist/check/${itemId}`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<boolean>(res);
  },

  // Purchases
  async getPurchases() {
    const res = await fetch(`${API_BASE}/purchases`, {
      headers: { ...getAuthHeader() },
    });
    return handleResponse<PurchaseHydrated[]>(res);
  },

  async buyNow(itemId: string) {
    const res = await fetch(`${API_BASE}/purchases/buy-now/${itemId}`, {
      method: "POST",
      headers: { ...getAuthHeader() },
    });
    return handleResponse<any>(res);
  },

  // Khalti Payment
  async initiateKhaltiPayment(itemId: string, amount?: number | string) {
    const res = await fetch(`${API_BASE}/payments/khalti/initiate`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...getAuthHeader(),
      },
      body: JSON.stringify({ item_id: itemId, amount: amount ? String(amount) : undefined }),
    });
    return handleResponse<{ transaction_id: string; pidx: string; payment_url: string }>(res);
  },

  async verifyKhaltiPayment(pidx: string) {
    const res = await fetch(`${API_BASE}/payments/khalti/verify`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ pidx }),
    });
    return handleResponse<PaymentTransaction>(res);
  },
};

