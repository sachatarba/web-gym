// import React, { useState, useEffect } from 'react';
// import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// import Navbar from './components/Navbar/Navbar';
// import RegisterPage from './pages/RegisterPage';
// import LoginPage from './pages/LoginPage';
// import HallsPage from './pages/HallsPage';
// import GymPage from './pages/GymPage';
// import TrainingPage from './pages/TrainingPage';
// import SchedulePage from './pages/SchedulePage';

// // function App() {
// //   const [isLoggedIn, setIsLoggedIn] = useState(false);
// //   // const [client, setClient] = useState(null);

// //   const handleLogout = () => {
// //     setIsLoggedIn(false);
// //   };

// //   const handleLogin = () => {
// //     setIsLoggedIn(true);
// //   };

// //   return (
// //     <Router>
// //       <Navbar isLoggedIn={isLoggedIn} handleLogout={handleLogout} />
// //       <Routes>
// //         <Route path="/register" element={<RegisterPage />} />
// //         <Route path="/login" element={<LoginPage handleLogin={handleLogin} />} />
// //         <Route path="/halls" element={<HallsPage />} />
// //         <Route path="/gym/:id" element={<GymPage />} />
// //         <Route path="/training/:id" element={<TrainingPage isLoggedIn={isLoggedIn}/>} />
// //         <Route path="/schedule/:id" element={<SchedulePage clientId={isLoggedIn}/>} />
// //       </Routes>
// //     </Router>
// //   );
// // }

// // export default App;

// // import React, { useState, useEffect } from 'react';
// import { Navigate } from 'react-router-dom';
// // import Navbar from './components/Navbar';
// // import LoginPage from './pages/LoginPage';
// // // import RegistrationPage from './pages/RegistrationPage';
// // import HallsPage from './pages/HallsPage';
// // import GymPage from './pages/GymPage';
// // import TrainingPage from './pages/TrainingPage';
// // import SchedulePage from './pages/SchedulePage';
// import PrivateRoute from './components/PrivateRoute/PrivateRoute';

// const App = () => {
//   const [isAuthenticated, setIsAuthenticated] = useState(false);
//   const [clientId, setClientId] = useState(null);

//   useEffect(() => {
//     // Здесь вы можете добавить логику для проверки, авторизован ли пользователь, например, проверка токена в localStorage
//     const token = localStorage.getItem('token');
//     if (token) {
//       setIsAuthenticated(true);
//       // Установите clientId из токена или другого источника
//       setClientId('example-client-id');
//     }
//   }, []);

//   const handleLogin = (clientId) => {
//     setIsAuthenticated(true);
//     setClientId(clientId);
//     localStorage.setItem('token', 'your-token'); // Сохраните токен в localStorage или используйте другой метод
//   };

//   const handleLogout = () => {
//     setIsAuthenticated(false);
//     setClientId(null);
//     localStorage.removeItem('token'); // Удалите токен из localStorage или используйте другой метод
//   };

//   return (
//     <Router>
//       <Navbar isAuthenticated={isAuthenticated} handleLogout={handleLogout} />
//       <Routes>
//         <Route path="/login">
//           <LoginPage handleLogin={handleLogin} />
//         </Route>
//         <Route path="/register" component={RegisterPage} />
//         <Route path="/gyms" component={HallsPage} />
//         <Route path="/gym/:id" component={GymPage} />
//         <PrivateRoute
//           path="/training/:trainer_id"
//           component={() => <TrainingPage isLoggedIn={isAuthenticated} />}
//           isAuthenticated={isAuthenticated}
//         />
//         <PrivateRoute
//           path="/schedule/:training_id"
//           component={() => <SchedulePage clientId={clientId} />}
//           isAuthenticated={isAuthenticated}
//         />
//         <Navigate from="/" to="/gyms" />
//       </Routes>
//     </Router>
//   );
// };

// export default App;

// import React, { useState, useEffect } from 'react';
// import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
// import Navbar from './components/Navbar/Navbar';
// import RegisterPage from './pages/RegisterPage';
// import LoginPage from './pages/LoginPage';
// import HallsPage from './pages/HallsPage';
// import GymPage from './pages/GymPage';
// import TrainingPage from './pages/TrainingPage';
// import SchedulePage from './pages/SchedulePage';
// import PrivateRoute from './components/PrivateRoute/PrivateRoute';

// const App = () => {
//   const [isAuthenticated, setIsAuthenticated] = useState(false);
//   const [clientId, setClientId] = useState(null);

//   // useEffect(() => {
//   //   const token = localStorage.getItem('token');
//   //   if (token) {
//   //     setIsAuthenticated(true);
//   //     setClientId('example-client-id');
//   //   }
//   // }, []);

//   const handleLogin = (clientId) => {
//     setIsAuthenticated(true);
//     setClientId(clientId);
//     localStorage.setItem('token', 'your-token');
//   };

//   const handleLogout = () => {
//     setIsAuthenticated(false);
//     setClientId(null);
//     localStorage.removeItem('token');
//   };

