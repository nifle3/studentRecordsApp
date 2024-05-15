'use client'

import useSWR from "swr";
import {useError} from "@/customHooks/useError";
import {FormEvent, useRef} from "react";
import Form from "react-bootstrap/Form";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import {Col, Row} from "react-bootstrap";
import "@/styles/addStudent/addStudent.css"

export default function Home() {
    const { mutate} = useSWR("/v1/students")
    const {alert, setError, setShow
        , setErrorCode} = useError()
    const formRef = useRef<HTMLFormElement>(null)

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('https://localhost:443/api/v1/worker/student', {
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
                <Input Name={"name"} Id={"FirstNameId"}
                       Type={"text"} LabelHolder={"Имя"} PlaceHolder={"Введите имя"} ClassName={"mt-2"}/>
                <Input Name={"last_name"} Id={"LastNameId"}
                       Type={"text"} LabelHolder={"Фамилия"} PlaceHolder={"Введите фамилилю"} ClassName={"mt-2"}/>
                <Input Name={"surname"} Id={"SurNameId"}
                       Type={"text"} LabelHolder={"Отчество"} PlaceHolder={"Введите отчество"} ClassName={"mt-2"}/>
                <Input Name={"email"} Id={"EmailId"}
                       Type={"email"} LabelHolder={"Почта"} PlaceHolder={"Введите почту"} ClassName={"mt-2"}/>
                <Input Name={"password"} Id={"PasswordId"}
                       Type={"password"} LabelHolder={"Пароль"} PlaceHolder={"Введите пароль"} ClassName={"mt-2"}/>
                <Input Name={"birthdate"} Id={"birthdateId"}
                       Type={"date"} LabelHolder={"День рождения"} PlaceHolder={"Введите день рождения"} ClassName={"mt-2"}/>
                <Row>
                    <Col>
                        <Input Name={"passport_number"} Id={"passport_numberId"}
                               Type={"number"} LabelHolder={"Номер паспорта"} PlaceHolder={"Введите номер паспорта"}
                               ClassName={"mt-2"}/>
                    </Col>
                    <Col>
                        <Input Name={"passport_seria"} Id={"passport_seriaШв"}
                               Type={"number"} LabelHolder={"Серия паспорта"} PlaceHolder={"Введите серию паспорта"} ClassName={"mt-2"}/>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <Input Name={"course"} Id={"CourseId"}
                               Type={"number"} LabelHolder={"Курс"} PlaceHolder={"Введите курс"} ClassName={"mt-2"}
                               Value={1}/>
                    </Col>
                    <Col>
                        <Input Name={"group"} Id={"GroupId"}
                               Type={"number"} LabelHolder={"Группа"} PlaceHolder={"Введите группу"} ClassName={"mt-2"}/>
                    </Col>
                </Row>
                <Input Name={"specialization"} Id={"specializationId"}
                       Type={"text"} LabelHolder={"Специализация"} PlaceHolder={"Введите специализацию"} ClassName={"mt-2"}/>

                <Row>
                    <Col>
                        <Input Name={"country"} Id={"countryId"} Type={"text"} LabelHolder={"Страна прописки"}
                       PlaceHolder={"Введите страну прописки"} ClassName={"mt-2"} Value={"Россия"}/>
                    </Col>
                    <Col>
                        <Input Name={"city"} Id={"cityId"} Type={"text"} LabelHolder={"Город прописки"}
                         PlaceHolder={"Введите город прописки"} ClassName={"mt-2"} Value={"Казань"}/>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <Input Name={"street"} Id={"streetId"} Type={"text"} LabelHolder={"Улица прописки"}
                       PlaceHolder={"Введите улицу прописки"} ClassName={"mt-2"}/>
                    </Col>
                    <Col>
                        <Input Name={"house"} Id={"houseШв"} Type={"text"} LabelHolder={"Дом прописки"}
                       PlaceHolder={"Введите дом прописки"} ClassName={"mt-2"}/>
                    </Col>
                    <Col>
                        <Input Name={"apartment"} Id={"apartmentId"} Type={"text"} LabelHolder={"Квартира прописки"}
                       PlaceHolder={"Введите квартиру прописки"} ClassName={"mt-2"} Value={0}/>
                    </Col>
                </Row>
                <Input Name={"image"} Id={"imageId"} Type={"file"}  Allowed={"image/jpeg"} LabelHolder={"Фотография"}
                       PlaceHolder={"Загрузите фотографию"} ClassName={"mt-2"}/>
                <Row>
                    <Col>
                        <Input Name={"phone1"} Id={"phone1Id"} Type={"tel"} LabelHolder={"Номер телефона"}
                               PlaceHolder={"Введите номер телефона"} ClassName={"mt-2"}/>
                    </Col>
                    <Col>
                        <Input Name={"description1"} Id={"description1Id"} Type={"text"} LabelHolder={"Описание номер телефона"}
                               PlaceHolder={"Введите описание номера телефона"} ClassName={"mt-2"}/>
                    </Col>
                </Row>
                <Button type={"submit"} className={"mt-3"}>Добавить</Button>
            </Form>
            {alert}
        </main>
    )
}