import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
});

export const registerClient = (clientReq) => {
  alert("send")
  return api.post('/register', clientReq);
};

export default api;
