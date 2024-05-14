'use client'

import Form from "react-bootstrap/Form";

export interface SelectRoleProps {
    Roles : string[]
    Id : string
    Name : string
    Title? : string
    ClassName? : string
}

export default function SelectRole({Roles, Id, Name, ClassName, Title} : SelectRoleProps) {
    return (
        <Form.Group controlId={Id} className={ClassName}>
            <Form.Label>{Title ?? "Ваша роль:"}</Form.Label>
            <Form.Select required name={Name}>
                {Roles.map((val, idx) =>
                    (
                        <option value={val} key={idx}>{val}</option>
                    ))
                }
            </Form.Select>
        </Form.Group>
    )
}