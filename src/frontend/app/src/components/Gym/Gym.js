// import React, { useState, useEffect } from 'react';
// import { useParams } from 'react-router-dom';
// import axios from 'axios';
// import './Gym.css';

// const Gym = () => {
//   const { id } = useParams();
//   const [gym, setGym] = useState(null);
//   const [memberships, setMemberships] = useState([]);
//   const [equipments, setEquipments] = useState([]);
//   const [trainers, setTrainers] = useState([]);
//   const [loading, setLoading] = useState(true);
//   const [view, setView] = useState('');

//   useEffect(() => {
//     axios.get(`http://localhost:8080/api/v1/gym/${id}`)
//       .then(response => {
//         setGym(response.data.gym);
//         setLoading(false);
//       })
//       .catch(error => {
//         console.error('Error fetching gym details:', error);
//         setLoading(false);
//       });
//   }, [id]);

//   const fetchMemberships = () => {
//     axios.get(`http://localhost:8080/api/v1/membershipType/${id}`)
//       .then(response => {
//         setMemberships(response.data.membershipTypes);
//         setView('memberships');
//       })
//       .catch(error => {
//         console.error('Error fetching memberships:', error);
//       });
//   };

//   const fetchEquipments = () => {
//     axios.get(`http://localhost:8080/api/v1/equipment/${id}`)
//       .then(response => {
//         setEquipments(response.data.equipments);
//         setView('equipments');
//       })
//       .catch(error => {
//         console.error('Error fetching equipments:', error);
//       });
//   };

//   const fetchTrainers = () => {
//     axios.get(`http://localhost:8080/api/v1/trainer/${id}`)
//       .then(response => {
//         setTrainers(response.data.trainers);
//         setView('trainers');
//       })
//       .catch(error => {
//         console.error('Error fetching trainers:', error);
//       });
//   };

//   const renderMemberships = () => {
//     if (view === 'memberships') {
//       return (
//       <div>
//       <h2>Абонементы</h2>
//       <ul className="list">
//         {memberships.map(membership => (
//           <li key={membership.ID} className="list-item">
//             <p>Тип: {membership.Type}</p>
//             <p>Описание: {membership.Description}</p>
//             <p>Цена: {membership.Price}</p>
//             <p>Длительность: {membership.DaysDuration} дней</p>
//           </li>
//         ))}
//       </ul>
//       </div>
//     )
//     }
//   };

//   const renderEquipments = () => {
//     if (view === "equipments") {
//       return (
//       <div>
//         <h2>Экипировка</h2>
//         <ul className="list">
//           {equipments.map(equipment => (
//             <li key={equipment.ID} className="list-item">
//               <p>Название: {equipment.Name}</p>
//               <p>Описание: {equipment.Description}</p>
//             </li>
//           ))}
//         </ul>
//       </div>
//     )
//    }
//   };

//   const renderTrainers = () => {
//   if (view === 'trainers') {
//     return (
//     <div>
//       <h2>Тренеры</h2>
//       <ul className="list">
//         {trainers.map(trainer => (
//           <li key={trainer.id} className="list-item">
//             <p>ФИО: {trainer.Fullname}</p>
//             <p>Email: {trainer.Email}</p>
//             <p>Телефон: {trainer.Phone}</p>
//             <p>Квалификация: {trainer.Qualification}</p>
//             <p>Цена в час: ${trainer.UnitPrice}</p>
//           </li>
//         ))}
//       </ul>
//     </div>
//    )
//     }
//   };

//   if (loading) {
//     return <div>Загрузка, пожалуйста, подождите...</div>;
//   }

//   return (
//     <div className="gym-page">
//       <h1>{gym.Name}</h1>
//       <p>Телефон: {gym.Phone}</p>
//       <p>Город: {gym.City}</p>
//       <p>Адрес: {gym.Addres}</p>
//       <p>Сетевой: {gym.IsChain ? 'Да' : 'Нет'}</p>

