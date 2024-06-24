import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Navbar.css';

const Navbar = ({ isAuthenticated, handleLogout }) => {
  const navigate = useNavigate();

  const handleLogoutClick = () => {
    handleLogout();
    navigate('/login');
  };

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/">Качалка.ру</Link>
      </div>
      <ul className="navbar-links">
        {/* <li>
          <Link to="/gyms">Halls</Link>
        </li> */}
        {!isAuthenticated ? (
          <>
            <li>
              <Link to="/register">Зарегистрироваться</Link>
            </li>
            <li>
              <Link to="/login">Войти</Link>
            </li>
          </>
        ) : (
          <>
          <li>
            <Link to="/profile">Профиль</Link>
          </li>
          <li>
            <Link onClick={handleLogoutClick}>Выйти</Link>
          </li>
          </>
        )}
      </ul>
    </nav>
  );
};

export default Navbar;
