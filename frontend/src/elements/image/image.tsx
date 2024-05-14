'use client'

import React, {useEffect, useRef} from "react";
import style from "@/styles/image/image.module.css"

export interface ImageProps {
    fetchUri : string
}

export default function Image({fetchUri} : ImageProps) {
    const ref = useRef<HTMLImageElement>(null)

    useEffect(() => {
        fetch(fetchUri, {
            method: "GET",
        }).then(response => response.blob()).then(
            blob => {
                if (!ref || !ref.current) {
                    throw "REF IS NULL"
                }

                ref.current.src = URL.createObjectURL(blob)
            }
        )
    }, [ref])

    return (
        <img ref={ref} src={""} alt={"Image"} className={style.image}/>
    )
}