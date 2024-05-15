import type { Metadata } from "next";
import { Inter } from "next/font/google";
import {ReactNode} from "react";
import "@/styles/global.css"
import "@/styles/layoutPage/layoutPage.module.css"
import style from "@/styles/layoutPage/layoutPage.module.css";
import Header from "@/elements/header/header";
import StudentProfile from "@/app/workerStudent/[studentId]/studentProfile";
import 'bootstrap/dist/css/bootstrap.min.css';
import Back from "@/app/workerStudent/[studentId]/back";
import Tabs from "@/app/workerStudent/[studentId]/tabs";
import Footer from "@/elements/footer/footer";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Студент",
    description: "Страница студента",
};

export default function WorkeStudentrLayout({ children }: Readonly<{ children: ReactNode; }>) {
    return (
        <html lang="ru">
            <body className={inter.className}>
                <div className={style.mainWrapper}>
                    <div className={style.content}>

                    <Header/>
                    <Back/>
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
