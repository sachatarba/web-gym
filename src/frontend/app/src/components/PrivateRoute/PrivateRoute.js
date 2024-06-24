// import React from 'react';
// import { Navigate } from 'react-router-dom';

// const PrivateRoute = ({ isAuthenticated, children }) => {
//   return isAuthenticated ? children : <Navigate to="/login" />;
// };

// export default PrivateRoute;

import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';

const PrivateRoute = ({ children, isAuthenticated, isAuthLoading }) => {
  const location = useLocation();
//   while (isAuthLoading) {
    // console.log("kek:", isAuthLoading)
//   }
//   console.log("kek:", isAuthLoading)
  if (isAuthLoading) {
    return <></>
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} />;
  }

  return children;
};

export default PrivateRoute;

