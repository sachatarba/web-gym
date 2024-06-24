import React from 'react';
import Gyms from '../components/Gym/Gym';
import '../components/Gym/Gym.css';

const GymPage = ({isAuthenticated, clientId}) => {
  return (
    <div className="gyms-page">
      <Gyms isAuthenticated={isAuthenticated} clientId={clientId}/>
    </div>
  );
};

export default GymPage;
