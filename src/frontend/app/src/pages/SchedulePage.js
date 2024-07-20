import React from 'react';
import Schedule from '../components/Schedule/Schedule';

const SchedulePage = ({ clientId }) => {
  return (
    <div className="schedule-page">
      <Schedule clientId={clientId} />
    </div>
  );
};

export default SchedulePage;
