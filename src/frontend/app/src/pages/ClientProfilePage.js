import React from 'react';
import ClientProfile from '../components/ClientProfile/ClientProfile';

const ClientProfilePage = ({ clientId }) => {
  return (
    <div>
      <ClientProfile clientId={clientId} />
    </div>
  );
};

export default ClientProfilePage;
