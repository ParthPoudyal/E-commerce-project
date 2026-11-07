export interface User {
  id: string;
  username: string;
  email: string;
  address?: string;
  phone_no?: string;
  created_at?: string;
}

export type ItemStatus = "available" | "unavailable" | "sold" | "bidding";

export interface Item {
  id: string;
  name: string;
  description: string;
  category_id?: string;
  category?: Category;
  image_url: string;
  posted_by_id?: string;
  posted_by?: User;
  price: string;
  status: ItemStatus;
  valid_from: string;
  valid_to?: string;
  is_auction: boolean;
  starting_bid?: string;
  current_bid?: string;
  sold_price?: string;
  sold_at?: string;
  created_at: string;
}

export interface Bid {
  id: string;
  item_id: string;
  bidder_id: string;
  bid_amount: string;
  bid_time: string;
}

export interface ItemDetailResponse {
  item: Item;
  highest_bid?: Bid;
  total_bids: number;
  time_left_seconds?: number;
}

export interface Category {
  id: string;
  name: string;
  slug: string;
  description: string;
  image_url: string;
}

export interface CartItemHydrated {
  cart_item: {
    Id: string;
    ItemID: string;
    UserID: string;
    AddedAt: string;
  };
  item?: Item;
}

export interface WishlistItemHydrated {
  wishlist_item: {
    id: string;
    user_id: string;
    item_id: string;
    added_at: string;
  };
  item?: Item;
}

export interface PurchaseHydrated {
  purchase: {
    id: string;
    user_id: string;
    item_id: string;
    price_paid: number;
    purchased_at: string;
  };
  item?: Item;
}

export interface PaymentTransaction {
  id: string;
  item_id: string;
  user_id: string;
  amount: string;
  provider: string;
  provider_txn_id: string;
  status: "pending" | "completed" | "failed";
  payment_url: string;
  created_at: string;
  completed_at?: string;
}

