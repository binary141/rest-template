import { ref } from 'vue';
import { defineStore } from 'pinia';
import api, { setGlobalRequestHeader } from '@/utils/api';

export const useUsersStore = defineStore('users', () => {
  const storedToken = localStorage.getItem('token');
  const storedUser = JSON.parse(localStorage.getItem('user') || 'null');
  const token = ref(storedToken);
  const loggedIn = ref(storedToken !== null);
  const user = ref(storedUser);

  // Restore auth header on page load
  if (storedToken) {
    setGlobalRequestHeader('Authorization', `Token ${storedToken}`);
  }

  function setToken(value) {
    token.value = value;
    if (value !== null) {
      localStorage.setItem('token', value);
      setGlobalRequestHeader('Authorization', `Token ${value}`);
    } else {
      localStorage.removeItem('token');
      setGlobalRequestHeader('Authorization', undefined);
    }
  }

  function logIn(userData) {
    user.value = userData;
    loggedIn.value = true;
    localStorage.setItem('user', JSON.stringify(userData));
  }

  async function logInAPI(credentials) {
    return api({ url: 'login', method: 'POST', json: credentials });
  }

  async function logOut() {
    const response = await api({ url: 'logout', method: 'POST' });
    loggedIn.value = false;
    user.value = null;
    localStorage.removeItem('user');
    setToken(null);
    return response;
  }

  return { token, loggedIn, user, setToken, logIn, logInAPI, logOut };
});
