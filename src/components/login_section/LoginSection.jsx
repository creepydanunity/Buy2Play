import "./LoginSection.scss";
import { useState } from "react";

function LoginSection() {
    const [email, setEmail] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("Email submitted:", email);
    };

    return (
        <div className="login-container">
            <div className="login-card">
                <img src="/logo.svg" alt="Logo" className="logo" />
                <h1>Вход</h1>
                <p>
                    Укажите ваш адрес электронной почты, на который мы отправим код для входа
                </p>
                <form onSubmit={handleSubmit}>
                    <input
                        type="email"
                        placeholder="Электронный адрес"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        className="input"
                    />
                    <button type="submit" className="button">
                        Продолжить
                    </button>
                </form>
                <a href="/privacy" className="privacy-link">
                    Конфиденциальность
                </a>
            </div>
        </div>
    );
}

export default LoginSection;
