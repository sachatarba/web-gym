// import React, { useState } from 'react';
// import axios from 'axios';
// import './Login.css';

// const Login = ({ handleLogin }) => {
//   const [login, setLogin] = useState('');
//   const [password, setPassword] = useState('');

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     axios.post('http://localhost:8080/api/v1/login', { login, password })
//       .then(response => {
//         handleLogin();
//         alert('Login successful');
//       })
//       .catch(error => {
//         alert(`Login failed ${error}`);
//       });
//   };

//   return (
//     <div className="login-page">
//       <h1>Добро пожаловать!</h1>
//       <form onSubmit={handleSubmit}>
//         {/* <label> */}
//           <input
//             type="text"
//             value={login}
//             placeholder='Логин'
//             onChange={(e) => setLogin(e.target.value)}
//           />
//         {/* </label> */}
//         {/* <label> */}
//           <input
//             type="password"
//             value={password}
//             placeholder='Пароль'
//             onChange={(e) => setPassword(e.target.value)}
//           />
//         {/* </label> */}
//         <button type="submit">Войти</button>
//       </form>
//     </div>
//   );
// };

// export default Login;


import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useLocation } from 'react-router-dom';
import './Login.css';

const   Login = ({ handleLogin }) => {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || '/gyms';

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/api/v1/login', { login, password }, {withCredentials : true})
      .then(response => {
        // console.log(response.data.session.ClientID)
        handleLogin(response.data.session.ClientID); // Assuming the server returns clientId in the response
        alert('Login successful');
        navigate(from, { replace: true });
      })
      .catch(error => {
        alert(`Login failed: ${error}`);
      });
  };

  return (
    <div className="login-page">
      <h1>Добро пожаловать!</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={login}
          placeholder="Логин"
          onChange={(e) => setLogin(e.target.value)}
        />
        <input
          type="password"
          value={password}
          placeholder="Пароль"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Войти</button>
      </form>
    </div>
  );
};

export default Login;
