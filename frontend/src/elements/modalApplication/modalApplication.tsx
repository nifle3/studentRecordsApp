'use client'

import {Dispatch, SetStateAction} from "react";
import {Modal} from "react-bootstrap";
import Button from "react-bootstrap/Button";

export interface ModalProps {
    Show : boolean
    SetShow : Dispatch<SetStateAction<boolean>>
    OnOk : () => void
}

export default function ApplicationModal({Show, SetShow, OnOk} : ModalProps) {
    return (
        <Modal show={Show} onHide={() => SetShow(false)} centered keyboard={false} backdrop={"static"}>
            <Modal.Header>
                <Modal.Title>Вы уверены, что хотите закрыть заявку?</Modal.Title>
            </Modal.Header>
            <Modal.Footer>
                <Button variant={"secondary"} onClick={() => SetShow(false)}>Отмена</Button>
                <Button variant={"danger"} onClick={OnOk}>Закрыть заявку</Button>
            </Modal.Footer>
        </Modal>
    )
}