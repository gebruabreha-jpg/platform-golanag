import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';
import { Toaster } from 'react-hot-toast';
import { AuthProvider } from '@/lib/authContext';
import { QueryProvider } from '@/lib/queryClient';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'ConnectMe - Ethiopian & Diaspora Community Platform',
  description:
    'Connect with Ethiopians worldwide. Peer-to-peer shipping, marketplace, housing, scholarships, jobs, currency exchange, and community networking.',
  keywords:
    'Ethiopia, diaspora, community, shipping, marketplace, housing, scholarships, jobs, remittance',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <QueryProvider>
          <AuthProvider>
            {children}
            <Toaster
              position="top-right"
              toastOptions={{
                duration: 4000,
                style: {
                  background: '#333',
                  color: '#fff',
                  borderRadius: '8px',
                },
              }}
            />
          </AuthProvider>
        </QueryProvider>
      </body>
    </html>
  );
}
