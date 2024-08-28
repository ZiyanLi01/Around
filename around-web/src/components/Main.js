import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";

import Login from "./Login";
import Register from "./Register";
import Home from "./Home";

function Main(props) {
    const { isLoggedIn, handleLoggedIn } = props;

    const showLogin = () => {
        return isLoggedIn ? (
            <Navigate to="/home" />
        ) : (
            <Login handleLoggedIn={handleLoggedIn} />
        );
    };

    const showHome = () => {
        return isLoggedIn ? <Home /> : <Navigate to="/login" />;
    };

    return (
        <div className="main">
            <Routes>
                <Route path="/" element={showLogin()} />
                <Route path="/login" element={showLogin()} />
                <Route path="/register" element={<Register />} />
                <Route path="/home" element={showHome()} />
                {/* Fallback route to redirect any unmatched routes */}
                <Route path="*" element={<Navigate to="/" />} />
            </Routes>
        </div>
    );
}

export default Main;
