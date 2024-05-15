'use client'

import {useRouter} from "next/navigation";
import style from "@/styles/back/back.module.css"

export default function Back() {
    const router = useRouter()

    return (
        <h1 onClick={() => router.push("/worker")} className={style.backText}>
            Назад
            <svg width="50" height="50" viewBox="0 0 50 50" fill="none" xmlns="http://www.w3.org/2000/svg">
                <g filter="url(#filter0_d_2_11)">
                    <path fill-rule="evenodd" clip-rule="evenodd"
                          d="M6.25 25C6.25 24.4475 6.4695 23.9175 6.86019 23.5269L21.4435 8.94354C22.2571 8.13 23.5763 8.13 24.3898 8.94354C25.2033 9.75708 25.2033 11.0762 24.3898 11.8898L13.3629 22.9167L41.6667 22.9167C42.8173 22.9167 43.75 23.8494 43.75 25C43.75 26.1506 42.8173 27.0833 41.6667 27.0833H13.3629L24.3898 38.1102C25.2033 38.9238 25.2033 40.2429 24.3898 41.0565C23.5763 41.8701 22.2571 41.8701 21.4435 41.0565L6.86019 26.4731C6.4695 26.0825 6.25 25.5525 6.25 25Z"
                          fill="black"/>
                </g>
                <defs>
                    <filter id="filter0_d_2_11" x="-4" y="0" width="58" height="58" filterUnits="userSpaceOnUse"
                            color-interpolation-filters="sRGB">
                        <feFlood flood-opacity="0" result="BackgroundImageFix"/>
                        <feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
                                       result="hardAlpha"/>
                        <feOffset dy="4"/>
                        <feGaussianBlur stdDeviation="2"/>
                        <feComposite in2="hardAlpha" operator="out"/>
                        <feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
                        <feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_2_11"/>
                        <feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_2_11" result="shape"/>
                    </filter>
                </defs>
            </svg>

        </h1>
    )
}