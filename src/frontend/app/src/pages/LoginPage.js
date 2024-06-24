import React from 'react';
import Login from '../components/Login/Login';
import '../components/Login/Login.css';

const LoginPage = ({ handleLogin }) => {
  return (
    <div className="login-page">
      <Login handleLogin={handleLogin} />
    </div>
  );
};

export default LoginPage;
