import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { format, parseISO } from 'date-fns';
import { formatInTimeZone } from 'date-fns-tz';
import { ru } from 'date-fns/locale';
import './ClientProfile.css';

const types = {
  'aerobic': 'аэробная',
  'anaerobic': 'анаэробная', 
  'flexibility': 'гибкость',
  'strength': 'силовая'
}

const ClientProfile = ({ clientId }) => {
  const [client, setClient] = useState(null);
  const [schedules, setSchedules] = useState([]);
  const [memberships, setMemberships] = useState([]);
  const [editing, setEditing] = useState(false);
  const [showSchedules, setShowSchedules] = useState(false);
  const [showMemberships, setShowMemberships] = useState(false);
  const [formData, setFormData] = useState({
    Login: '',
    Password: '',
    Fullname: '',
    Email: '',
    Phone: '',
    Birthdate: ''
  });

  useEffect(() => {
    axios.get(`http://localhost:8080/api/v1/client/${clientId}`)
      .then(response => {
        const clientData = response.data.client;
        setClient(clientData);
        setFormData({
          Login: clientData.Login,
          Password: clientData.Password,
          Fullname: clientData.Fullname,
          Email: clientData.Email,
          Phone: clientData.Phone,
          Birthdate: clientData.Birthdate
        });
      })
      .catch(error => {
        console.error('Error fetching client data:', error);
      });
  }, [clientId]);

  const handleEditChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleEditSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/api/v1/client/change', { ...formData, id: clientId })
      .then(response => {
        setClient({ ...formData, id: clientId });
        setEditing(false);
        alert('Profile updated successfully');
      })
      .catch(error => {
        console.error('Error updating profile:', error);
        alert('Failed to update profile');
      });
  };

  const fetchSchedules = () => {
    setShowMemberships(false);
    axios.get(`http://localhost:8080/api/v1/schedule/${clientId}`)
      .then(response => {
        setSchedules(response.data.schedules);
        setShowSchedules(true);
      })
      .catch(error => {
        console.error('Error fetching schedules:', error);
      });
  };

  const fetchMemberships = () => {
    setShowSchedules(false);
    axios.get(`http://localhost:8080/api/v1/client_membership/${clientId}`)
      .then(response => {
        setMemberships(response.data.clientMemberships);
        setShowMemberships(true);
      })
      .catch(error => {
        console.error('Error fetching memberships:', error);
      });
  };

  const formatDate = (dateString) => {
    return format(parseISO(dateString), 'dd MMMM yyyy', { locale: ru });
  };

  const formatTime = (timeString) => {
    const timeZone = 'Europe/Moscow';
    const time = parseISO(timeString);
    return formatInTimeZone(time.setHours(time.getHours() - 3), timeZone, 'HH:mm', { locale: ru });
  };

  const calcDays = (start, end) => {
    const startDay = parseISO(start);
    const endDay = parseISO(end);
    const today = new Date();

    if (today < startDay) {
        return ((endDay - startDay) / ( 1000 * 60 * 60 * 24)).toFixed(0);
    }
    // console.log(endDay, startDay)
    const days = ((endDay - today) / ( 1000 * 60 * 60 * 24)).toFixed(0);
    return days;
  };

  return (
    <div className="client-profile">
      <h1>Личный кабинет</h1>
      {client && (
        <>
          {!editing ? (
            <div className="client-info">
              <p>Полное имя: {client.Fullname}</p>
              <p>Email: {client.Email}</p>
              <p>Телефон: {client.Phone}</p>
              <p>Дата рождения: {formatDate(client.Birthdate)}</p>
              <button onClick={() => setEditing(true)}>Редактировать</button>
            </div>
          ) : (
            <form onSubmit={handleEditSubmit} className="edit-form">
              <label>
                Полное имя:
                <input
                  type="text"
                  name="Fullname"
                  value={formData.Fullname}
                  onChange={handleEditChange}
                />
              </label>
              <label>
                Email:
                <input
                  type="email"
                  name="Email"
                  value={formData.Email}
                  onChange={handleEditChange}
                />
              </label>
              <label>
                Телефон:
                <input
                  type="text"
                  name="Phone"
                  value={formData.Phone}
                  onChange={handleEditChange}
                />
              </label>
              <label>
                Дата рождения:
                <input
                  type="date"
                  name="Birthdate"
                  value={formData.Birthdate}
                  onChange={handleEditChange}
                />
              </label>
              <button type="submit">Сохранить</button>
              <button type="button" onClick={() => setEditing(false)}>Отмена</button>
            </form>
          )}
          <div className="buttons">
            <button onClick={fetchSchedules}>Показать записи на тренировки</button>
            <button onClick={fetchMemberships}>Показать абонементы</button>
          </div>

          {showSchedules && (
            <div className="schedules">
              {schedules.map(schedule => (
                <div key={schedule.ID} className="schedule-item">
                  <p>Дата: {formatDate(schedule.DayOfTheWeek)}</p>
                  <p>Начало: {formatTime(schedule.StartTime)}</p>
                  <p>Конец: {formatTime(schedule.EndTime)}</p>
                  <div className="training-details">
                    <p>Название тренировки: {schedule.Training.Title}</p>
                    <p>Описание: {schedule.Training.Description}</p>
                    <p>Тип тренировки: {types[schedule.Training.TrainingType]}</p>
                  </div>
                </div>
              ))}
            </div>
          )}

          {showMemberships && (
            <div className="memberships">
              {memberships.map(membership => (
                <div key={membership.ID} className="membership-item">
                  <p>Начало: {formatDate(membership.StartDate)}</p>
                  <p>Конец: {formatDate(membership.EndDate)}</p>
                  <p>Тип абонемента: {membership.MembershipType.Type}</p>
                  <p>Описание: {membership.MembershipType.Description}</p>
                  {/* <p>Цена: {membership.MembershipType.Price}</p> */}
                  <p>Длительность: {(calcDays(membership.StartDate, membership.EndDate) >=0 ? calcDays(membership.StartDate, membership.EndDate) + ' дней' : "Не действителен")}</p>
                </div>
              ))}
            </div>
          )}
        </>
      )}
    </div>
  );
};

export default ClientProfile;




