'use client'

import {Dispatch, SetStateAction} from "react";
import {Modal} from "react-bootstrap";
import Button from "react-bootstrap/Button";

export interface DeleteModalProps {
    Show : boolean
    SetShow : Dispatch<SetStateAction<boolean>>
    OnOk : () => void
    DeletedObjectName : string
}

export default function DeleteModal({Show, SetShow, OnOk, DeletedObjectName} : DeleteModalProps) {
    return (
        <Modal show={Show} onHide={() => SetShow(false)} centered keyboard={false} backdrop={"static"}>
            <Modal.Header>
                <Modal.Title>Вы уверены, что хотите удалить {DeletedObjectName}</Modal.Title>
            </Modal.Header>
            <Modal.Footer>
                <Button variant={"secondary"} onClick={() => SetShow(false)}>Отмена</Button>
                <Button variant={"danger"} onClick={OnOk}>Удалить</Button>
            </Modal.Footer>
        </Modal>
    )
}