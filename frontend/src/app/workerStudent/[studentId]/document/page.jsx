'use client'

import useSWR from "swr";
import React, {useState} from "react";
import {Spinner, Table} from "react-bootstrap";
import style from "@/styles/table/table.module.css";
import DownloadFile from "@/elements/downloadFile/downloadFile";
import {useParams} from "next/navigation";
import UpdateModal from "@/app/workerStudent/[studentId]/document/updateModal";
import UpdateFileModal from "@/app/workerStudent/[studentId]/document/updateFIleModal";

export default function Home() {
    const {studentId} = useParams()
    const fetcher = () => fetch("http://localhost:80/api/v1/worker/student/document/" + studentId, {
        method: "GET",
    }).then(response => response.json())


    const {data, isLoading, mutate} = useSWR("/v1/document", fetcher)
    const [selectData, setSelectedData] = useState(null)
    const [showUpdate, setShowUpdate] = useState(false)
    const [setedId, setId] = useState<string>("")
    const [showUpdateFile, setShowUpdateFile] = useState(false)

    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    const onUpdate = (val ) => {
        return () => {
            setSelectedData(val)
            setShowUpdate(true)
        }
    }

    const onUpdateFile = (id) => {
        return () => {
            setId(id)
            setShowUpdateFile(true)

        }
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Table>
                <thead>
                <tr>
                    <th>Название</th>
                    <th>Тип</th>
                    <th>Создана</th>
                    <th>Обновить</th>
                    <th>Обновить файл</th>
                    <th>Скачать</th>
                </tr>
                </thead>
                <tbody>
                {data.map((val, idx) =>
                    <tr key={idx}>
                        <td>{val.name}</td>
                        <td>{val.type}</td>
                        <td>{val.created_at.substring(0, 10)}</td>
                        <td><span className={style.tableAction} onClick={onUpdate(val)}>Тык!</span></td>
                        <td><span className={style.tableAction} onClick={onUpdateFile(val.id)}>Тык!</span></td>
                        <td>
                            <DownloadFile ClassName={style.tableAction}
                                          Fetch={"http://localhost:80/api/v1/worker/document/download/" + val.link}
                                          FileName={"document_" + val.name + "_" + ".pdf"}/>
                        </td>
                    </tr>
                )}
                </tbody>
            </Table>
            <UpdateModal Data={selectData} onClose={() => setShowUpdate(false)} mutate={mutate} Show={showUpdate}/>
            <UpdateFileModal Id={setedId} onClose={() => setShowUpdateFile(false)} mutate={mutate} Show={showUpdateFile}/>
        </main>
    );
}