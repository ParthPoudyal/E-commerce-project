export const GA_TRACKING_ID = process.env.NEXT_PUBLIC_GA_ID || "G-ECOM-AUCTION26";

// Log page views
export const pageview = (url: string) => {
  if (typeof window !== "undefined" && (window as any).gtag) {
    (window as any).gtag("config", GA_TRACKING_ID, {
      page_path: url,
    });
  }
};

// Log specific events
export const event = ({
  action,
  category,
  label,
  value,
  params,
}: {
  action: string;
  category?: string;
  label?: string;
  value?: number;
  params?: Record<string, any>;
}) => {
  if (typeof window !== "undefined" && (window as any).gtag) {
    (window as any).gtag("event", action, {
      event_category: category,
      event_label: label,
      value: value,
      ...params,
    });
  }
};

// E-commerce tracking helpers
export const trackViewItem = (item: { id: string; name: string; price: number; category?: string }) => {
  event({
    action: "view_item",
    category: "ecommerce",
    params: {
      currency: "NPR",
      value: item.price,
      items: [
        {
          item_id: item.id,
          item_name: item.name,
          price: item.price,
          item_category: item.category,
        },
      ],
    },
  });
};

export const trackBidPlaced = (itemId: string, itemName: string, bidAmount: number) => {
  event({
    action: "bid_placed",
    category: "auction",
    label: itemName,
    value: bidAmount,
    params: {
      item_id: itemId,
      amount: bidAmount,
      currency: "NPR",
    },
  });
};

export const trackPaymentCompleted = (transactionId: string, amount: number, itemId: string) => {
  event({
    action: "purchase",
    category: "ecommerce",
    label: transactionId,
    value: amount,
    params: {
      transaction_id: transactionId,
      currency: "NPR",
      value: amount,
      item_id: itemId,
    },
  });
};

