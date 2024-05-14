'use client'


import useSWR from "swr";
import Form from "react-bootstrap/Form"
import {Col, Row, Spinner} from "react-bootstrap"
import Input from "@/elements/input/input"
import Button from "react-bootstrap/Button";
import {FormEvent} from "react";
import {useError} from "@/customHooks/useError";

const fetcher = async () => fetch("http://localhost:80/api/v1/admin", {
    method: "GET",
}).then(response => response.json())

export default function AdminSelfInfo() {
    const {data, isLoading, mutate} = useSWR("/v1/admin", fetcher)
    const {alert, setError, setShow
        , setErrorCode} = useError()

    if (isLoading) {
        return (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('http://localhost:80/api/v1/admin', {
            method: 'PATCH',
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
    }


    return (
        <>
            <Form onSubmit={onSubmit} method={"POST"}>
                <Row>
                    <Col>
                        <Input Name={"first_name"} Id={"FirstNameId"}
                               Type={"text"} LabelHolder={"Имя"} PlaceHolder={"Введите ваше имя"}
                               Value={data.first_name} ClassName={"mt-2"}/>
                        <Input Name={"last_name"} Id={"LastNameId"}
                               Type={"text"} LabelHolder={"Фамилия"} PlaceHolder={"Введите ваше фамилилю"}
                               Value={data.last_name} ClassName={"mt-2"}/>
                    </Col>
                    <Col>
                        <Input Name={"surname"} Id={"SurNameId"}
                               Type={"text"} LabelHolder={"Отчество"} PlaceHolder={"Введите ваше отчество"}
                               Value={data.surname} ClassName={"mt-2"}/>
                        <Input Name={"email"} Id={"EmailId"}
                               Type={"email"} LabelHolder={"Почта"} PlaceHolder={"Введите вашу почту"}
                               Value={data.email} ClassName={"mt-2"}/>
                    </Col>
                </Row>
                <Button type={"submit"} className={"mt-3"}>Обновить</Button>
            </Form>
            {alert}
        </>
    )
}