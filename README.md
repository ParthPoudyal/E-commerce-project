# 🏷️ AuctionHub — eBay-Style Live Auction & E-Commerce Marketplace

> **6th Semester E-Commerce Engineering Capstone Project**  
> *Developed by Parth Poudyal & Arun Mainali*

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![Next.js](https://img.shields.io/badge/Next.js-14.2-black?style=flat&logo=next.js)](https://nextjs.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)](https://www.postgresql.org)
[![Khalti](https://img.shields.io/badge/Payment-Khalti%20v2-5C2D91?style=flat)](https://khalti.com)
[![Architecture](https://img.shields.io/badge/Architecture-Clean%20%2F%20Hexagonal-orange?style=flat)](#architecture)

---

## 📖 Project Preview & Goal

**AuctionHub** is a full-stack, enterprise-grade e-commerce marketplace engineered to replicate **eBay's real-time live auctioning practicalities**. The project was conceived to explore and demonstrate how complex, high-concurrency e-commerce systems actually function behind the scenes:
- Real-time bid streaming and synchronized countdown timers across distributed clients.
- Clean Architecture (Hexagonal / Ports and Adapters) maintaining strict separation between pure business domain rules, use case services, and infrastructure delivery mechanisms.
- Production-grade digital wallet integration using **Khalti's v2 Merchant Sandbox**.
- Real-time telemetry via **Google Analytics 4 (GA4)** tracking e-commerce funnels and bidding interactions.
- A **bespoke Hallmark "Studio" design system** (inspired by `nutlope/hallmark` *Field Office, an identity studio*): warm linen canvas (`#F9F9F6`), architectural 1px hairline rules (`#E4E4DC`), high-contrast editorial serifs (`Newsreader`), precision monospace tickers (`JetBrains Mono`), registration marks (`+`), and zero generic "AI-slop" gradients or bubbly badges.

---

## 🏛️ Architecture Overview

The backend strictly adheres to **Clean Architecture Principles**:

```
E-commerce-project/
├── internal/
│   ├── domain/               # Pure business models (User, Item, Bid, Category, Payment, Cart, Wishlist)
│   └── usecase/              # Business logic & repository port interfaces
│       ├── auth/services/    # Registration & login with Bcrypt + JWT
│       ├── bid/services/     # Live bid validation (increments, timer, seller checks)
│       ├── auction/services/ # Auction lifecycle & auto-close logic
│       ├── payment/services/ # Khalti v2 API initiate & lookup verification
│       ├── cart/services/    # Cart management use cases
│       └── wishlist/services/# User wishlist use cases
├── infrastructure/
│   ├── database/             # PostgreSQL connection, AutoMigrate & Seeding
│   ├── models/               # GORM relational persistence models
│   ├── mapper/               # Bidirectional mappers (Domain ↔ GORM Model)
│   ├── repository/           # Concrete GORM repository implementations
│   └── auth/                 # JWT token generation & verification
├── api/
│   ├── handlers/             # Gin HTTP handlers for each resource
│   ├── middleware/           # JWT authentication, CORS, upload handlers
│   ├── ws/                   # Real-time WebSocket hub & live room manager
│   ├── dto/                  # Request / Response schemas & validation
│   └── router.go             # Route configuration
└── web/                      # Next.js 14 App Router, TypeScript & Tailwind CSS
```

---

## ⚡ Core Functional Features

### 1. 🔨 Live eBay-Style Bidding Engine
- **Live Auction Rooms**: Dedicated WebSocket connections (`/ws/auction/:id`) stream bids in real-time to all active viewers.
- **Atomic Bid Validations**: Enforces minimum bid increment (1% of current highest bid or 1 NPR minimum), prevents self-bidding by consignors, and rejects bids once the countdown timer expires.
- **Outbid Telemetry**: Instant outbid alerts pushed directly to displaced bidders over WebSocket.
- **Tabular Numerics**: Bid values and second-by-second countdown timers utilize tabular figures (`font-variant-numeric: tabular-nums`) to prevent layout shifts.

### 2. 🗂️ Categorical Display & Filtering
- Dynamic horizontal category pills with custom iconography across 8 major sectors: *Electronics, Watches & Jewelry, Art & Collectibles, Sneakers & Fashion, Antiques & Vintage, Sports Memorabilia, Automotive & Motors, and Books & Manuscripts*.
- Multi-mode filtering: Toggle between **All Items**, **Live Auctions Only**, and **Buy It Now (Fixed Price)**.
- Real-time client-side sorting by newest listings, price ascending, and price descending.

### 3. 💳 Khalti Digital Wallet Sandbox Integration
- **Full Sandbox API (dev.khalti.com)**:
  - `POST /epayment/initiate/`: Generates payment intent and unique `pidx`.
  - `POST /epayment/lookup/`: Verifies payment settlement and transaction status.
  - Interactive test dialog with pre-filled test numbers, MPIN, and OTP.
  - Automatically marks won auction lots and buy-it-now items as `SOLD` upon verification.

### 4. 📈 Google Analytics 4 (GA4) Telemetry
- Integrated with custom e-commerce and auction telemetry events:
  - `view_item`: Dispatched whenever an auction or item detail page is loaded.
  - `bid_placed`: Records bid value, currency, and item name.
  - `purchase`: Dispatched upon Khalti payment confirmation with transaction PIDX and total.

---

## 🛠️ Prerequisites & Tech Stack

| Component | Technology | Version |
| :--- | :--- | :--- |
| **Backend Language** | Go (Golang) | `1.25+` / `1.26` |
| **HTTP Framework** | Gin Web Framework (`gin-gonic/gin`) | `v1.12` |
| **WebSockets** | Gorilla WebSocket (`gorilla/websocket`) | `v1.5` |
| **Database ORM** | GORM (`gorm.io/gorm`) + PostgreSQL Driver | PostgreSQL 16 |
| **Frontend Framework** | Next.js (App Router) | `14.2+` |
| **Styling** | Tailwind CSS + Lucide React | `3.4+` |
| **Package Manager** | `pnpm` (or `npm`) | `v10+` |

---

## 🚀 Step-by-Step Setup Instructions

### 1. Environment Configuration

Check or create `.env` in the backend root (`E-commerce-project/.env`):

```ini
DB_NAME=ecom_db
DB_PASSWORD=parth grey1
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_DSN=postgres://postgres:parth%20grey1@localhost:5432/ecom_db?sslmode=disable
JWT_SECRET=auctionhub_jwt_secure_super_secret_key_2026
KHALTI_BASE_URL=https://dev.khalti.com/api/v2
KHALTI_SECRET_KEY=test_secret_key_ecom_nepal
SERVER_PORT=8080
FRONTEND_URL=http://localhost:3000
```

### 2. Start PostgreSQL via Docker

In the backend root:
```powershell
cd c:\Users\acer\Desktop\Auction\E-commerce-project
docker compose up -d
```
> *Note: If running PostgreSQL natively without Docker, verify that a database named `ecom_db` is created on port `5432` matching your `.env` credentials.*

### 3. Run the Go Backend API

```powershell
cd c:\Users\acer\Desktop\Auction\E-commerce-project
go run ./api
```
The server will initialize on **`http://localhost:8080`**.
On first startup, the database auto-migrates all tables and seeds categories, demo users, and live auction lots.

### 4. Run the Next.js Frontend

Open a second terminal window:
```powershell
cd c:\Users\acer\Desktop\Auction\E-commerce-project\web
pnpm install
pnpm run dev
```
Open **`http://localhost:3000`** in your browser.

---

## 🎯 Important Actions & Feature Walkthrough

### Action 1: Authenticate with Pre-Seeded Accounts
The database automatically seeds demo accounts for testing:

| Role | Email | Password | Username |
| :--- | :--- | :--- | :--- |
| **Consignor / Seller** | `parth@ecom.np` | `password123` | `parth_consignor` |
| **Active Bidder** | `viraj@ecom.np` | `password123` | `viraj_collector` |

*Or register a new account instantly via `/register`.*

---

### Action 2: Experience Real-Time Live Bidding
1. Open `http://localhost:3000` in **Window A** and sign in as `viraj@ecom.np`.
2. Open an Incognito/Private window in **Window B** at `http://localhost:3000`, sign in as `parth@ecom.np` or register a new user.
3. On both windows, click on the **Featured Live Auction** (e.g. *1972 Vintage Rolex Submariner*).
4. Notice the **Viewer Count** indicator (`• 2 viewing now`) syncs over WebSocket.
5. In Window A, click an increment button (e.g., `+NPR 1,000` or enter a custom bid) and click **Bid**.
6. **Watch Window B**: Without refreshing the page, Window B's current price will immediately update, a new bid item will appear in the **Live Bid History Feed**, and Window B will receive an outbid alert!

---

### Action 3: Test Khalti Sandbox Checkout
Khalti provides an isolated sandbox environment with pre-configured developer credentials:

1. Navigate to an available lot or click **Buy It Now** on any fixed-price item (e.g., *Original Contemporary Acrylic on Canvas* or *Restored Bell Star Helmet*).
2. Click **"Pay with Khalti"**.
3. The interactive **Khalti Payment Gateway Modal** opens:
   - **Test Khalti Phone Numbers**: `9800000000` through `9800000005`
   - **Test MPIN**: `1111`
   - **Test OTP Code**: `987654`
4. Click **Authorize / Pay**.
5. The backend initiates and verifies the transaction via Khalti's lookup API, creates the purchase record, marks the item as **SOLD**, and redirects to your receipt in the **Dashboard**.

---

### Action 4: List a New Auction Lot
1. Click **"+ List Item"** in the top navigation bar (or visit `/create`).
2. Choose **Live Auction**:
   - Set starting bid (e.g., `NPR 25,000`).
   - Select auction duration (1, 3, 5, 7, or 10 days).
   - Pick a category and use one of the quick photo presets or paste an image URL.
3. Click **Publish Listing**.
4. The item immediately becomes available in `/browse?auction=true` and is ready for live bidding!

---

### Action 5: Review Dashboard & Bidding History
1. Click your username avatar in the top right and select **Dashboard & Bids** (`/dashboard`).
2. View your real-time stats:
   - **Bids Placed**: All historical bids with direct links back to the auction lots.
   - **Purchased & Won Items**: All completed settlements with Khalti transaction badges and purchase timestamps.
   - **Total Spent**: Auto-calculated aggregate sum in NPR.

---

## 📡 API Reference Overview

| Method | Endpoint | Access | Purpose |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/auth/register` | Public | Register new user account |
| `POST` | `/api/auth/login` | Public | Sign in and retrieve JWT Bearer token |
| `GET` | `/api/auth/me` | Protected | Get profile of logged-in user |
| `GET` | `/api/categories` | Public | List all 8 marketplace categories |
| `GET` | `/api/items` | Public | Query listings with filter & pagination params |
| `GET` | `/api/items/auctions` | Public | Get all active live bidding lots |
| `GET` | `/api/items/:id` | Public | Get item details + highest bid + time remaining |
| `POST` | `/api/items` | Protected | Create new auction or buy-now listing |
| `POST` | `/api/upload` | Protected | Upload local image file to server `/uploads` |
| `POST` | `/api/bids` | Protected | Place a bid on an active auction |
| `GET` | `/api/bids/item/:id` | Public | Retrieve scrolling bid history for a lot |
| `GET` | `/api/bids/user` | Protected | List all bids placed by the current user |
| `GET` | `/api/cart` | Protected | Get current user's shopping cart |
| `POST` | `/api/cart` | Protected | Add item to cart |
| `DELETE`| `/api/cart/:itemId` | Protected | Remove item from cart |
| `GET` | `/api/wishlist` | Protected | List saved / bookmarked lots |
| `POST` | `/api/wishlist` | Protected | Add lot to wishlist |
| `POST` | `/api/purchases/buy-now/:id` | Protected | Instant direct buy-now execution |
| `POST` | `/api/payments/khalti/initiate`| Protected | Initiate Khalti v2 sandbox checkout intent |
| `POST` | `/api/payments/khalti/verify` | Public | Verify Khalti `pidx` and finalize purchase |
| `GET` | `/ws/auction/:id` | Public/Auth | Live Bidding Room WebSocket connection |

---

## 🎨 Design Philosophy: Hallmark Studio Theme

Inspired by **Hallmark's Studio theme** (`Field Office, an identity studio`), **AuctionHub** discards generic "AI-slop" conventions (floating colored bubbles, rainbow gradients, pill badges with sparkle emojis) in favor of an **editorial, architectural auction catalogue layout**:
- **Warm Linen Canvas**: Unbleached paper and warm linen background (`#F9F9F6`) with stark white catalogue plates (`#FFFFFF`) and muted warm grey details (`#F2F2EC`).
- **Hairline Architectural Grids**: Sharp 1px dividing borders (`border-[#E4E4DC]` and `border-[#141413]`) with registration crosshairs (`+`) reminiscent of Swiss graphic design and Sotheby's auction catalogues.
- **Editorial Typography**: High-contrast serif headlines (`Newsreader` / `Playfair Display`) paired with precision monospace tickers and lot numbers (`JetBrains Mono` / `IBM Plex Mono`).
- **Gavel Ledger Floor**: Live bidding presented as a monospace gavel ledger tape with bidder paddle numbers (`PADDLE #402`), rapid increment controls, and instant Khalti sandbox settlement certificates.

---

## 📄 License
This project was developed for educational purposes as an E-Commerce Engineering Capstone Project.
