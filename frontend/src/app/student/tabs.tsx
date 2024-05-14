'use client'

import {Nav} from "react-bootstrap";
import {useParams, useRouter} from "next/navigation";

export interface TabsProperty {
    ClassName? : string
}

export default function Tabs({ClassName} : TabsProperty) {
    const router = useRouter()
    return (
        <Nav fill variant="tabs" className={ClassName} defaultActiveKey={"/student"} justify>
            <Nav.Item defaultChecked>
                <Nav.Link onClick={ () => router.push("/student")} eventKey={"/student"}>
                    Посмотреть заявки
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/student/document")} eventKey={"/student/document"}>
                    Посмотреть документы
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/student/addApplication")} eventKey={"/student/addApplication"}>
                    Добавить заявку
                </Nav.Link>
            </Nav.Item>
        </Nav>
    )
}