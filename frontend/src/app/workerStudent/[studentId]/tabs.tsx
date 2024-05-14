'use client'

import {Nav} from "react-bootstrap";
import {useParams, useRouter} from "next/navigation";

export interface TabsProperty {
    ClassName? : string
}

export default function Tabs({ClassName} : TabsProperty) {
    const router = useRouter()
    const {studentId} = useParams<{studentId : string}>()
    return (
        <Nav fill variant="tabs" className={ClassName} defaultActiveKey={"/workerStudent/"} justify>
            <Nav.Item defaultChecked>
                <Nav.Link onClick={ () => router.push("/workerStudent/" + studentId)} eventKey={"/workerStudent/"}>
                    Посмотреть заявки
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/workerStudent/" + studentId + "/document")} eventKey={"/workerStudent/document"}>
                    Посмотреть документы
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/workerStudent/" + studentId + "/addDocument")} eventKey={"/workerStudent/addDocument"}>
                    Добавить документ
                </Nav.Link>
            </Nav.Item>
        </Nav>
    )
}