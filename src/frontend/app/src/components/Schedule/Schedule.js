import React, { useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { v4 as uuidv4 } from 'uuid';
import './Schedule.css';

const Schedule = ({ clientId }) => {
  const { training_id } = useParams();
  const [dayOfTheWeek, setDayOfTheWeek] = useState('');
  const [startTime, setStartTime] = useState('');
  const [endTime, setEndTime] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    const newSchedule = {
        ID: uuidv4(),
        dayOfTheWeek,
        startTime: `${startTime}:00`,
        endTime: `${endTime}:00`,
        ClientId: clientId,
        TrainingId: training_id
    };
    
    console.log(newSchedule);
    axios.post('http://localhost:8080/api/v1/schedule/new', newSchedule)
      .then(response => {
        alert('Successfully signed up for training');
      })
      .catch(error => {
        console.error('Error signing up for training:', error);
        alert('Failed to sign up for training');
      });
  };

  return (
    <div className="schedule-page">
      <h1>Запись на тренировку</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Дата:
          <input
            type="date"
            value={dayOfTheWeek}
            onChange={(e) => setDayOfTheWeek(e.target.value)}
          />
        </label>
        <label>
          Время начала:
          <input
            type="time"
            value={startTime}
            onChange={(e) => setStartTime(e.target.value)}
          />
        </label>
        <label>
          Время окончания:
          <input
            type="time"
            value={endTime}
            onChange={(e) => setEndTime(e.target.value)}
          />
        </label>
        <button type="submit">Записаться</button>
      </form>
    </div>
  );
};

export default Schedule;
