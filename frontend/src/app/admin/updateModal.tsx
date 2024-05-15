'use client'

import {Col, Modal, Row} from "react-bootstrap";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import React, {FormEvent} from "react";
import {KeyedMutator} from "swr";
import {useError} from "@/customHooks/useError";
import {retry} from "next/dist/compiled/@next/font/dist/google/retry";


export default function UpdateModal({Data, onClose, mutate, Show}
                                        : { Data: any; onClose : () => void; mutate :  KeyedMutator<any>; Show : boolean}) {
    const {alert, setError, setShow
        , setErrorCode} = useError()
    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('https://localhost:443/api/v1/admin/worker/' + Data.id, {
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
        onClose()
    }

    if (Data == null) {
        return
    }

    return (
        <Modal show={Show} onHide={onClose} centered keyboard={false} backdrop={"static"}>
            <Modal.Header>
                <Modal.Title>
                    Обновление пользователя
                </Modal.Title>
            </Modal.Header>
                <form onSubmit={onSubmit}>
                    <Modal.Body>
                        <Row>
                            <Col>
                                <Input Name={"first_name"} Id={"FirstNameId"}
                                       Type={"text"} LabelHolder={"Имя"} PlaceHolder={"Введите ваше имя"}
                                       Value={Data.first_name} ClassName={"mt-2"}/>
                                <Input Name={"last_name"} Id={"LastNameId"}
                                       Type={"text"} LabelHolder={"Фамилия"} PlaceHolder={"Введите ваше фамилилю"}
                                       Value={Data.last_name} ClassName={"mt-2"}/>
                            </Col>
                            <Col>
                                <Input Name={"surname"} Id={"SurNameId"}
                                       Type={"text"} LabelHolder={"Отчество"} PlaceHolder={"Введите ваше отчество"}
                                       Value={Data.surname} ClassName={"mt-2"}/>
                                <Input Name={"email"} Id={"EmailId"}
                                       Type={"email"} LabelHolder={"Почта"} PlaceHolder={"Введите вашу почту"}
                                       Value={Data.email} ClassName={"mt-2"}/>
                            </Col>
                        </Row>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant={"secondary"} onClick={onClose}>Отмена</Button>
                        <Button variant={"primary"} type={"submit"}>Обновить</Button>
                    </Modal.Footer>
                </form>
        </Modal>
    )
}