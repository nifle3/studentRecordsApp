import { Inter } from "next/font/google";
import {ReactNode} from "react";
import "@/styles/global.css"
import {useRouter} from "next/navigation";

const inter = Inter({ subsets: ["latin"] });
export default function RootLayout({
  children,
}: Readonly<{
  children: ReactNode;
}>) {
    return (
    <html lang="ru">
      <body className={inter.className}>{children}</body>
    </html>
  );
}
