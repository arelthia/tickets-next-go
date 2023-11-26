import Footing from '@/components/Footing'
import Heading from '@/components/Heading'
import type { Metadata } from 'next'

import { Inter } from 'next/font/google'
import "@/app/app.css";

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>        <div className="wrapper">
        <Heading title="Support Dashboard" />
        <main>
          {children}
        </main>
        <Footing />
      </div></body>
    </html>
  )
}
