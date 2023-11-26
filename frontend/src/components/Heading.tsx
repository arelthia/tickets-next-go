'use client';
import { ReactElement } from "react";
import { usePathname } from 'next/navigation';
import Link from "next/link";

type HeadingProps = { title: string };

const Heading = ({ title }: HeadingProps): ReactElement => {
  const pathname = usePathname()

  return (
    <header className="branding">
      <h1 className="site-title">{title}</h1>
      <nav>
        <Link href="/" className={pathname === "/" ? "active" : ""}>Dashboard</Link>
        <Link href="/create-ticket" className={pathname === "/create-ticket" ? "active" : ""}>Create New Ticket</Link>
      </nav>
    </header>
  )
}

export default Heading;