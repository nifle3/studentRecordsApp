import style from "@/styles/alert/alert.module.css"
import {Alert} from "react-bootstrap";
import {Dispatch, SetStateAction} from "react";

export interface CustomAlertInterface {
    IsShow : boolean
    SetIsShow :  Dispatch<SetStateAction<boolean>>
    Error : string
    ErrorCode : number
}

export default function CustomAlert({IsShow, SetIsShow, Error, ErrorCode} : CustomAlertInterface) {
    if (!IsShow) {
        return
    }

    return (
        <Alert variant={"danger"} className={style.alert} onClose={() => SetIsShow(false)} dismissible>
            <Alert.Heading>{ErrorCode}</Alert.Heading>
            <p>
                {Error}
            </p>
        </Alert>
    )
}