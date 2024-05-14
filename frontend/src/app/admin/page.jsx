'use client'

import useSWR from "swr";
import React, {useState} from "react";
import {Spinner, Table} from "react-bootstrap";
import tableStyle from "@/styles/table/table.module.css"
import DeleteModal from "@/elements/deleteModal/deleteModal";
import UpdateModal from "@/app/admin/updateModal";

const fetcher = () => fetch("http://localhost:80/api/v1/admin/worker", {
    method: "GET",
}).then(response => response.json())

export default function Home() {
    const {data, isLoading, mutate} = useSWR("/v1/workers", fetcher)
    const [showModal, setShowModal] = useState(false)
    const [showUpdate, setShowUpdate] = useState(false)
    const [selectId, setSelectId] = useState("")
    const [selectData, setSelectData] = useState(null)
    

    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }
    const onDelete = (id ) => {
        return () => {
            setSelectId(id)
            setShowModal(true)
        }
    }

    const onUpdate = (val) => {
        return () => {
            setSelectData(val)
            setShowUpdate(true)
        }
    }

    const onOk = async () => {
        await fetch("http://localhost:80/api/v1/admin/worker/" + selectId, {
            method: "DELETE",
        })
        await mutate()
        setShowModal(false)
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Table>
                <thead>
                    <tr>
                        <th>Имя</th>
                        <th>Фамилия</th>
                        <th>Отчество</th>
                        <th>Почта</th>
                        <th>Обновить</th>
                        <th>Удалить</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((val, idx) => (
                        <>
                            <tr key={idx}>
                                <td>{val.first_name}</td>
                                <td>{val.last_name}</td>
                                <td>{val.surname}</td>
                                <td>{val.email}</td>
                                <td><span onClick={onUpdate(val)}
                                    className={tableStyle.tableAction}>Обновить</span></td>
                                <td><span onClick={onDelete(val.id)}
                                    className={tableStyle.tableAction}>Удалить</span></td>
                            </tr>
                        </>
                        )
                    )}
                </tbody>
            </Table>
            <UpdateModal Data={selectData} mutate={mutate} onClose={() => setShowUpdate(false)} Show={showUpdate}/>
            <DeleteModal Show={showModal} SetShow={setShowModal} OnOk={onOk} DeletedObjectName={"работника"}/>
        </main>
    );
}