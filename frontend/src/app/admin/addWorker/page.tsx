'use client'

import useSWR from "swr";
import {useError} from "@/customHooks/useError";
import Form from "react-bootstrap/Form";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import {FormEvent, useRef} from "react";
import SelectRole from "@/elements/selectRole/SelectRole";


// TODO сделать уведомление про добавление
export default function Home() {
    const { mutate} = useSWR("/v1/workers")
    const {alert, setError, setShow
        , setErrorCode} = useError()
    const formRef = useRef<HTMLFormElement>(null)

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('http://localhost:80/api/v1/admin/worker', {
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

        if (formRef && formRef.current)
            formRef.current.reset()

        await mutate()
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Form onSubmit={onSubmit} method={"POST"} ref={formRef}>
                <Input Name={"first_name"} Id={"FirstNameId"}
                       Type={"text"} LabelHolder={"Имя"} PlaceHolder={"Введите имя"} ClassName={"mt-2"}/>
                <Input Name={"last_name"} Id={"LastNameId"}
                       Type={"text"} LabelHolder={"Фамилия"} PlaceHolder={"Введите фамилилю"} ClassName={"mt-2"}/>
                <Input Name={"surname"} Id={"SurNameId"}
                       Type={"text"} LabelHolder={"Отчество"} PlaceHolder={"Введите отчество"} ClassName={"mt-2"}/>
                <Input Name={"email"} Id={"EmailId"}
                       Type={"email"} LabelHolder={"Почта"} PlaceHolder={"Введите почту"} ClassName={"mt-2"}/>
                <Input Name={"password"} Id={"PasswordId"}
                       Type={"password"} LabelHolder={"Пароль"} PlaceHolder={"Введите пароль"} ClassName={"mt-2"}/>
                <SelectRole Roles={["Админ", "Работник"]} Id={"RoleSelectId"} Name={"role"} Title={"Роль"}/>
                <Button type={"submit"} className={"mt-3"}>Добавить</Button>
            </Form>
            {alert}
        </main>
    )
}