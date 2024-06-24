// import React, { useState, useEffect } from 'react';
// import axios from 'axios';
// import {v4 as uuidv4} from 'uuid';
// import './AdminProfile.css';

// const AdminProfile = () => {
//   const [gyms, setGyms] = useState([]);
//   const [trainers, setTrainers] = useState([]);
//   const [view, setView] = useState('gyms');
//   const [selectedGym, setSelectedGym] = useState(null);
//   const [equipments, setEquipments] = useState(null);
//   const [newGym, setNewGym] = useState({ Name: '', Phone: '', City: '', Address: '', IsChain: true });
//   const [newTrainer, setNewTrainer] = useState({ Fullname: '', Email: '', Phone: '', Qualification: '', UnitPrice: '' });
//   const [newEquipment, setNewEquipment] = useState({ ID: uuidv4(), Name: '', Description: '', GymID: null });
//   const [newMembershipType, setNewMembershipType] = useState({ ID: uuidv4(), Type: '', Description: '', Price: 0, DaysDuration: 0, GymID: null });
//   const [membershipTypes, setMembershipTypes] = useState(null);

//   useEffect(() => {
//     fetchGyms();
//     fetchTrainers();
//   }, []);

//   const fetchGyms = () => {
//     axios.get('http://localhost:8080/api/v1/gym/all')
//       .then(response => {
//         setGyms(response.data.gyms);
//       })
//       .catch(error => {
//         console.error('Error fetching gyms:', error);
//       });
//   };

//   const fetchTrainers = () => {
//     axios.get('http://localhost:8080/api/v1/trainer/all')
//       .then(response => {
//         setTrainers(response.data.trainers);
//       })
//       .catch(error => {
//         console.error('Error fetching trainers:', error);
//       });
//   };

//   const handleAddGym = () => {
//     axios.post('http://localhost:8080/api/v1/gym/new', newGym)
//       .then(_ => {
//         setGyms([...gyms, newGym]);
//         setNewGym({ Name: '', Phone: '', City: '', Address: '', IsChain: true });
//       })
//       .catch(error => {
//         console.error('Error adding gym:', error);
//       });
//   };

//   const handleAddTrainer = () => {
//     axios.post('http://localhost:8080/api/v1/trainer/new', newTrainer)
//       .then(response => {
//         setTrainers([...trainers, response.data.trainer]);
//         setNewTrainer({ Fullname: '', Email: '', Phone: '', Qualification: '', UnitPrice: '' });
//       })
//       .catch(error => {
//         console.error('Error adding trainer:', error);
//       });
//   };

//   const handleAddEquipment = (gymId) => {
//     newEquipment.GymID = gymId;
//     axios.post(`http://localhost:8080/api/v1/equipment/new`, newEquipment)
//       .then(response => {
//         setSelectedGym({
//           ...selectedGym,
//           // Equipments: [...selectedGym.Equipments, response.data.equipment]
//         });
//         setEquipments([...equipments, newEquipment]);
//         setNewEquipment({ ID: uuidv4(), Name: '', Description: '', GymID: gymId});
//       })
//       .catch(error => {
//         console.error('Error adding equipment:', error);
//       });
//   };

//   const handleAddMembershipType = (gymId) => {
//     axios.post(`http://localhost:8080/api/v1/membershipType/new`, newMembershipType)
//       .then(response => {
//         setSelectedGym({
//           ...selectedGym,
//           MembershipTypes: [...selectedGym.MembershipTypes, response.data.membershipType]
//         });
//         setNewMembershipType({ ID: uuidv4(), Type: '', Description: '', Price: 0, DaysDuration: 0, GymID: gymId });
//       })
//       .catch(error => {
//         console.error('Error adding membership type:', error);
//       });
//   };

//   const handleEditGym = (gym) => {
//     axios.put(`http://localhost:8080/api/v1/gym/${gym.ID}`, gym)
//       .then(response => {
//         setGyms(gyms.map(g => (g.ID === gym.ID ? response.data.gym : g)));
//       })
//       .catch(error => {
//         console.error('Error editing gym:', error);
//       });
//   };

//   const handleEditTrainer = (trainer) => {
//     axios.put(`http://localhost:8080/api/v1/trainer/${trainer.ID}`, trainer)
//       .then(response => {
//         setTrainers(trainers.map(t => (t.ID === trainer.ID ? response.data.trainer : t)));
//       })
//       .catch(error => {
//         console.error('Error editing trainer:', error);
//       });
//   };

