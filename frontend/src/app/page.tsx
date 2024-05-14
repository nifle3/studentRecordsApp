import style from "@/styles/login/page.module.css"
import backgroundImage from "@/icons/app/background.jpg"
import AuthForm from "./authForm";
import "@/styles/global.css"
import 'bootstrap/dist/css/bootstrap.min.css';

// TODO: Add LOGO
export default function Home() {
  return (
    <main className={style.wrapper}>
        <img src={backgroundImage.src} alt={""} className={style.image}/>
        <div className={style.loginWrapper}>
            <AuthForm/>
        </div>
    </main>
  );
}
