import type { Metadata } from "next";
import { Inter } from "next/font/google";
import {ReactNode} from "react";
import "@/styles/global.css"
import "@/styles/layoutPage/layoutPage.module.css"
import style from "@/styles/layoutPage/layoutPage.module.css";
import Header from "@/elements/header/header";
import AdminSelfInfo from "@/app/admin/adminSelfInfo";
import Tabs from "@/app/admin/tabs";
import 'bootstrap/dist/css/bootstrap.min.css';
import Footer from "@/elements/footer/footer";


const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Админ",
    description: "Страница админа",
};

export default function RootLayout({children}: Readonly<{ children: ReactNode;
}>) {
    return (
        <html lang="ru">
            <body className={inter.className}>
            <div className={style.mainWrapper}>
                <div className={style.content}>
                    <Header/>
                    <div className={style.profileTextWrapper}>
                        <h3>Профиль</h3>
                    </div>
                    <AdminSelfInfo/>
                    <Tabs ClassName={style.tabs}/>
                    {children}
                </div>
               <Footer/>
            </div>
            </body>
        </html>
    );
}