//       <div className="buttons">
//         <button onClick={fetchMemberships}>Абонементы</button>
//         <button onClick={fetchEquipments}>Экипировка</button>
//         <button onClick={fetchTrainers}>Тренеры</button>
//       </div>

//       {renderMemberships()}
//       {renderEquipments()}
//       {renderTrainers()}
//       {/* {view === 'memberships' && (
//         <div>
//           <h2>Membership Types</h2>
//           <ul className="list">
//             {memberships.map(membership => (
//               <li key={membership.id} className="list-item">
//                 <p>Type: {membership.type}</p>
//                 <p>Description: {membership.description}</p>
//                 <p>Price: {membership.price}</p>
//                 <p>Duration: {membership.daysDuration} days</p>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )}

//       {view === 'equipments' && (
//         <div>
//           <h2>Equipments</h2>
//           <ul className="list">
//             {equipments.map(equipment => (
//               <li key={equipment.id} className="list-item">
//                 <p>Name: {equipment.name}</p>
//                 <p>Description: {equipment.description}</p>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )}

//       {view === 'trainers' && (
//         <div>
//           <h2>Trainers</h2>
//           <ul className="list">
//             {trainers.map(trainer => (
//               <li key={trainer.id} className="list-item">
//                 <p>Fullname: {trainer.fullname}</p>
//                 <p>Email: {trainer.email}</p>
//                 <p>Phone: {trainer.phone}</p>
//                 <p>Qualification: {trainer.qualification}</p>
//                 <p>Unit Price: ${trainer.unitPrice}</p>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )} */}
//     </div>
//   );
// };

// export default Gym;

// import React, { useState, useEffect } from 'react';
// import { useParams, Link } from 'react-router-dom';
// import axios from 'axios';
// import './Gym.css';

// const Gym = ({ isLoggedIn }) => {
//   const { id } = useParams();
//   const [gym, setGym] = useState(null);
//   const [memberships, setMemberships] = useState([]);
//   const [equipments, setEquipments] = useState([]);
//   const [trainers, setTrainers] = useState([]);
//   const [loading, setLoading] = useState(true);
//   const [view, setView] = useState('');

//   useEffect(() => {
//     axios.get(`http://localhost:8080/api/v1/gym/${id}`)
//       .then(response => {
//         setGym(response.data.gym);
//         setLoading(false);
//       })
//       .catch(error => {
//         console.error('Error fetching gym details:', error);
//         setLoading(false);
//       });
//   }, [id]);

//   const fetchMemberships = () => {
//     axios.get(`http://localhost:8080/api/v1/membershipType/${id}`)
//       .then(response => {
//         setMemberships(response.data.membershipTypes);
//         setView('memberships');
//       })
//       .catch(error => {
//         console.error('Error fetching memberships:', error);
//       });
//   };

//   const fetchEquipments = () => {
//     axios.get(`http://localhost:8080/api/v1/equipment/${id}`)
//       .then(response => {
//         setEquipments(response.data.equipments);
//         setView('equipments');
//       })
//       .catch(error => {
//         console.error('Error fetching equipments:', error);
//       });
//   };

//   const fetchTrainers = () => {
//     axios.get(`http://localhost:8080/api/v1/trainer/${id}`)
//       .then(response => {
//         setTrainers(response.data.trainers);
//         setView('trainers');
//       })
//       .catch(error => {
//         console.error('Error fetching trainers:', error);
//       });
//   };

//   if (loading) {
//     return <div>Loading...</div>;
//   }

//   return (
//     <div className="halls-page">
//       <h1>{gym.Name}</h1>
//       <p>Phone: {gym.Phone}</p>
//       <p>City: {gym.City}</p>
//       <p>Address: {gym.Addres}</p>
//       <p>Chain: {gym.IsChain ? 'Yes' : 'No'}</p>

