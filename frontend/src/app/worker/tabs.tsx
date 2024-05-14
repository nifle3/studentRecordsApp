'use client'

import {Nav} from "react-bootstrap";
import {useRouter} from "next/navigation";

export interface TabsProperty {
    ClassName? : string
}

export default function Tabs({ClassName} : TabsProperty) {
    const router = useRouter()

    return (
        <Nav fill variant="tabs" className={ClassName} defaultActiveKey={"/worker"} justify>
            <Nav.Item defaultChecked>
                <Nav.Link onClick={ () => router.push("/worker")} eventKey={"/worker"}>
                    Посмотреть студентов
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/worker/addStudent")} eventKey={"/worker/addStudent"}>
                    Добавить студента
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/worker/application")} eventKey={"/worker/application"}>
                    Посмотреть все заявки
                </Nav.Link>
            </Nav.Item>
        </Nav>
    )
}