import React, { useState } from 'react';
import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';
import './RegistrationForm.css';

const RegistrationForm = () => {
  const [formData, setFormData] = useState({
    fullname: '',
    login: '',
    password: '',
    birthdate: '',
    email: '',
    phone: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const clientReq = {
      id: uuidv4(),
      ...formData,
    };
    axios.post('http://localhost:8080/api/v1/register', clientReq)
      .then((response) => {
        alert('Registration successful');
      })
      .catch((error) => {
        alert('Registration failed');
      });
  };

  return (
    <form onSubmit={handleSubmit} className="registration-form">
      <label>
        ФИО:
        <input
          type="text"
          name="fullname"
          value={formData.fullname}
          onChange={handleChange}
        />
      </label>
      <label>
        Логин:
        <input
          type="text"
          name="login"
          value={formData.login}
          onChange={handleChange}
        />
      </label>
      <label>
        Пароль:
        <input
          type="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
        />
      </label>
      <label>
        Дата рождения:
        <input
          type="date"
          name="birthdate"
          value={formData.birthdate}
          onChange={handleChange}
        />
      </label>
      <label>
        Email:
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
        />
      </label>
      <label>
        Номер телефона:
        <input
          type="tel"
          name="phone"
          value={formData.phone}
          onChange={handleChange}
        />
      </label>
      <button type="submit">Зарегистрироваться</button>
    </form>
  );
};

export default RegistrationForm;
