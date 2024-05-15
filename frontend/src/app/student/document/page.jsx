'use client'

import useSWR from "swr";
import React, {useState} from "react";
import {Spinner, Table} from "react-bootstrap";
import style from "@/styles/table/table.module.css";
import DownloadFile from "@/elements/downloadFile/downloadFile";

export default function Home() {
    const fetcher = () => fetch("https://localhost:443/api/v1/student/document", {
        method: "GET",
    }).then(response => response.json())


    const {data, isLoading} = useSWR("/v1/document/student", fetcher)

    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Table>
                <thead>
                <tr>
                    <th>Название</th>
                    <th>Тип</th>
                    <th>Создана</th>
                    <th>Статус</th>
                    <th>Скачать</th>
                </tr>
                </thead>
                <tbody>
                {data.map((val, idx) =>
                    <tr key={idx}>
                        <td>{val.name}</td>
                        <td>{val.type}</td>
                        <td>{val.created_at.substring(0, 10)}</td>
                        <td>{val.status}</td>
                        <td>
                            <DownloadFile ClassName={style.tableAction}
                                          Fetch={"https://localhost:443/api/v1/student/document/download/" + val.link}
                                          FileName={"document_" + val.name + "_" + ".pdf"}/>
                        </td>
                    </tr>
                )}
                </tbody>
            </Table>
        </main>
    );
}