//   return (
//     <Router>
//       <Navbar isAuthenticated={isAuthenticated} handleLogout={handleLogout} />
//       <Routes>
//         <Route path="/login" element={<LoginPage handleLogin={handleLogin} />} />
//         <Route path="/register" element={<RegisterPage />} />
//         <Route path="/halls" element={<HallsPage />} />
//         <Route path="/gym/:id" element={<GymPage />} />
//         <Route
//           path="/training/:trainer_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <TrainingPage isLoggedIn={isAuthenticated} />
//             </PrivateRoute>
//           }
//         />
//         <Route
//           path="/schedule/:training_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <SchedulePage clientId={clientId} />
//             </PrivateRoute>
//           }
//         />
//         <Route path="/" element={<Navigate to="/gyms" />} />
//       </Routes>
//     </Router>
//   );
// };

// export default App;

// import React, { useState, useEffect } from 'react';
// import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
// import Navbar from './components/Navbar/Navbar';
// import RegisterPage from './pages/RegisterPage';
// import LoginPage from './pages/LoginPage';
// import HallsPage from './pages/HallsPage';
// import GymPage from './pages/GymPage';
// import TrainingPage from './pages/TrainingPage';
// import SchedulePage from './pages/SchedulePage';
// import PrivateRoute from './components/PrivateRoute/PrivateRoute';

// const App = () => {
//   const [isAuthenticated, setIsAuthenticated] = useState(false);
//   const [clientId, setClientId] = useState(null);

//   useEffect(() => {
//     const token = localStorage.getItem('token');
//     if (token) {
//       setIsAuthenticated(true);
//       setClientId('example-client-id');
//     }
//   }, []);

//   const handleLogin = (clientId) => {
//     setIsAuthenticated(true);
//     setClientId(clientId);
//     localStorage.setItem('token', 'your-token');
//   };

//   const handleLogout = () => {
//     setIsAuthenticated(false);
//     setClientId(null);
//     localStorage.removeItem('token');
//   };

//   return (
//     <Router>
//       <Navbar isAuthenticated={isAuthenticated} handleLogout={handleLogout} />
//       <Routes>
//         <Route path="/login" element={<LoginPage handleLogin={handleLogin} />} />
//         <Route path="/register" element={<RegisterPage />} />
//         <Route path="/gyms" element={<HallsPage />} />
//         <Route path="/gym/:id" element={<GymPage />} />
//         <Route
//           path="/training/:trainer_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <TrainingPage isLoggedIn={isAuthenticated} />
//             </PrivateRoute>
//           }
//         />
//         <Route
//           path="/schedule/:training_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <SchedulePage clientId={clientId} />
//             </PrivateRoute>
//           }
//         />
//         <Route path="/" element={<Navigate to="/gyms" />} />
//       </Routes>
//     </Router>
//   );
// };

// export default App;

// import React, { useState, useEffect } from 'react';
// import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
// import Navbar from './components/Navbar/Navbar';
// import RegisterPage from './pages/RegisterPage';
// import LoginPage from './pages/LoginPage';
// import HallsPage from './pages/HallsPage';
// import GymPage from './pages/GymPage';
// import TrainingPage from './pages/TrainingPage';
// import SchedulePage from './pages/SchedulePage';
// import ClientProfile from './pages/ClientProfile';
// import PrivateRoute from './components/PrivateRoute/PrivateRoute';

// const App = () => {
//   const [isAuthenticated, setIsAuthenticated] = useState(false);
//   const [clientId, setClientId] = useState(null);

//   useEffect(() => {
//     const token = localStorage.getItem('token');
//     if (token) {
//       setIsAuthenticated(true);
//       setClientId('example-client-id'); 
//     }
//   }, []);

//   const handleLogin = (clientId) => {
//     setIsAuthenticated(true);
//     setClientId(clientId);
//     localStorage.setItem('token', 'your-token'); 
//   };

//   const handleLogout = () => {
//     setIsAuthenticated(false);
//     setClientId(null);
//     localStorage.removeItem('token');
//   };

//   return (
//     <Router>
//       <Navbar isAuthenticated={isAuthenticated} handleLogout={handleLogout} />
//       <Routes>
//         <Route path="/login" element={<LoginPage handleLogin={handleLogin} />} />
//         <Route path="/register" element={<RegisterPage />} />
//         <Route path="/gyms" element={<HallsPage />} />
//         <Route path="/gym/:id" element={<GymPage />} />
//         <Route
//           path="/training/:trainer_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <TrainingPage isLoggedIn={isAuthenticated} />
//             </PrivateRoute>
//           }
//         />
//         <Route
//           path="/schedule/:training_id"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <SchedulePage clientId={clientId} />
//             </PrivateRoute>
//           }
//         />
//         <Route
//           path="/profile"
//           element={
//             <PrivateRoute isAuthenticated={isAuthenticated}>
//               <ClientProfile clientId={clientId} />
//             </PrivateRoute>
//           }
//         />
//         <Route path="/" element={<Navigate to="/gyms" />} />
//       </Routes>
//     </Router>
//   );
// };

// export default App;


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
      // alert('Failed to sign up for training');
    });
    // /api/v1/isauthorize
    // if (token) {
    //   setIsAuthenticated(true);
    //   setClientId('example-client-id'); // Replace with logic to extract clientId from token
    // }
  }, []);

  const handleLogin = (clientId) => {
    setIsAuthenticated(true);
    setClientId(clientId);
    // localStorage.setItem('token', 'your-token'); // Replace 'your-token' with actual token
  };

  const handleLogout = () => {
    setIsAuthenticated(false);
    setClientId(null);
    // localStorage.removeItem('token');
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


