import React from 'react';
import Training from '../components/Training/Training';
import '../components/Training/Training.css';

const TrainingPage = ({ isLoggedIn }) => {
  return (
    <div className="training-page">
      <Training isLoggedIn={isLoggedIn} />
    </div>
  );
};

export default TrainingPage;
