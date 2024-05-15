'use client'

import useSWR from "swr";
import {useError} from "@/customHooks/useError";
import React, {useState} from "react";
import {Spinner, Table} from "react-bootstrap";
import style from "@/styles/table/table.module.css";
import DownloadFile from "@/elements/downloadFile/downloadFile";
import ApplicationModal from "@/elements/modalApplication/modalApplication";
import {useParams} from "next/navigation";

export default function Home() {
    const {studentId} = useParams()
    const fetcher = () => fetch("https://localhost:443/api/v1/worker/student/application/" + studentId, {
        method: "GET",
    }).then(response => response.json())


    const {data, isLoading, mutate} = useSWR("/v1/application", fetcher)
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
            const response = await fetch("https://localhost:443/api/v1/worker/application/close/" + id, {
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
                    <th>Созадана</th>
                    <th>Текст</th>
                    <th>Контактная информация</th>
                    <th>Закрыть</th>
                    <th>Скачать</th>
                </tr>
                </thead>
                <tbody>
                {data.map((val, idx) =>
                    <tr key={idx}>
                        <td>{val.name}</td>
                        <td>{val.created_at.substring(0, 10)}</td>
                        <td>{val.text}</td>
                        <td>{val.contact_info}</td>
                        <td><span className={style.tableAction} onClick={onUpdate(val.id)}>Тык!</span></td>
                        <td>
                            <DownloadFile ClassName={style.tableAction} Fetch={"https://localhost:443/api/v1/worker/application/download/" + val.link}
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