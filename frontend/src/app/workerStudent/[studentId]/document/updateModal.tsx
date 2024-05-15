'use client'

import {Col, Modal, Row} from "react-bootstrap";
import Input from "@/elements/input/input";
import Button from "react-bootstrap/Button";
import React, {FormEvent} from "react";
import {KeyedMutator} from "swr";
import {useError} from "@/customHooks/useError";
import {retry} from "next/dist/compiled/@next/font/dist/google/retry";
import {useParams} from "next/navigation";


export default function UpdateModal({Data, onClose, mutate, Show}
                                        : { Data: any; onClose : () => void; mutate :  KeyedMutator<any>; Show : boolean}) {
    const {alert, setError, setShow
        , setErrorCode} = useError()

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)

        const response = await fetch('https://localhost:443/api/v1/worker/document/' + Data.id, {
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
                    Обновление документа
                </Modal.Title>
            </Modal.Header>
            <form onSubmit={onSubmit}>
                <Modal.Body>
                    <Input Name={"name"} Id={"NameId"} Type={"text"}  LabelHolder={"Имя документа"} PlaceHolder={"Введите имя документа"}
                           ClassName={"mt-3"} Value={Data.name}/>
                    <Input Name={"type"} Id={"TypeId"} Type={"text"}  LabelHolder={"Тип документа"} PlaceHolder={"Введите тип документа"}
                           ClassName={"mt-3"} Value={Data.type}/>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant={"secondary"} onClick={onClose}>Отмена</Button>
                    <Button variant={"primary"} type={"submit"}>Обновить</Button>
                </Modal.Footer>
            </form>
            {alert}
        </Modal>
    )
}