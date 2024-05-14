'use client'

import {useRef} from "react";

export interface DownloadFileProps {
    ClassName? : string
    Fetch : string
    FileName : string
}

export default function DownloadFile({ClassName, Fetch, FileName} : DownloadFileProps) {
    async function onClick() {
        const blob = await fetch(Fetch, {
            method: "GET"
        }).then(response => response.blob()).then(blob => {
            const url = URL.createObjectURL(new Blob([blob]))
            const link = document.createElement('a')
            link.href = url
            link.setAttribute('download', FileName)
            link.setAttribute('target', '_blank')
            link.click()
            URL.revokeObjectURL(url)
        })
    }


    return (
        <span className={ClassName}  onClick={onClick}>
            Скачать
        </span>
    )
}