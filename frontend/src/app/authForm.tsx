'use client'

import style from "@/styles/login/page.module.css";
import SelectRole from "@/elements/selectRole/SelectRole";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import {FormEvent, useLayoutEffect, useState} from "react";
import {useRouter} from "next/navigation";
import CustomAlert from "@/elements/customAlert/customAlert";

export default function AuthForm() {
    const router = useRouter()
    const [showError, setShowError] = useState<boolean>(false)
    const [errorCode, setErrorCode] = useState<number>(200)
    const [error, setError] = useState<string>("")

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)
        const response = await fetch('https://localhost:443/api/v1/auth', {
            method: 'POST',
            body: formData,
        })

        if (!response.ok) {
            const text = await response.text().then(text => text)
            setShowError(true)
            setError(text)
            setErrorCode(response.status)
            return
        }

        const role = formData.get("role")
        if (role == null) {
            const text = await response.text().then(text => text)
            setShowError(true)
            setError("Роль пустая")
            setErrorCode(400)
            return
        }


        if (role == "Админ") {
            router.push("/admin")
        }

        if (role == "Работник") {
            router.push("/worker")
        }

        if (role == "Студент") {
            router.push("/student")
        }
    }

    return (
        <>
            <Form className={style.login} onSubmit={onSubmit} method={"POST"}>
                <SelectRole Roles={["Админ", "Работник", "Студент"]} Id={"RoleSelectId"} Name={"role"}
                            ClassName={style.input}/>
                <Input  Type={"email"} LabelHolder={"Почта:"}
                        Name={"login"} PlaceHolder={"Введите почту"} Id={"EmailInput"} ClassName={style.input + " mt-2"}/>
                <Input Name={"password"} PlaceHolder={"Введите пароль"}
                       Type={"password"} LabelHolder={"Пароль:"} Id={"PasswordInput"} ClassName={style.input + " mt-2"}/>
                <Button type={"submit"} className={"mt-3 " + style.input} variant={"primary"}>
                    Войти
                </Button>
            </Form>
            <CustomAlert IsShow={showError} SetIsShow={setShowError} Error={error} ErrorCode={errorCode}/>
        </>
    )
}