//   const renderGyms = () => (
//     <div>
//       <h2>Залы</h2>
//       {gyms.map(gym => (
//         <div key={gym.ID} className="gym-item">
//           <p>Название: {gym.Name}</p>
//           <p>Телефон: {gym.Phone}</p>
//           <p>Город: {gym.City}</p>
//           <p>Адрес: {gym.Addres}</p>
//           {/* <p>Сеть: {gym.IsChain ? 'Да' : 'Нет'}</p> */}
//           <button onClick={() => controlGym(gym)}>Управление</button>
//         </div>
//       ))}
//       <div className="new-gym-form">
//         <h3>Добавить новый зал</h3>
//         <label>
//           Название:
//           <input type="text" value={newGym.Name} onChange={(e) => setNewGym({ ...newGym, Name: e.target.value })} />
//         </label>
//         <label>
//           Телефон:
//           <input type="text" value={newGym.Phone} onChange={(e) => setNewGym({ ...newGym, Phone: e.target.value })} />
//         </label>
//         <label>
//           Город:
//           <input type="text" value={newGym.City} onChange={(e) => setNewGym({ ...newGym, City: e.target.value })} />
//         </label>
//         <label>
//           Адрес:
//           <input type="text" value={newGym.Address} onChange={(e) => setNewGym({ ...newGym, Address: e.target.value })} />
//         </label>
//         {/* <label>
//           Сеть:
//           <input type="checkbox" checked={newGym.IsChain} onChange={(e) => setNewGym({ ...newGym, IsChain: e.target.checked })} />
//         </label> */}
//         <button onClick={handleAddGym}>Добавить зал</button>
//       </div>
//     </div>
//   );

//   const controlGym = (gym) => {
//     setSelectedGym(gym);
//     axios.get(`http://localhost:8080/api/v1/equipment/${gym.ID}`)
//     .then( response => {
//       setEquipments(response.data.equipments);
//     })
//     .catch( error => {
//       console.log(error)
//     })
//     axios.get(`http://localhost:8080/api/v1/membershipType/${gym.ID}`)
//     .then( response => {
//       setMembershipTypes(response.data.membershipTypes);
//       console.log(response.data.membershipTypes);
//       console.log(membershipTypes);
//     })
//     .catch( error => {
//       console.log(error)
//     })
//   }

//   const renderTrainers = () => (
//     <div>
//       <h2>Тренеры</h2>
//       {trainers.map(trainer => (
//         <div key={trainer.ID} className="trainer-item">
//           <p>ФИО: {trainer.Fullname}</p>
//           <p>Email: {trainer.Email}</p>
//           <p>Телефон: {trainer.Phone}</p>
//           <p>Квалификация: {trainer.Qualification}</p>
//           <p>Стоимость за единицу: {trainer.UnitPrice}</p>
//           <button onClick={() => setSelectedGym(trainer)}>Редактировать</button>
//         </div>
//       ))}
//       <div className="new-trainer-form">
//         <h3>Добавить нового тренера</h3>
//         <label>
//           ФИО:
//           <input type="text" value={newTrainer.Fullname} onChange={(e) => setNewTrainer({ ...newTrainer, Fullname: e.target.value })} />
//         </label>
//         <label>
//           Email:
//           <input type="email" value={newTrainer.Email} onChange={(e) => setNewTrainer({ ...newTrainer, Email: e.target.value })} />
//         </label>
//         <label>
//           Телефон:
//           <input type="text" value={newTrainer.Phone} onChange={(e) => setNewTrainer({ ...newTrainer, Phone: e.target.value })} />
//         </label>
//         <label>
//           Квалификация:
//           <input type="text" value={newTrainer.Qualification} onChange={(e) => setNewTrainer({ ...newTrainer, Qualification: e.target.value })} />
//         </label>
//         <label>
//           Стоимость за единицу:
//           <input type="number" value={newTrainer.UnitPrice} onChange={(e) => setNewTrainer({ ...newTrainer, UnitPrice: e.target.value })} />
//         </label>
//         <button onClick={handleAddTrainer}>Добавить тренера</button>
//       </div>
//     </div>
//   );

//   return (
//     <div className="admin-profile">
//       <h1>Панель администратора</h1>
//       <div className="buttons">
//         <button onClick={() => setView('gyms')}>Залы</button>
//         <button onClick={() => setView('trainers')}>Тренеры</button>
//       </div>

//       {view === 'gyms' && renderGyms()}
//       {view === 'trainers' && renderTrainers()}

