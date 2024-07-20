import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar/Navbar';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';
import HallsPage from './pages/HallsPage';
import GymPage from './pages/GymPage';
import TrainingPage from './pages/TrainingPage';
import SchedulePage from './pages/SchedulePage';
import ClientProfilePage from './pages/ClientProfilePage';
import PrivateRoute from './components/PrivateRoute/PrivateRoute';
import axios from 'axios';
import AdminProfile from './components/AdminProfile/AdminProfile';

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [clientId, setClientId] = useState(null);
  const [isAuthLoading, setIsAuthLoading] = useState(true);

  useEffect(() => {
    setIsAuthLoading(true);
    axios.get("http://localhost:8080/api/v1/isauthorize", {withCredentials: true})
    .then(response => {
      if (response.data.session) {
        setIsAuthenticated(true);
        setClientId(response.data.session.ClientID);
      } else {
        setIsAuthenticated(false);
        setClientId(null);
      }
      setIsAuthLoading(false);
    })
    .catch(error => {
      console.error('Unathorized:', error);
      setIsAuthLoading(false);
    });
  }, []);

  const handleLogin = (clientId) => {
    setIsAuthenticated(true);
    setClientId(clientId);
  };

  const handleLogout = () => {
    setIsAuthenticated(false);
    setClientId(null);
  };

  return (
    <Router>
      <Navbar isAuthenticated={isAuthenticated} handleLogout={handleLogout} />
      <Routes>
        <Route path="/login" element={<LoginPage handleLogin={handleLogin} />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/gyms" element={<HallsPage />} />
        <Route path="/gym/:id" element={<GymPage isAuthenticated={isAuthenticated} clientId={clientId}/>} />
        <Route
          path="/training/:trainer_id"
          element={
            <PrivateRoute isAuthenticated={isAuthenticated} isAuthLoading={isAuthLoading}>
              <TrainingPage isLoggedIn={isAuthenticated} />
            </PrivateRoute>
          }
        />
        <Route
          path="/schedule/:training_id"
          element={
            <PrivateRoute isAuthenticated={isAuthenticated} isAuthLoading={isAuthLoading}>
              <SchedulePage clientId={clientId} />
            </PrivateRoute>
          }
        />
        <Route
          path="/profile"
          element={
            <PrivateRoute isAuthenticated={isAuthenticated} isAuthLoading={isAuthLoading}>
              <ClientProfilePage clientId={clientId} />
            </PrivateRoute>
          }
        />
        <Route path="/" element={<Navigate to="/gyms" />} />
        <Route path="/admin" element={<AdminProfile  />} />
      </Routes>
    </Router>
    );
};

export default App;