//       <div className="buttons">
//         <button onClick={fetchMemberships}>Show Memberships</button>
//         <button onClick={fetchEquipments}>Show Equipments</button>
//         <button onClick={fetchTrainers}>Show Trainers</button>
//       </div>

//       {view === 'memberships' && (
//         <div>
//           <h2>Membership Types</h2>
//           <ul className="list">
//             {memberships.map(membership => (
//               <li key={membership.ID} className="list-item">
//                 <p>Type: {membership.Type}</p>
//                 <p>Description: {membership.Description}</p>
//                 <p>Price: {membership.Price}</p>
//                 <p>Duration: {membership.DaysDuration} days</p>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )}

//       {view === 'equipments' && (
//         <div>
//           <h2>Equipments</h2>
//           <ul className="list">
//             {equipments.map(equipment => (
//               <li key={equipment.ID} className="list-item">
//                 <p>Name: {equipment.Name}</p>
//                 <p>Description: {equipment.Description}</p>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )}

//       {view === 'trainers' && (
//         <div>
//           <h2>Trainers</h2>
//           <ul className="list">
//             {trainers.map(trainer => (
//               <li key={trainer.id} className="list-item">
//                 <p>Fullname: {trainer.Fullname}</p>
//                 <p>Email: {trainer.Email}</p>
//                 <p>Phone: {trainer.Phone}</p>
//                 <p>Qualification: {trainer.Qualification}</p>
//                 <p>Unit Price: ${trainer.UnitPrice}</p>
//                 {/* <button></> */}
//                 <Link to={`/training/${trainer.ID}`}>View Trainings</Link>
//               </li>
//             ))}
//           </ul>
//         </div>
//       )}
//     </div>
//   );
// };

// export default Gym;

import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import axios from 'axios';
import './Gym.css';
import { v4 as uuidv4 } from 'uuid';
import { format, parseISO } from 'date-fns';
import { ru } from 'date-fns/locale';