//       {view === 'gyms' && equipments && membershipTypes && selectedGym &&(
//         <div className="gym-management">
//           <h2>Управление залом: {selectedGym.Name}</h2>
//           <div>
//             <h3>Оборудование</h3>
//             <ul>
//               {equipments.map(equipment => (
//                 <li key={equipment.ID}>
//                   {equipment.Name}: {equipment.Description}
//                 </li>
//               ))}
//             </ul>
//             <div>
//               <h4>Добавить оборудование</h4>
//               <label>
//                 Название:
//                 <input type="text" value={newEquipment.Name} onChange={(e) => setNewEquipment({ ...newEquipment, Name: e.target.value })} />
//               </label>
//               <label>
//                 Описание:
//                 <input type="text" value={newEquipment.Description} onChange={(e) => setNewEquipment({ ...newEquipment, Description: e.target.value })} />
//               </label>
//               <button onClick={() => handleAddEquipment(selectedGym.ID)}>Добавить оборудование</button>
//             </div>
//           </div>

//           <div>
//             <h3>Типы абонементов</h3>
//             <ul>
//               {membershipTypes.map(membershipType => (
//                 <li key={membershipType.ID}>
//                   {membershipType.Type}: {membershipType.Description}, {membershipType.Price} руб, {membershipType.DaysDuration} дней
//                 </li>
//               ))}
//             </ul>
//             <div>
//               <h4>Добавить тип абонемента</h4>
//               <label>
//                 Тип:
//                 <input type="text" value={newMembershipType.Type} onChange={(e) => setNewMembershipType({ ...newMembershipType, Type: e.target.value })} />
//               </label>
//               <label>
//                 Описание:
//                 <input type="text" value={newMembershipType.Description} onChange={(e) => setNewMembershipType({ ...newMembershipType, Description: e.target.value })} />
//               </label>
//               <label>
//                 Цена:
//                 <input type="number" value={newMembershipType.Price} onChange={(e) => setNewMembershipType({ ...newMembershipType, Price: e.target.value })} />
//               </label>
//               <label>
//                 Длительность (дней):
//                 <input type="number" value={newMembershipType.DaysDuration} onChange={(e) => setNewMembershipType({ ...newMembershipType, DaysDuration: e.target.value })} />
//               </label>
//               <button onClick={() => handleAddMembershipType(selectedGym.ID)}>Добавить тип абонемента</button>
//             </div>
//           </div>
//         </div>
//       )}
//     </div>
//   );
// };

// export default AdminProfile;

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';
import './AdminProfile.css';

