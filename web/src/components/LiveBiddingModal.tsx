"use client";

import React, { useState, useEffect, useRef } from "react";
import { Clock, AlertTriangle, CheckCircle2 } from "lucide-react";
import { Item, Bid } from "../types";
import { formatCurrency, formatTimeRemaining } from "../lib/utils";
import { api } from "../lib/api";
import { useAuth } from "../context/AuthContext";
import { trackBidPlaced } from "../lib/gtag";

interface LiveBiddingModalProps {
  item: Item;
  highestBid?: Bid;
  totalBids: number;
  timeLeftSeconds?: number;
  onBidPlaced?: () => void;
}

export const LiveBiddingModal: React.FC<LiveBiddingModalProps> = ({
  item,
  highestBid: initialHighest,
  totalBids: initialCount,
  timeLeftSeconds = 0,
  onBidPlaced,
}) => {
  const { user } = useAuth();
  const [currentHighest, setCurrentHighest] = useState<Bid | undefined>(initialHighest);
  const [bidsCount, setBidsCount] = useState<number>(initialCount);
  const [bidsList, setBidsList] = useState<Bid[]>([]);
  const [customAmount, setCustomAmount] = useState<string>("");
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [outbidAlert, setOutbidAlert] = useState<string | null>(null);
  const [successAlert, setSuccessAlert] = useState<string | null>(null);
  const [errorAlert, setErrorAlert] = useState<string | null>(null);
  const [timeLeft, setTimeLeft] = useState<number>(timeLeftSeconds);
  const socketRef = useRef<WebSocket | null>(null);

  // Fetch initial bid history
  useEffect(() => {
    if (item.id) {
      api
        .getBidHistory(item.id)
        .then((history) => {
          setBidsList(history);
          if (history.length > 0) {
            setCurrentHighest(history[0]);
            setBidsCount(history.length);
          }
        })
        .catch(() => {});
    }
  }, [item.id]);

  // Countdown timer
  useEffect(() => {
    if (timeLeftSeconds > 0) {
      setTimeLeft(timeLeftSeconds);
      const timer = setInterval(() => {
        setTimeLeft((prev) => Math.max(0, prev - 1));
      }, 1000);
      return () => clearInterval(timer);
    }
  }, [timeLeftSeconds]);

  // WebSocket Live Sync
  useEffect(() => {
    if (!item.id) return;

    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    const host = process.env.NEXT_PUBLIC_WS_URL || "localhost:8080";
    const token = typeof window !== "undefined" ? localStorage.getItem("auctionhub_token") : null;
    const wsUrl = `${protocol}//${host}/ws/auction/${item.id}${token ? `?token=${token}` : ""}`;

    try {
      const ws = new WebSocket(wsUrl);
      socketRef.current = ws;

      ws.onmessage = (event) => {
        try {
          const payload = JSON.parse(event.data);

          if (payload.type === "bid_placed") {
            const newBid = payload.data as Bid;
            setCurrentHighest(newBid);
            setBidsCount((c) => c + 1);
            setBidsList((prev) => [newBid, ...prev]);

            if (user && newBid.bidder_id !== user.id) {
              setOutbidAlert(`You were outbid: ${formatCurrency(newBid.bid_amount)}`);
            }
          } else if (payload.type === "outbid") {
            setOutbidAlert(payload.message || "You have been outbid!");
          }
        } catch {
          // ignore
        }
      };
    } catch {
      // WS connection fallback
    }

    return () => {
      if (socketRef.current) {
        socketRef.current.close();
      }
    };
  }, [item.id, user]);

  const currentPriceNum = currentHighest
    ? parseFloat(currentHighest.bid_amount)
    : parseFloat(item.price);

  const increment500 = currentPriceNum + 500;
  const increment1000 = currentPriceNum + 1000;
  const increment5000 = currentPriceNum + 5000;

  const handlePlaceBid = async (amount: number) => {
    if (!user) {
      window.location.href = "/login";
      return;
    }

    setIsSubmitting(true);
    setErrorAlert(null);
    setSuccessAlert(null);
    setOutbidAlert(null);

    try {
      const res = await api.placeBid(item.id, amount);
      trackBidPlaced(item.id, item.name, amount);
      setSuccessAlert(`Bid placed: ${formatCurrency(amount)}`);
      setCurrentHighest(res.bid);
      setBidsCount((c) => c + 1);
      setCustomAmount("");
      onBidPlaced?.();
    } catch (err: any) {
      setErrorAlert(err.message || "Failed to place bid");
    } finally {
      setIsSubmitting(false);
    }
  };

  const timerInfo = formatTimeRemaining(timeLeft);
  const isAuctionClosed = timerInfo.isExpired || item.status === "sold";

  return (
    <div className="bg-white border border-[#E4E4DC] p-5 flex flex-col gap-5 text-[#141413]">
      {/* Header */}
      <div className="flex items-center justify-between pb-3 border-b border-[#E4E4DC] text-xs">
        <span className="font-bold flex items-center gap-1.5">
          <span className="w-2 h-2 rounded-full bg-[#B83A24] animate-pulse" />
          Live Bidding
        </span>

        <span className="text-[#706F6A] font-mono flex items-center gap-1">
          <Clock className="w-3.5 h-3.5 text-[#706F6A]" />
          {timerInfo.formatted}
        </span>
      </div>

      {/* Highest Bid */}
      <div className="flex items-baseline justify-between bg-[#F9F9F6] p-4 border border-[#E4E4DC]">
        <div>
          <span className="text-[11px] text-[#706F6A] block">
            {currentHighest ? "Current Bid" : "Starting Bid"}
          </span>
          <span className="font-mono text-2xl font-bold text-[#141413]">
            {formatCurrency(currentPriceNum)}
          </span>
        </div>
        <span className="text-xs text-[#706F6A]">
          {bidsCount} {bidsCount === 1 ? "bid" : "bids"}
        </span>
      </div>

      {/* Alerts */}
      {outbidAlert && (
        <div className="p-2.5 bg-amber-50 border border-amber-200 text-amber-800 text-xs flex items-center gap-2">
          <AlertTriangle className="w-4 h-4 flex-shrink-0" />
          <span>{outbidAlert}</span>
        </div>
      )}

      {successAlert && (
        <div className="p-2.5 bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs flex items-center gap-2">
          <CheckCircle2 className="w-4 h-4 flex-shrink-0 text-emerald-600" />
          <span>{successAlert}</span>
        </div>
      )}

      {errorAlert && (
        <div className="p-2.5 bg-red-50 border border-red-200 text-red-800 text-xs flex items-center gap-2">
          <AlertTriangle className="w-4 h-4 flex-shrink-0" />
          <span>{errorAlert}</span>
        </div>
      )}

      {/* Quick Increments */}
      {!isAuctionClosed ? (
        <div className="space-y-3">
          <div className="grid grid-cols-3 gap-2 text-xs">
            <button
              onClick={() => handlePlaceBid(increment500)}
              disabled={isSubmitting}
              className="py-2 px-1 border border-[#E4E4DC] hover:border-[#141413] bg-white text-center font-medium transition-colors disabled:opacity-50"
            >
              +NPR 500
            </button>

            <button
              onClick={() => handlePlaceBid(increment1000)}
              disabled={isSubmitting}
              className="py-2 px-1 border border-[#E4E4DC] hover:border-[#141413] bg-white text-center font-medium transition-colors disabled:opacity-50"
            >
              +NPR 1,000
            </button>

            <button
              onClick={() => handlePlaceBid(increment5000)}
              disabled={isSubmitting}
              className="py-2 px-1 border border-[#E4E4DC] hover:border-[#141413] bg-white text-center font-medium transition-colors disabled:opacity-50"
            >
              +NPR 5,000
            </button>
          </div>

          {/* Custom Bid */}
          <div className="flex gap-2 text-xs">
            <input
              type="number"
              placeholder={`Min ${currentPriceNum + 1}`}
              value={customAmount}
              onChange={(e) => setCustomAmount(e.target.value)}
              className="flex-1 px-3 py-2 bg-white border border-[#E4E4DC] focus:outline-none focus:border-[#141413] text-[#141413] font-mono"
            />
            <button
              onClick={() => {
                const val = parseFloat(customAmount);
                if (!isNaN(val)) handlePlaceBid(val);
              }}
              disabled={isSubmitting || !customAmount}
              className="px-5 py-2 bg-[#141413] hover:bg-[#2B2B28] text-white uppercase tracking-wider text-xs font-medium transition-colors disabled:opacity-50"
            >
              Bid
            </button>
          </div>
        </div>
      ) : (
        <div className="py-4 text-center text-xs text-[#706F6A] bg-[#F2F2EC]">
          This auction has ended.
        </div>
      )}

      {/* Bids List */}
      {bidsList.length > 0 && (
        <div className="pt-2 border-t border-[#E4E4DC]">
          <span className="text-[11px] font-bold text-[#706F6A] block mb-2">
            Recent Bids
          </span>
          <div className="space-y-1.5 max-h-36 overflow-y-auto text-xs font-mono pr-1">
            {bidsList.slice(0, 5).map((bid, i) => (
              <div
                key={bid.id || i}
                className="flex items-center justify-between py-1 border-b border-[#F2F2EC] text-[#706F6A]"
              >
                <span>Bidder #{bid.bidder_id ? bid.bidder_id.substring(0, 4) : "anon"}</span>
                <span className="font-bold text-[#141413]">
                  {formatCurrency(bid.bid_amount)}
                </span>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};
