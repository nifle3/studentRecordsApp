'use client'

import useSWR from "swr";
import {useParams, useRouter} from "next/navigation";
import {Col, Row, Spinner} from "react-bootstrap";
import React, {FormEvent} from "react";
import Form from "react-bootstrap/Form";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import {useError} from "@/customHooks/useError";
import Image from "@/elements/image/image";

export default function StudentProfile() {
    const {alert, setError, setShow
        , setErrorCode} = useError()

    const fetcher = async () => await fetch("https://localhost:443/api/api/v1/student", {
        method: "GET",
    }) .then(response => response.json())
    const {data, isLoading, mutate} = useSWR("/v1/student", fetcher)
    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('https://localhost:443/api/v1/student', {
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
        <Form method={"POST"} onSubmit={onSubmit} className={"mb-3"}>
            <Row>
                <Col>
                    <Image fetchUri={"https://localhost:443/api/v1/student/photo"}/>
                </Col>
                <Col>
                    <Input Name={"name"} Id={"FirstNameId"}
                           ReadOnly
                           Type={"text"} LabelHolder={"Имя"} PlaceHolder={"Введите имя"} ClassName={"mt-2"} Value={data.first_name}/>
                    <Input Name={"last_name"} Id={"LastNameId"} Value={data.last_name}
                           ReadOnly
                           Type={"text"} LabelHolder={"Фамилия"} PlaceHolder={"Введите фамилилю"} ClassName={"mt-2"}/>
                    <Input Name={"surname"} Id={"SurNameId"} Value={data.surname}
                           ReadOnly
                           Type={"text"} LabelHolder={"Отчество"} PlaceHolder={"Введите отчество"} ClassName={"mt-2"}/>
                    <Input Name={"email"} Id={"EmailId"} Value={data.email}
                           ReadOnly
                           Type={"email"} LabelHolder={"Почта"} PlaceHolder={"Введите почту"} ClassName={"mt-2"}/>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Input Name={"passport_number"} Id={"passport_numberId"} Value={data.passport_number}
                           ReadOnly
                           Type={"number"} LabelHolder={"Номер паспорта"} PlaceHolder={"Введите номер паспорта"}
                           ClassName={"mt-2"}/>
                </Col>
                <Col>
                    <Input Name={"passport_seria"} Id={"passport_seriaId"} Value={data.passport_seria}
                           ReadOnly
                           Type={"number"} LabelHolder={"Серия паспорта"} PlaceHolder={"Введите серию паспорта"} ClassName={"mt-2"}/>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Input Name={"course"} Id={"CourseId"} Value={data.course}
                           ReadOnly
                           Type={"number"} LabelHolder={"Курс"} PlaceHolder={"Введите курс"} ClassName={"mt-2"}/>
                </Col>
                <Col>
                    <Input Name={"group"} Id={"GroupId"} Value={data.group}
                           ReadOnly
                           Type={"number"} LabelHolder={"Группа"} PlaceHolder={"Введите группу"} ClassName={"mt-2"}/>
                </Col>
            </Row>
            <Input Name={"specialization"} Id={"specializationId"} Value={data.specialization}
                   ReadOnly
                   Type={"text"} LabelHolder={"Специализация"} PlaceHolder={"Введите специализацию"} ClassName={"mt-2"}/>

            <Row>
                <Col>
                    <Input Name={"country"} Id={"countryId"} Type={"text"} LabelHolder={"Страна прописки"}
                           ReadOnly
                           PlaceHolder={"Введите страну прописки"} ClassName={"mt-2"}
                    Value={data.country}/>
                </Col>
                <Col>
                    <Input Name={"city"} Id={"cityId"} Type={"text"} LabelHolder={"Город прописки"}
                           ReadOnly
                           PlaceHolder={"Введите город прописки"} ClassName={"mt-2"} Value={data.city}/>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Input Name={"street"} Id={"streetId"} Type={"text"} LabelHolder={"Улица прописки"}
                           ReadOnly
                           PlaceHolder={"Введите улицу прописки"} ClassName={"mt-2"} Value={data.street}/>
                </Col>
                <Col>
                    <Input Name={"house"} Id={"houseШв"} Type={"text"} LabelHolder={"Дом прописки"}
                           PlaceHolder={"Введите дом прописки"} ClassName={"mt-2"} Value={data.house_number}
                           ReadOnly/>
                </Col>
                <Col>
                    <Input Name={"apartment"} Id={"apartmentId"} Type={"text"} LabelHolder={"Квартира прописки"}
                           PlaceHolder={"Введите квартиру прописки"} ClassName={"mt-2"} Value={data.apartment_number}
                           ReadOnly/>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Input Name={"phone1"} Id={"phone1Id"} Type={"tel"} LabelHolder={"Номер телефона"}
                           PlaceHolder={"Введите номер телефона"} ClassName={"mt-2"} Value={data.phoneNumbers[0].phone}
                           ReadOnly/>
                </Col>
                <Col>
                    <Input Name={"description1"} Id={"description1Id"} Type={"text"} LabelHolder={"Описание номер телефона"}
                           PlaceHolder={"Введите описание номера телефона"} ClassName={"mt-2"} Value={data.phoneNumbers[0].description}
                    ReadOnly/>
                </Col>
            </Row>
        </Form>
            {alert}
        </>
    )
}