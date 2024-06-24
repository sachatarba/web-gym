import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, Navigate, useNavigate } from 'react-router-dom';
import './Training.css';

const types = {
  'aerobic': 'аэробная',
  'anaerobic': 'анаэробная', 
  'flexibility': 'гибкость',
  'strength': 'силовая'
}

const Training = ({ isLoggedIn }) => {
  const { trainer_id } = useParams();
  const [trainings, setTrainings] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    axios.get(`http://localhost:8080/api/v1/training/${trainer_id}`)
      .then(response => {
        setTrainings(response.data.trainings);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching trainings:', error);
        setLoading(false);
      });
  }, [trainer_id]);

  const toSchedule = (id) => {
    return () => {
      navigate(`/schedule/${id}`)
    } 
  }

  return (
    <div className="training-page">
      <h1>Тренировки</h1>
      {loading ? (
        <p>Загрузка...</p>
      ) : (
        <ul className="training-list">
          {trainings.map(training => (
            <li key={training.ID} className="training-item">
              <h3>{training.Title}</h3>
              <p>Описание: {training.Description}</p>
              <p>Тип: {types[training.TrainingType]}</p>
              {isLoggedIn && (
                <button onClick={toSchedule(training.ID)}>
                  Записаться
                </button>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Training;
