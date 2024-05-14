'use client'

import {Nav} from "react-bootstrap";
import {useRouter} from "next/navigation";

export interface TabsProperty {
    ClassName? : string
}

export default function Tabs({ClassName} : TabsProperty) {
    const router = useRouter()

    return (
        <Nav fill variant="tabs" className={ClassName} defaultActiveKey={"/admin"} justify>
            <Nav.Item defaultChecked>
                <Nav.Link onClick={ () => router.push("/admin")} eventKey={"/admin"}>
                    Посмотреть работников
                </Nav.Link>
            </Nav.Item>
            <Nav.Item>
                <Nav.Link onClick={() => router.push("/admin/addWorker")} eventKey={"/admin/addWorker"}>
                    Добавить работника
                </Nav.Link>
            </Nav.Item>
        </Nav>
    )
}