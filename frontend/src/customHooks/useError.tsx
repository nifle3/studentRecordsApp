import {Dispatch, SetStateAction, useState} from "react";
import CustomAlert from "@/elements/customAlert/customAlert";

export interface UserErrorReturning {
    alert: JSX.Element
    setError: Dispatch<SetStateAction<string>>
    setErrorCode: Dispatch<SetStateAction<number>>
    setShow: Dispatch<SetStateAction<boolean>>
}

export function useError() : UserErrorReturning {
    const [showError, setShowError] = useState<boolean>(false)
    const [errorCode, setErrorCode] = useState<number>(200)
    const [error, setError] = useState<string>("")

    const alertError = <CustomAlert IsShow={showError} SetIsShow={setShowError} Error={error}
                               ErrorCode={errorCode}/>

    return {
        setShow: setShowError,
        setErrorCode: setErrorCode,
        setError: setError,
        alert: alertError,
    }
}