'use client'

import {Spinner, Table} from "react-bootstrap";
import React from "react";
import useSWR from "swr";
import style from "@/styles/table/table.module.css"
import {useRouter} from "next/navigation";

const fetcher = () => fetch("https://localhost:443/api/v1/worker/student", {
    method: "GET",
}).then(response => response.json())

export default function Home() {
    const router = useRouter()
    const {data, isLoading} = useSWR("/v1/students", fetcher)

    if (isLoading) {
        return  (
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        )
    }

    const onClick = (id) => {
        return () => {
            router.push("/workerStudent/" + id)
        }
    }

    return (
        <main style={{marginTop:"2rem"}}>
            <Table>
                <thead>
                <tr>
                    <th>Имя</th>
                    <th>Фамилия</th>
                    <th>Отчество</th>
                    <th>Год поступления</th>
                    <th>Специализация</th>
                    <th>Курс</th>
                    <th>Группа</th>
                    <th>Подробнее</th>
                </tr>
                </thead>
                <tbody>
                {data.map((val, idx) =>
                        <tr key={idx}>
                            <td>{val.first_name}</td>
                            <td>{val.last_name}</td>
                            <td>{val.surname}</td>
                            <td>{val.enroll_year.substring(0, 10)}</td>
                            <td>{val.specialization}</td>
                            <td>{val.course}</td>
                            <td>{val.group}</td>
                            <td><span className={style.tableAction} onClick={onClick(val.id)}>Тык!</span></td>
                        </tr>
                )}
                </tbody>
            </Table>
        </main>
    );
}