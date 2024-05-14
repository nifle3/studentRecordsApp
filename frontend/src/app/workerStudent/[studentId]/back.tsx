'use client'

import {useRouter} from "next/navigation";
import style from "@/styles/back/back.module.css"

export default function Back() {
    const router = useRouter()

    return (
        <h1 onClick={() => router.push("/worker")} className={style.backText}>Назад</h1>
    )
}