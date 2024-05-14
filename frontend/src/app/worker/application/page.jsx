'use client'

import {Spinner, Table} from "react-bootstrap";
import React, {useRef, useState} from "react";
import useSWR from "swr";
import style from "@/styles/table/table.module.css"
import DownloadFile from "@/elements/downloadFile/downloadFile";
import {useError} from "@/customHooks/useError";
import ApplicationModal from "@/elements/modalApplication/modalApplication";

const fetcher = () => fetch("http://localhost:80/api/v1/worker/application", {
    method: "GET",
}).then(response => response.json())

export default function Home() {
    const {data, isLoading, mutate} = useSWR("/v1/worker/application", fetcher)
    const {setError, setErrorCode,
        setShow,alert } = useError()
    const [showModal, setShowModal] = useState(false)
    const [selectId, setSelectId] = useState("")

    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    const onUpdate = (id) => {
        return () => {
            setSelectId(id)
            setShowModal(true)
        }
    }

    const onOK = (id) => {
        return async () => {
            const response = await fetch("http://localhost:80/api/v1/worker/application/close/" + id, {
                method: "PATCH",
            }).then(response => response)
            if (!response.ok) {
                const text = await response.text()
                setErrorCode(response.status)
                setError(text)
                setShow(true)
                return
            }
            await mutate()
            setShowModal(false)
        }
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Table>
                <thead>
                <tr>
                    <th>Название</th>
                    <th>Имя отправителя</th>
                    <th>Группа отправителя</th>
                    <th>Курс отпрвителя</th>
                    <th>Созадана</th>
                    <th>Текст</th>
                    <th>Контактная информация</th>
                    <th>Закрыть</th>
                    <th>Скачать</th>
                </tr>
                </thead>
                <tbody>
                {data && data.map((val, idx) =>
                    <tr key={idx}>
                        <td>{val.name}</td>
                        <td>{val.fio}</td>
                        <td>{val.group}</td>
                        <td>{val.course}</td>
                        <td>{val.created_at.substring(0, 10)}</td>
                        <td>{val.text}</td>
                        <td>{val.contact_info}</td>
                        <td><span className={style.tableAction} onClick={onUpdate(val.id)}>Тык!</span></td>
                        <td>
                            <DownloadFile ClassName={style.tableAction} Fetch={"http://localhost:80/api/v1/worker/application/download/" + val.link}
                                FileName={"document_"+val.name+"_" + val.fio + ".pdf"}/>
                        </td>
                    </tr>
                )}
                </tbody>
            </Table>
            {alert}
            <ApplicationModal Show={showModal} SetShow={setShowModal} OnOk={onOK(selectId)}/>
        </main>
    );
}