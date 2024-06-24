import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Halls.css';

const Gyms = () => {
  const [gyms, setGyms] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get('http://localhost:8080/api/v1/gym/all')
      .then(response => {
        if (response.data && Array.isArray(response.data.gyms)) {
          setGyms(response.data.gyms);
        } else {
          console.error('Unexpected response format:', response.data);
        }
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching gyms:', error);
        setLoading(false);
      });
  }, []);

  if (loading)
    return <p>Загрузка</p>

  if (gyms.length === 0)
    return <h2>Пока здесь пусто</h2>

  return (
    <div className="halls-page">
      <h1>Залы</h1>
        <ul className="halls-list">
          {gyms.map(gym => (
            <li key={gym.ID} className="hall-item">
              <Link to={`/gym/${gym.ID}`}>
                <h2>{gym.Name}</h2>
                <p>Телефон: {gym.Phone}</p>
                <p>Город: {gym.City}</p>
                <p>Адрес: {gym.Addres}</p>
                {/* <p>Chain: {gym.IsChain ? 'Да' : 'Нет'}</p> */}
              </Link>
            </li>
          ))}
        </ul>
    </div>
  );
};

export default Gyms;
