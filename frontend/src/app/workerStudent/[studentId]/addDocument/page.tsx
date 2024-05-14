'use client'

import useSWR from "swr";
import {useError} from "@/customHooks/useError";
import React, {FormEvent, useRef, useState} from "react";
import {Spinner, Table} from "react-bootstrap";
import style from "@/styles/table/table.module.css";
import DownloadFile from "@/elements/downloadFile/downloadFile";
import ApplicationModal from "@/elements/modalApplication/modalApplication";
import {useParams} from "next/navigation";
import Form from "react-bootstrap/Form";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";

export default function Home() {
    const {studentId} = useParams<{studentId : string}>()
    const formRef = useRef<HTMLFormElement>(null)
    const {mutate} = useSWR("/v1/document")
    const {setError, setErrorCode,
        setShow,alert } = useError()

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)
        formData.set("student_id", studentId)

        const response = await fetch('http://localhost:80/api/v1/worker/document', {
            method: 'POST',
            body: formData,
        })
        if (!response.ok) {
            const text = await response.text().then(text => text)
            setShow(true)
            setError(text)
            setErrorCode(response.status)
            return
        }

        await mutate()
        if (formRef && formRef.current)
            formRef.current.reset()
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Form method={"POST"} onSubmit={onSubmit} ref={formRef}>
                <Input Name={"name"} Id={"NameId"} Type={"text"}  LabelHolder={"Имя документа"} PlaceHolder={"Введите имя документа"} ClassName={"mt-3"}/>
                <Input Name={"type"} Id={"TypeId"} Type={"text"}  LabelHolder={"Тип документа"} PlaceHolder={"Введите тип документа"} ClassName={"mt-3"}/>
                <Input Name={"file"} Id={"FileId"} Type={"file"} Allowed={"application/pdf"} LabelHolder={"Файл"} PlaceHolder={"Файл документа"} ClassName={"mt-3"}/>
                <Button type={"submit"} className={"mt-3"}>Добавить</Button>
            </Form>
            {alert}
        </main>
    );
}