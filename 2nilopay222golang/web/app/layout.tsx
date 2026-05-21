import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './styles/globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Nilo Commerce - Direct Bill Payments for Ethiopia',
  description:
    'Secure cross-border bill payment platform. Pay for school tuition, healthcare, and groceries directly to merchants in Ethiopia from anywhere in the world.',
  keywords: 'Ethiopia, diaspora, bill payment, tuition, healthcare, supermarket, cross-border',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  );
}