const AdminProfile = () => {
  const [gyms, setGyms] = useState([]);
  const [trainers, setTrainers] = useState([]);
  const [view, setView] = useState('gyms');
  const [selectedGym, setSelectedGym] = useState(null);
  const [equipments, setEquipments] = useState([]);
  const [membershipTypes, setMembershipTypes] = useState([]);
  const [newGym, setNewGym] = useState({ Name: '', Phone: '', City: '', Address: '', IsChain: true });
  const [newTrainer, setNewTrainer] = useState({ ID: uuidv4(), Fullname: '', Email: '', Phone: '', Qualification: '', UnitPrice: '', GymsID: [] });
  const [newEquipment, setNewEquipment] = useState({ ID: uuidv4(), Name: '', Description: '', GymID: null });
  const [newMembershipType, setNewMembershipType] = useState({ ID: uuidv4(), Type: '', Description: '', Price: 0, DaysDuration: 0, GymID: null });

  useEffect(() => {
    fetchGyms();
    fetchTrainers();
  }, []);

  const fetchGyms = () => {
    axios.get('http://localhost:8080/api/v1/gym/all')
      .then(response => {
        setGyms(response.data.gyms);
      })
      .catch(error => {
        console.error('Error fetching gyms:', error);
      });
  };

  const fetchTrainers = () => {
    axios.get('http://localhost:8080/api/v1/trainer/all')
      .then(response => {
        setTrainers(response.data.trainers);
      })
      .catch(error => {
        console.error('Error fetching trainers:', error);
      });
  };

  const handleAddGym = () => {
    axios.post('http://localhost:8080/api/v1/gym/new', newGym)
      .then(_ => {
        setGyms([...gyms, newGym]);
        setNewGym({ Name: '', Phone: '', City: '', Address: '', IsChain: true });
      })
      .catch(error => {
        console.error('Error adding gym:', error);
      });
  };

  const handleAddTrainer = () => {
    axios.post('http://localhost:8080/api/v1/trainer/new', newTrainer)
      .then(response => {
        setTrainers([...trainers, newTrainer]);
        setNewTrainer({ Fullname: '', Email: '', Phone: '', Qualification: '', UnitPrice: '', GymsID: [] });
      })
      .catch(error => {
        console.error('Error adding trainer:', error);
      });
  };

  const handleAddEquipment = (gymId) => {
    const equipment = { ...newEquipment, GymID: gymId };
    axios.post('http://localhost:8080/api/v1/equipment/new', equipment)
      .then(response => {
        setEquipments([...equipments, response.data.equipment]);
        setNewEquipment({ ID: uuidv4(), Name: '', Description: '', GymID: gymId });
      })
      .catch(error => {
        console.error('Error adding equipment:', error);
      });
  };

  const handleAddMembershipType = (gymId) => {
    const membershipType = { ...newMembershipType, GymID: gymId };
    axios.post('http://localhost:8080/api/v1/membershipType/new', membershipType)
      .then(response => {
        setMembershipTypes([...membershipTypes, response.data.membershipType]);
        setNewMembershipType({ ID: uuidv4(), Type: '', Description: '', Price: 0, DaysDuration: 0, GymID: gymId });
      })
      .catch(error => {
        console.error('Error adding membership type:', error);
      });
  };

  const handleEditEquipment = (equipment) => {
    axios.put(`http://localhost:8080/api/v1/equipment/${equipment.ID}`, equipment)
      .then(response => {
        setEquipments(equipments.map(e => (e.ID === equipment.ID ? response.data.equipment : e)));
      })
      .catch(error => {
        console.error('Error editing equipment:', error);
      });
  };

  const handleEditMembershipType = (membershipType) => {
    axios.put(`http://localhost:8080/api/v1/membershipType/${membershipType.ID}`, membershipType)
      .then(response => {
        setMembershipTypes(membershipTypes.map(mt => (mt.ID === membershipType.ID ? response.data.membershipType : mt)));
      })
      .catch(error => {
        console.error('Error editing membership type:', error);
      });
  };

  const handleEditGym = (gym) => {
    axios.put(`http://localhost:8080/api/v1/gym/${gym.ID}`, gym)
      .then(response => {
        setGyms(gyms.map(g => (g.ID === gym.ID ? response.data.gym : g)));
      })
      .catch(error => {
        console.error('Error editing gym:', error);
      });
  };

  const handleEditTrainer = (trainer) => {
    axios.put(`http://localhost:8080/api/v1/trainer/${trainer.ID}`, trainer)
      .then(response => {
        setTrainers(trainers.map(t => (t.ID === trainer.ID ? response.data.trainer : t)));
      })
      .catch(error => {
        console.error('Error editing trainer:', error);
      });
  };

  const controlGym = (gym) => {
    setSelectedGym(gym);
    axios.get(`http://localhost:8080/api/v1/equipment/${gym.ID}`)
      .then(response => {
        setEquipments(response.data.equipments);
      })
      .catch(error => {
        console.error(error);
      });
    axios.get(`http://localhost:8080/api/v1/membershipType/${gym.ID}`)
      .then(response => {
        setMembershipTypes(response.data.membershipTypes);
      })
      .catch(error => {
        console.error(error);
      });
  };

  const renderGyms = () => (
    <div>
      <h2>Залы</h2>
      {gyms.map(gym => (
        <div key={gym.ID} className="gym-item">
          <p>Название: {gym.Name}</p>
          <p>Телефон: {gym.Phone}</p>
          <p>Город: {gym.City}</p>
          <p>Адрес: {gym.Address}</p>
          <button onClick={() => controlGym(gym)}>Управление</button>
        </div>
      ))}
      <div className="new-gym-form">
        <h3>Добавить новый зал</h3>
        <label>
          Название:
          <input type="text" value={newGym.Name} onChange={(e) => setNewGym({ ...newGym, Name: e.target.value })} />
        </label>
        <label>
          Телефон:
          <input type="text" value={newGym.Phone} onChange={(e) => setNewGym({ ...newGym, Phone: e.target.value })} />
        </label>
        <label>
          Город:
          <input type="text" value={newGym.City} onChange={(e) => setNewGym({ ...newGym, City: e.target.value })} />
        </label>
        <label>
          Адрес:
          <input type="text" value={newGym.Address} onChange={(e) => setNewGym({ ...newGym, Address: e.target.value })} />
        </label>
        <button onClick={handleAddGym}>Добавить зал</button>
      </div>
    </div>
  );

  const renderTrainers = () => (
    <div>
      <h2>Тренеры</h2>
      {trainers.map(trainer => (
        <div key={trainer.ID} className="trainer-item">
          <p>ФИО: {trainer.Fullname}</p>
          <p>Email: {trainer.Email}</p>
          <p>Телефон: {trainer.Phone}</p>
          <p>Квалификация: {trainer.Qualification}</p>
          <p>Стоимость за единицу: {trainer.UnitPrice}</p>
          <p>Зал: {gyms.find(gym => gym.ID === trainer.GymID)?.Name || 'Не назначен'}</p>
          <button onClick={() => setSelectedGym(trainer)}>Редактировать</button>
        </div>
      ))}
      <div className="new-trainer-form">
        <h3>Добавить нового тренера</h3>
        <label>
          ФИО:
          <input type="text" value={newTrainer.Fullname} onChange={(e) => setNewTrainer({ ...newTrainer, Fullname: e.target.value })} />
        </label>
        <label>
          Email:
          <input type="email" value={newTrainer.Email} onChange={(e) => setNewTrainer({ ...newTrainer, Email: e.target.value })} />
        </label>
        <label>
          Телефон:
          <input type="text" value={newTrainer.Phone} onChange={(e) => setNewTrainer({ ...newTrainer, Phone: e.target.value })} />
        </label>
        <label>
          Квалификация:
          <input type="text" value={newTrainer.Qualification} onChange={(e) => setNewTrainer({ ...newTrainer, Qualification: e.target.value })} />
        </label>
        <label>
          Стоимость за единицу:
          <input type="number" value={newTrainer.UnitPrice} onChange={(e) => setNewTrainer({ ...newTrainer, UnitPrice: parseFloat(e.target.value) })} />
        </label>
        <label>
          Зал:
          <select value={newTrainer.GymsID} onChange={(e) => setNewTrainer({ ...newTrainer, GymsID: [e.target.value] })}>
            <option value="">Выбрать зал</option>
            {gyms.map(gym => (
              <option key={gym.ID} value={gym.ID}>{gym.Name}</option>
            ))}
          </select>
        </label>
        <button onClick={handleAddTrainer}>Добавить тренера</button>
      </div>
    </div>
  );

  return (
    <div className="admin-profile">
      <h1>Панель администратора</h1>
      <div className="buttons">
        <button onClick={() => setView('gyms')}>Залы</button>
        <button onClick={() => setView('trainers')}>Тренеры</button>
      </div>

      {view === 'gyms' && renderGyms()}
      {view === 'trainers' && renderTrainers()}

      {view === 'gyms' && selectedGym && (
        <div className="gym-management">
          <h2>Управление залом: {selectedGym.Name}</h2>
          <div>
            <h3>Оборудование</h3>
            <ul>
              {equipments.map(equipment => (
                <li key={equipment.ID}>
                  {equipment.Name}: {equipment.Description}
                  <button onClick={() => handleEditEquipment(equipment)}>Редактировать</button>
                </li>
              ))}
            </ul>
            <div>
              <h4>Добавить оборудование</h4>
              <label>
                Название:
                <input type="text" value={newEquipment.Name} onChange={(e) => setNewEquipment({ ...newEquipment, Name: e.target.value })} />
              </label>
              <label>
                Описание:
                <input type="text" value={newEquipment.Description} onChange={(e) => setNewEquipment({ ...newEquipment, Description: e.target.value })} />
              </label>
              <button onClick={() => handleAddEquipment(selectedGym.ID)}>Добавить оборудование</button>
            </div>
          </div>

          <div>
            <h3>Типы абонементов</h3>
            <ul>
              {membershipTypes.map(membershipType => (
                <li key={membershipType.ID}>
                  {membershipType.Type}: {membershipType.Description}, {membershipType.Price} руб, {membershipType.DaysDuration} дней
                  <button onClick={() => handleEditMembershipType(membershipType)}>Редактировать</button>
                </li>
              ))}
            </ul>
            <div>
              <h4>Добавить тип абонемента</h4>
              <label>
                Тип:
                <input type="text" value={newMembershipType.Type} onChange={(e) => setNewMembershipType({ ...newMembershipType, Type: e.target.value })} />
              </label>
              <label>
                Описание:
                <input type="text" value={newMembershipType.Description} onChange={(e) => setNewMembershipType({ ...newMembershipType, Description: e.target.value })} />
              </label>
              <label>
                Цена:
                <input type="number" value={newMembershipType.Price} onChange={(e) => setNewMembershipType({ ...newMembershipType, Price: e.target.value })} />
              </label>
              <label>
                Длительность (дней):
                <input type="number" value={newMembershipType.DaysDuration} onChange={(e) => setNewMembershipType({ ...newMembershipType, DaysDuration: e.target.value })} />
              </label>
              <button onClick={() => handleAddMembershipType(selectedGym.ID)}>Добавить тип абонемента</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default AdminProfile;

