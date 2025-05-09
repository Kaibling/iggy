import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export const showError = (msg: string) => {
    toast.error(msg);
  };

  export const showWarning = () => {
    toast.warn('This is a warning message!');
  };

  export const showInfo = (msg: string) => {
    toast.info(msg);
  };

  export const showSuccess = (msg: string) => {
    toast.success(msg);
  };



export const NotifyContainer = () => {
    return (
        <ToastContainer/>
    )
}