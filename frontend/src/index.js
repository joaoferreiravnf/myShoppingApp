import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { Auth0Provider } from "@auth0/auth0-react";
import App from './App';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
    <Auth0Provider
        domain="dev-d8ok352njndoo8ks.eu.auth0.com" // Replace with your Auth0 domain
        clientId="98fMsKnzJxIEqMQG4tBvggpGhDTV7vVi" // Replace with your Auth0 client ID
        authorizationParams={{
            redirect_uri: window.location.origin, // Redirect back to your app after login
        }}
    >
        <App />
    </Auth0Provider>,
    document.getElementById("root")
);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
