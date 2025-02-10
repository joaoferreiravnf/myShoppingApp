import React from "react";
import { useAuth0 } from "@auth0/auth0-react";

const LoginPage = () => {
    const { loginWithRedirect } = useAuth0();

    return (
        <div className="login-page">
            <h1>Welcome to My App</h1>
            <button onClick={() => loginWithRedirect()}>Log In</button>
        </div>
    );
};

export default LoginPage;