import type { Metadata } from "next";
import Script from "next/script";
import "./globals.css";
import { AuthProvider } from "../context/AuthContext";
import { Navbar } from "../components/Navbar";
import { Footer } from "../components/Footer";
import { GA_TRACKING_ID } from "../lib/gtag";

export const metadata: Metadata = {
  title: "AUCTION // STUDIO — Live Antiquities & Fine Horology Floor",
  description:
    "Field office for synchronized live bidding on rare horology, contemporary design, and antiquities. Concurrently executed over Go WebSocket channels with instant Khalti digital settlement.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        {/* Google Analytics 4 Script */}
        <Script
          strategy="afterInteractive"
          src={`https://www.googletagmanager.com/gtag/js?id=${GA_TRACKING_ID}`}
        />
        <Script
          id="google-analytics"
          strategy="afterInteractive"
          dangerouslySetInnerHTML={{
            __html: `
              window.dataLayer = window.dataLayer || [];
              function gtag(){dataLayer.push(arguments);}
              gtag('js', new Date());
              gtag('config', '${GA_TRACKING_ID}', {
                page_path: window.location.pathname,
              });
            `,
          }}
        />
      </head>
      <body className="min-h-screen flex flex-col bg-[#F9F9F6] text-[#141413] antialiased selection:bg-[#141413] selection:text-[#F9F9F6] font-sans">
        <AuthProvider>
          <Navbar />
          <main className="flex-1">{children}</main>
          <Footer />
        </AuthProvider>
      </body>
    </html>
  );
}
