'use client'

import Form from "react-bootstrap/Form";

type inputType = "email" | "text" | "tel" | "number" | "date" | "password" | "file"

export interface InputProps {
    Name : string
    Id : string
    Type : inputType
    LabelHolder : string
    PlaceHolder : string
    ClassName? : string
    Value? : string | number | string[]
    Allowed? : string
    ReadOnly? : boolean
}

export default function Input({Name, Type, LabelHolder, PlaceHolder, Value, Id, ClassName, Allowed, ReadOnly} : InputProps) {
    return (
        <Form.Group controlId={Id} className={ClassName}>
            <Form.Label>{LabelHolder}</Form.Label>
            <Form.Control name={Name} type={Type}
                placeholder={PlaceHolder}  maxLength={50} required defaultValue={Value} accept={Allowed}
            readOnly={ReadOnly}/>
        </Form.Group>
    )
}