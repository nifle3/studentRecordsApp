'use client'

import {FormEvent, useRef} from "react";
import Form from "react-bootstrap/Form";
import {useError} from "@/customHooks/useError";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import useSWR from "swr";

export default function Home() {
    const {setShow, setError, setErrorCode,
    alert} = useError()
    const ref = useRef<HTMLFormElement>(null)
    const {mutate} = useSWR("/v1/student/application")
    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('https://localhost:443/api/v1/student/application', {
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
        
        if (ref && ref.current)
            ref.current.reset()

        await mutate()
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Form method={"POST"} onSubmit={onSubmit} ref={ref}>
                <Input Id={"ContactInfoId"} LabelHolder={"Контактная информация"} Name={"contact_info"}
                       PlaceHolder={"Введите контактную информацию"}
                       Type={"text"} ClassName={"mt-3"}/>
                <Input Id={"TextId"} LabelHolder={"Текст"} Name={"text"} PlaceHolder={"Введтие текст"}
                       Type={"text"} ClassName={"mt-3"}/>
                <Input Id={"NameId"} LabelHolder={"Название"} Name={"name"} PlaceHolder={"Введите название заявки"}
                       Type={"text"} ClassName={"mt-3"}/>
                <Input Id={"FileId"} LabelHolder={"Файл"} Name={"file"} PlaceHolder={"Выберите файл"}
                       Type={"file"} Allowed={"application/pdf"} ClassName={"mt-3"}/>
                <Button type={"submit"} className={"mt-3 "}>Добавить</Button>
            </Form>
            {alert}
        </main>
    );
}