const Gym = ({ isAuthenticated, clientId }) => {
  const { id } = useParams();
  const [gym, setGym] = useState(null);
  const [memberships, setMemberships] = useState([]);
  const [equipments, setEquipments] = useState([]);
  const [trainers, setTrainers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [view, setView] = useState('');
  const [selectedMembership, setSelectedMembership] = useState(null);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [price, setPrice] = useState(0);

  useEffect(() => {
    axios.get(`http://localhost:8080/api/v1/gym/${id}`)
      .then(response => {
        setGym(response.data.gym);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching gym details:', error);
        setLoading(false);
      });
  }, [id]);

  const fetchMemberships = () => {
    axios.get(`http://localhost:8080/api/v1/membershipType/${id}`)
      .then(response => {
        setMemberships(response.data.membershipTypes);
        setView('memberships');
      })
      .catch(error => {
        console.error('Error fetching memberships:', error);
      });
  };

  const fetchEquipments = () => {
    axios.get(`http://localhost:8080/api/v1/equipment/${id}`)
      .then(response => {
        setEquipments(response.data.equipments);
        setView('equipments');
      })
      .catch(error => {
        console.error('Error fetching equipments:', error);
      });
  };

  const fetchTrainers = () => {
    axios.get(`http://localhost:8080/api/v1/trainer/${id}`)
      .then(response => {
        setTrainers(response.data.trainers);
        setView('trainers');
      })
      .catch(error => {
        console.error('Error fetching trainers:', error);
      });
  };

  const handlePurchase = (membership) => {
    setSelectedMembership(membership);
  };

  const calculatePrice = (start, end, membership) => {
    const startDay = parseISO(start);
    const endDay = parseISO(end);
    const days = (endDay - startDay) / (1000 * 60 * 60 * 24);
    return ((days / membership.DaysDuration) * membership.Price).toFixed(2);
  };

  const handleStartDateChange = (e) => {
    setStartDate(e.target.value);
    if (endDate) {
      setPrice(calculatePrice(e.target.value, endDate, selectedMembership));
    }
  };

  const handleEndDateChange = (e) => {
    setEndDate(e.target.value);
    if (startDate) {
      setPrice(calculatePrice(startDate, e.target.value, selectedMembership));
    }
  };

  const handlePayment = () => {
    const paymentData = {
      amount: {
        value: price,
        currency: 'RUB'
      },
      confirmation: {
        type: 'redirect',
        return_url: 'http://localhost:3000/'
      },
      capture: true,
      description: `Purchase of membership: ${selectedMembership.Type}`
    };

    axios.post('http://localhost:8080/create-payment', paymentData)
      .then(response => {
        console.log(response.data)
        window.location.href = response.data.confirmation.confirmation_url;
        axios.post('http://localhost:8080/api/v1/client_membership/new', {
          id: uuidv4(),
          startdate: startDate,
          enddate: endDate,
          membershiptypeid: selectedMembership.ID,
          clientid: clientId 
        })
        .then(() => {
          // alert('Membership purchased successfully');
        })
        .catch(error => {
          console.error('Error creating client membership:', error);
          // alert('Failed to create client membership');
        });
      })
      .catch(error => {
        console.error('Error creating payment:', error);
        // alert('Failed to create payment');
      });
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <div className="halls-page">
      <h1>{gym.Name}</h1>
      <p>Телефон: {gym.Phone}</p>
      <p>Город: {gym.City}</p>
      <p>Адрес: {gym.Addres}</p>
      {/* <p>Chain: {gym.IsChain ? 'Yes' : 'No'}</p> */}

      <div className="buttons">
        { isAuthenticated && <button onClick={fetchMemberships}>Абонементы зала</button>}
        <button onClick={fetchEquipments}>Экипировка зала</button>
        <button onClick={fetchTrainers}>Тренеры</button>
      </div>

      {view === 'memberships' &&  (
        <div>
          <h2>Абонементы зала</h2>
          <ul className="list">
            {memberships.map(membership => (
              <li key={membership.ID} className="list-item">
                <p>Тип: {membership.Type}</p>
                <p>Описание: {membership.Description}</p>
                <p>Цена: {membership.Price} руб</p>
                <p>За время: {membership.DaysDuration} дней</p>
                <button onClick={() => handlePurchase(membership)}>Купить</button>
              </li>
            ))}
          </ul>
        </div>
      )}

      {view === 'equipments' && (
        <div>
          <h2>Экипировка</h2>
          <ul className="list">
            {equipments.map(equipment => (
              <li key={equipment.ID} className="list-item">
                <p>Имя: {equipment.Name}</p>
                <p>Описание: {equipment.Description}</p>
              </li>
            ))}
          </ul>
        </div>
      )}

      {view === 'trainers' && (
        <div>
          <h2>Тренеры</h2>
          <ul className="list">
            {trainers.map(trainer => (
              <li key={trainer.id} className="list-item">
                <p>ФИО: {trainer.Fullname}</p>
                <p>Email: {trainer.Email}</p>
                <p>Телефон: {trainer.Phone}</p>
                <p>Квалификация: {trainer.Qualification}</p>
                <p>Цена в час: {trainer.UnitPrice} руб</p>
                <Link to={`/training/${trainer.ID}`}>Показать тренировки</Link>
              </li>
            ))}
          </ul>
        </div>
      )}

      {selectedMembership && (
        <div className="purchase-form">
          <h2>Покупка абонемента</h2>
          <label>
            Начальная дата:
            <input type="date" value={startDate} onChange={handleStartDateChange} />
          </label>
          <label>
            Конечная дата:
            <input type="date" value={endDate} onChange={handleEndDateChange} />
          </label>
          <p>Итоговая цена: {price} руб</p>
          <button onClick={handlePayment}>Перейти к оплате</button>
        </div>
      )}
    </div>
  );
};

export default Gym;


