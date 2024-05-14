import type { Metadata } from "next";
import { Inter } from "next/font/google";
import {ReactNode} from "react";
import "@/styles/global.css"
import "@/styles/layoutPage/layoutPage.module.css"
import style from "@/styles/layoutPage/layoutPage.module.css";
import Header from "@/elements/header/header";
import StudentProfile from "@/app/student/studentProfile";
import 'bootstrap/dist/css/bootstrap.min.css';
import Tabs from "@/app/student/tabs";
import Footer from "@/elements/footer/footer";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Студент",
    description: "Страница студента",
};

export default function RootLayout({ children }: Readonly<{ children: ReactNode; }>) {
    return (
        <html lang="ru">
            <body className={inter.className}>
            <div className={style.mainWrapper}>
                <div className={style.content}>
                    <Header/>
                    <div className={style.profileTextWrapper}>
                        <h3>Профиль</h3>
                    </div>
                    <StudentProfile/>
                    <Tabs/>
                    {children}
                </div>
                <Footer/>
            </div>
            </body>
        </html>
    );
}
