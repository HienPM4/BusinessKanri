import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "BusinessKanri",
  description: "Sales Data Web System",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
